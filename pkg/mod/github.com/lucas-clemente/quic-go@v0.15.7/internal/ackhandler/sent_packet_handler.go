package ackhandler

import (
	"fmt"
	"math"
	"time"

	"github.com/lucas-clemente/quic-go/internal/congestion"
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/qerr"
	"github.com/lucas-clemente/quic-go/internal/utils"
	"github.com/lucas-clemente/quic-go/internal/wire"
	"github.com/lucas-clemente/quic-go/qlog"
	"github.com/lucas-clemente/quic-go/quictrace"
)

const (
	// Maximum reordering in time space before time based loss detection considers a packet lost.
	// Specified as an RTT multiplier.
	timeThreshold = 9.0 / 8
	// Maximum reordering in packets before packet threshold loss detection considers a packet lost.
	packetThreshold = 3
)

type packetNumberSpace struct {
	history *sentPacketHistory
	pns     *packetNumberGenerator

	lossTime                       time.Time
	lastSentAckElicitingPacketTime time.Time

	largestAcked protocol.PacketNumber
	largestSent  protocol.PacketNumber
}

func newPacketNumberSpace(initialPN protocol.PacketNumber) *packetNumberSpace {
	return &packetNumberSpace{
		history:      newSentPacketHistory(),
		pns:          newPacketNumberGenerator(initialPN, protocol.SkipPacketAveragePeriodLength),
		largestSent:  protocol.InvalidPacketNumber,
		largestAcked: protocol.InvalidPacketNumber,
	}
}

type sentPacketHandler struct {
	nextSendTime time.Time

	initialPackets   *packetNumberSpace
	handshakePackets *packetNumberSpace
	appDataPackets   *packetNumberSpace

	peerCompletedAddressValidation bool
	handshakeComplete              bool

	// lowestNotConfirmedAcked is the lowest packet number that we sent an ACK for, but haven't received confirmation, that this ACK actually arrived
	// example: we send an ACK for packets 90-100 with packet number 20
	// once we receive an ACK from the peer for packet 20, the lowestNotConfirmedAcked is 101
	// Only applies to the application-data packet number space.
	lowestNotConfirmedAcked protocol.PacketNumber

	bytesInFlight protocol.ByteCount

	congestion congestion.SendAlgorithmWithDebugInfos
	rttStats   *congestion.RTTStats

	// The number of times a PTO has been sent without receiving an ack.
	ptoCount uint32
	ptoMode  SendMode
	// The number of PTO probe packets that should be sent.
	// Only applies to the application-data packet number space.
	numProbesToSend int

	// The alarm timeout
	alarm time.Time

	perspective protocol.Perspective

	traceCallback func(quictrace.Event)
	qlogger       qlog.Tracer
	logger        utils.Logger
}

var _ SentPacketHandler = &sentPacketHandler{}
var _ sentPacketTracker = &sentPacketHandler{}

func newSentPacketHandler(
	initialPacketNumber protocol.PacketNumber,
	rttStats *congestion.RTTStats,
	pers protocol.Perspective,
	traceCallback func(quictrace.Event),
	qlogger qlog.Tracer,
	logger utils.Logger,
) *sentPacketHandler {
	congestion := congestion.NewCubicSender(
		congestion.DefaultClock{},
		rttStats,
		true, // use Reno
	)

	return &sentPacketHandler{
		peerCompletedAddressValidation: pers == protocol.PerspectiveServer,
		initialPackets:                 newPacketNumberSpace(initialPacketNumber),
		handshakePackets:               newPacketNumberSpace(0),
		appDataPackets:                 newPacketNumberSpace(0),
		rttStats:                       rttStats,
		congestion:                     congestion,
		perspective:                    pers,
		traceCallback:                  traceCallback,
		qlogger:                        qlogger,
		logger:                         logger,
	}
}

func (h *sentPacketHandler) DropPackets(encLevel protocol.EncryptionLevel) {
	if h.perspective == protocol.PerspectiveClient && encLevel == protocol.EncryptionInitial {
		// This function is called when the crypto setup seals a Handshake packet.
		// If this Handshake packet is coalesced behind an Initial packet, we would drop the Initial packet number space
		// before SentPacket() was called for that Initial packet.
		return
	}
	h.dropPackets(encLevel)
}

func (h *sentPacketHandler) dropPackets(encLevel protocol.EncryptionLevel) {
	// The server won't await address validation after the handshake is confirmed.
	// This applies even if we didn't receive an ACK for a Handshake packet.
	if h.perspective == protocol.PerspectiveClient && encLevel == protocol.EncryptionHandshake {
		h.peerCompletedAddressValidation = true
	}
	// remove outstanding packets from bytes_in_flight
	if encLevel == protocol.EncryptionInitial || encLevel == protocol.EncryptionHandshake {
		pnSpace := h.getPacketNumberSpace(encLevel)
		pnSpace.history.Iterate(func(p *Packet) (bool, error) {
			if p.includedInBytesInFlight {
				h.bytesInFlight -= p.Length
			}
			return true, nil
		})
	}
	// drop the packet history
	switch encLevel {
	case protocol.EncryptionInitial:
		h.initialPackets = nil
	case protocol.EncryptionHandshake:
		h.handshakePackets = nil
	case protocol.Encryption0RTT:
		// TODO(#2067): invalidate sent data
		h.appDataPackets.history.Iterate(func(p *Packet) (bool, error) {
			if p.EncryptionLevel != protocol.Encryption0RTT {
				return false, nil
			}
			h.queueFramesForRetransmission(p)
			if p.includedInBytesInFlight {
				h.bytesInFlight -= p.Length
			}
			h.appDataPackets.history.Remove(p.PacketNumber)
			return true, nil
		})
	default:
		panic(fmt.Sprintf("Cannot drop keys for encryption level %s", encLevel))
	}
	h.setLossDetectionTimer()
	if h.qlogger != nil && h.ptoCount != 0 {
		h.qlogger.UpdatedPTOCount(0)
	}
	h.ptoCount = 0
	h.numProbesToSend = 0
	h.ptoMode = SendNone
}

func (h *sentPacketHandler) packetsInFlight() int {
	packetsInFlight := h.appDataPackets.history.Len()
	if h.handshakePackets != nil {
		packetsInFlight += h.handshakePackets.history.Len()
	}
	if h.initialPackets != nil {
		packetsInFlight += h.initialPackets.history.Len()
	}
	return packetsInFlight
}

func (h *sentPacketHandler) SentPacket(packet *Packet) {
	// For the client, drop the Initial packet number space when the first Handshake packet is sent.
	if h.perspective == protocol.PerspectiveClient && packet.EncryptionLevel == protocol.EncryptionHandshake && h.initialPackets != nil {
		h.dropPackets(protocol.EncryptionInitial)
	}
	isAckEliciting := h.sentPacketImpl(packet)
	if isAckEliciting {
		h.getPacketNumberSpace(packet.EncryptionLevel).history.SentPacket(packet)
	}
	if h.qlogger != nil && isAckEliciting {
		h.qlogger.UpdatedMetrics(h.rttStats, h.congestion.GetCongestionWindow(), h.bytesInFlight, h.packetsInFlight())
	}
	if isAckEliciting || !h.peerCompletedAddressValidation {
		h.setLossDetectionTimer()
	}
}

func (h *sentPacketHandler) getPacketNumberSpace(encLevel protocol.EncryptionLevel) *packetNumberSpace {
	switch encLevel {
	case protocol.EncryptionInitial:
		return h.initialPackets
	case protocol.EncryptionHandshake:
		return h.handshakePackets
	case protocol.Encryption0RTT, protocol.Encryption1RTT:
		return h.appDataPackets
	default:
		panic("invalid packet number space")
	}
}

func (h *sentPacketHandler) sentPacketImpl(packet *Packet) bool /* is ack-eliciting */ {
	pnSpace := h.getPacketNumberSpace(packet.EncryptionLevel)

	if h.logger.Debug() && pnSpace.history.HasOutstandingPackets() {
		for p := utils.MaxPacketNumber(0, pnSpace.largestSent+1); p < packet.PacketNumber; p++ {
			h.logger.Debugf("Skipping packet number %d", p)
		}
	}

	pnSpace.largestSent = packet.PacketNumber
	isAckEliciting := len(packet.Frames) > 0

	if isAckEliciting {
		pnSpace.lastSentAckElicitingPacketTime = packet.SendTime
		packet.includedInBytesInFlight = true
		h.bytesInFlight += packet.Length
		if h.numProbesToSend > 0 {
			h.numProbesToSend--
		}
	}
	h.congestion.OnPacketSent(packet.SendTime, h.bytesInFlight, packet.PacketNumber, packet.Length, isAckEliciting)

	h.nextSendTime = utils.MaxTime(h.nextSendTime, packet.SendTime).Add(h.congestion.TimeUntilSend(h.bytesInFlight))
	return isAckEliciting
}

func (h *sentPacketHandler) ReceivedAck(ack *wire.AckFrame, encLevel protocol.EncryptionLevel, rcvTime time.Time) error {
	pnSpace := h.getPacketNumberSpace(encLevel)

	largestAcked := ack.LargestAcked()
	if largestAcked > pnSpace.largestSent {
		return qerr.NewError(qerr.ProtocolViolation, "Received ACK for an unsent packet")
	}

	pnSpace.largestAcked = utils.MaxPacketNumber(pnSpace.largestAcked, largestAcked)

	if !pnSpace.pns.Validate(ack) {
		return qerr.NewError(qerr.ProtocolViolation, "Received an ACK for a skipped packet number")
	}

	// Servers complete address validation when a protected packet is received.
	if h.perspective == protocol.PerspectiveClient && !h.peerCompletedAddressValidation &&
		(encLevel == protocol.EncryptionHandshake || encLevel == protocol.Encryption1RTT) {
		h.peerCompletedAddressValidation = true
		h.logger.Debugf("Peer doesn't await address validation any longer.")
		// Make sure that the timer is reset, even if this ACK doesn't acknowledge any (ack-eliciting) packets.
		h.setLossDetectionTimer()
	}

	// maybe update the RTT
	if p := pnSpace.history.GetPacket(ack.LargestAcked()); p != nil {
		// don't use the ack delay for Initial and Handshake packets
		var ackDelay time.Duration
		if encLevel == protocol.Encryption1RTT {
			ackDelay = utils.MinDuration(ack.DelayTime, h.rttStats.MaxAckDelay())
		}
		h.rttStats.UpdateRTT(rcvTime.Sub(p.SendTime), ackDelay, rcvTime)
		if h.logger.Debug() {
			h.logger.Debugf("\tupdated RTT: %s (??: %s)", h.rttStats.SmoothedRTT(), h.rttStats.MeanDeviation())
		}
		h.congestion.MaybeExitSlowStart()
		if h.qlogger != nil {
			h.qlogger.UpdatedMetrics(h.rttStats, h.congestion.GetCongestionWindow(), h.bytesInFlight, h.packetsInFlight())
		}
	}

	priorInFlight := h.bytesInFlight
	ackedPackets, err := h.detectAndRemoveAckedPackets(ack, encLevel)
	if err != nil || len(ackedPackets) == 0 {
		return err
	}
	lostPackets, err := h.detectAndRemoveLostPackets(rcvTime, encLevel)
	if err != nil {
		return err
	}
	for _, p := range lostPackets {
		h.congestion.OnPacketLost(p.PacketNumber, p.Length, priorInFlight)
	}
	for _, p := range ackedPackets {
		if p.includedInBytesInFlight {
			h.congestion.OnPacketAcked(p.PacketNumber, p.Length, priorInFlight, rcvTime)
		}
	}

	if h.qlogger != nil && h.ptoCount != 0 {
		h.qlogger.UpdatedPTOCount(0)
	}
	h.ptoCount = 0
	h.numProbesToSend = 0

	h.setLossDetectionTimer()
	return nil
}

func (h *sentPacketHandler) GetLowestPacketNotConfirmedAcked() protocol.PacketNumber {
	return h.lowestNotConfirmedAcked
}

func (h *sentPacketHandler) detectAndRemoveAckedPackets(ack *wire.AckFrame, encLevel protocol.EncryptionLevel) ([]*Packet, error) {
	pnSpace := h.getPacketNumberSpace(encLevel)
	var ackedPackets []*Packet
	ackRangeIndex := 0
	lowestAcked := ack.LowestAcked()
	largestAcked := ack.LargestAcked()
	err := pnSpace.history.Iterate(func(p *Packet) (bool, error) {
		// Ignore packets below the lowest acked
		if p.PacketNumber < lowestAcked {
			return true, nil
		}
		// Break after largest acked is reached
		if p.PacketNumber > largestAcked {
			return false, nil
		}

		if ack.HasMissingRanges() {
			ackRange := ack.AckRanges[len(ack.AckRanges)-1-ackRangeIndex]

			for p.PacketNumber > ackRange.Largest && ackRangeIndex < len(ack.AckRanges)-1 {
				ackRangeIndex++
				ackRange = ack.AckRanges[len(ack.AckRanges)-1-ackRangeIndex]
			}

			if p.PacketNumber >= ackRange.Smallest { // packet i contained in ACK range
				if p.PacketNumber > ackRange.Largest {
					return false, fmt.Errorf("BUG: ackhandler would have acked wrong packet %d, while evaluating range %d -> %d", p.PacketNumber, ackRange.Smallest, ackRange.Largest)
				}
				ackedPackets = append(ackedPackets, p)
			}
		} else {
			ackedPackets = append(ackedPackets, p)
		}
		return true, nil
	})
	if h.logger.Debug() && len(ackedPackets) > 0 {
		pns := make([]protocol.PacketNumber, len(ackedPackets))
		for i, p := range ackedPackets {
			pns[i] = p.PacketNumber
		}
		h.logger.Debugf("\tnewly acked packets (%d): %d", len(pns), pns)
	}

	for _, p := range ackedPackets {
		if packet := pnSpace.history.GetPacket(p.PacketNumber); packet == nil {
			continue
		}
		if p.LargestAcked != protocol.InvalidPacketNumber && encLevel == protocol.Encryption1RTT {
			h.lowestNotConfirmedAcked = utils.MaxPacketNumber(h.lowestNotConfirmedAcked, p.LargestAcked+1)
		}

		for _, f := range p.Frames {
			if f.OnAcked != nil {
				f.OnAcked(f.Frame)
			}
		}
		if p.includedInBytesInFlight {
			h.bytesInFlight -= p.Length
		}
		if err := pnSpace.history.Remove(p.PacketNumber); err != nil {
			return nil, err
		}
	}

	return ackedPackets, err
}

func (h *sentPacketHandler) getEarliestLossTimeAndSpace() (time.Time, protocol.EncryptionLevel) {
	var encLevel protocol.EncryptionLevel
	var lossTime time.Time

	if h.initialPackets != nil {
		lossTime = h.initialPackets.lossTime
		encLevel = protocol.EncryptionInitial
	}
	if h.handshakePackets != nil && (lossTime.IsZero() || (!h.handshakePackets.lossTime.IsZero() && h.handshakePackets.lossTime.Before(lossTime))) {
		lossTime = h.handshakePackets.lossTime
		encLevel = protocol.EncryptionHandshake
	}
	if h.handshakeComplete &&
		(lossTime.IsZero() || (!h.appDataPackets.lossTime.IsZero() && h.appDataPackets.lossTime.Before(lossTime))) {
		lossTime = h.appDataPackets.lossTime
		encLevel = protocol.Encryption1RTT
	}
	return lossTime, encLevel
}

// same logic as getEarliestLossTimeAndSpace, but for lastSentAckElicitingPacketTime instead of lossTime
func (h *sentPacketHandler) getEarliestSentTimeAndSpace() (time.Time, protocol.EncryptionLevel) {
	var encLevel protocol.EncryptionLevel
	var sentTime time.Time

	if h.initialPackets != nil {
		sentTime = h.initialPackets.lastSentAckElicitingPacketTime
		encLevel = protocol.EncryptionInitial
	}
	if h.handshakePackets != nil && (sentTime.IsZero() || (!h.handshakePackets.lastSentAckElicitingPacketTime.IsZero() && h.handshakePackets.lastSentAckElicitingPacketTime.Before(sentTime))) {
		sentTime = h.handshakePackets.lastSentAckElicitingPacketTime
		encLevel = protocol.EncryptionHandshake
	}
	if h.handshakeComplete &&
		(sentTime.IsZero() || (!h.appDataPackets.lastSentAckElicitingPacketTime.IsZero() && h.appDataPackets.lastSentAckElicitingPacketTime.Before(sentTime))) {
		sentTime = h.appDataPackets.lastSentAckElicitingPacketTime
		encLevel = protocol.Encryption1RTT
	}
	return sentTime, encLevel
}

func (h *sentPacketHandler) hasOutstandingCryptoPackets() bool {
	var hasInitial, hasHandshake bool
	if h.initialPackets != nil {
		hasInitial = h.initialPackets.history.HasOutstandingPackets()
	}
	if h.handshakePackets != nil {
		hasHandshake = h.handshakePackets.history.HasOutstandingPackets()
	}
	return hasInitial || hasHandshake
}

func (h *sentPacketHandler) hasOutstandingPackets() bool {
	// We only send application data probe packets once the handshake completes,
	// because before that, we don't have the keys to decrypt ACKs sent in 1-RTT packets.
	return (h.handshakeComplete && h.appDataPackets.history.HasOutstandingPackets()) ||
		h.hasOutstandingCryptoPackets()
}

func (h *sentPacketHandler) setLossDetectionTimer() {
	oldAlarm := h.alarm // only needed in case qlog is enabled
	if lossTime, encLevel := h.getEarliestLossTimeAndSpace(); !lossTime.IsZero() {
		// Early retransmit timer or time loss detection.
		h.alarm = lossTime
		if h.qlogger != nil && h.alarm != oldAlarm {
			h.qlogger.SetLossTimer(qlog.TimerTypeACK, encLevel, h.alarm)
		}
		return
	}

	// Cancel the alarm if no packets are outstanding
	if !h.hasOutstandingPackets() && h.peerCompletedAddressValidation {
		h.alarm = time.Time{}
		h.logger.Debugf("Canceling loss detection timer. No packets in flight.")
		if h.qlogger != nil && !oldAlarm.IsZero() {
			h.qlogger.LossTimerCanceled()
		}
		return
	}

	// PTO alarm
	sentTime, encLevel := h.getEarliestSentTimeAndSpace()
	if sentTime.IsZero() {
		if h.peerCompletedAddressValidation {
			panic("didn't expect sentTime to be zero")
		}
		sentTime = time.Now()
	}
	h.alarm = sentTime.Add(h.rttStats.PTO(encLevel == protocol.Encryption1RTT) << h.ptoCount)
	if h.qlogger != nil && h.alarm != oldAlarm {
		h.qlogger.SetLossTimer(qlog.TimerTypePTO, encLevel, h.alarm)
	}
}

func (h *sentPacketHandler) detectAndRemoveLostPackets(now time.Time, encLevel protocol.EncryptionLevel) ([]*Packet, error) {
	pnSpace := h.getPacketNumberSpace(encLevel)
	pnSpace.lossTime = time.Time{}

	maxRTT := float64(utils.MaxDuration(h.rttStats.LatestRTT(), h.rttStats.SmoothedRTT()))
	lossDelay := time.Duration(timeThreshold * maxRTT)

	// Minimum time of granularity before packets are deemed lost.
	lossDelay = utils.MaxDuration(lossDelay, protocol.TimerGranularity)

	// Packets sent before this time are deemed lost.
	lostSendTime := now.Add(-lossDelay)

	var lostPackets []*Packet
	if err := pnSpace.history.Iterate(func(packet *Packet) (bool, error) {
		if packet.PacketNumber > pnSpace.largestAcked {
			return false, nil
		}

		if packet.SendTime.Before(lostSendTime) {
			lostPackets = append(lostPackets, packet)
			if h.qlogger != nil {
				h.qlogger.LostPacket(packet.EncryptionLevel, packet.PacketNumber, qlog.PacketLossTimeThreshold)
			}
		} else if pnSpace.largestAcked >= packet.PacketNumber+packetThreshold {
			lostPackets = append(lostPackets, packet)
			if h.qlogger != nil {
				h.qlogger.LostPacket(packet.EncryptionLevel, packet.PacketNumber, qlog.PacketLossReorderingThreshold)
			}
		} else if pnSpace.lossTime.IsZero() {
			// Note: This conditional is only entered once per call
			lossTime := packet.SendTime.Add(lossDelay)
			if h.logger.Debug() {
				h.logger.Debugf("\tsetting loss timer for packet %d (%s) to %s (in %s)", packet.PacketNumber, encLevel, lossDelay, lossTime)
			}
			pnSpace.lossTime = lossTime
		}
		return true, nil
	}); err != nil {
		return nil, err
	}

	if h.logger.Debug() && len(lostPackets) > 0 {
		pns := make([]protocol.PacketNumber, len(lostPackets))
		for i, p := range lostPackets {
			pns[i] = p.PacketNumber
		}
		h.logger.Debugf("\tlost packets (%d): %d", len(pns), pns)
	}

	for _, p := range lostPackets {
		h.queueFramesForRetransmission(p)
		// the bytes in flight need to be reduced no matter if this packet will be retransmitted
		if p.includedInBytesInFlight {
			h.bytesInFlight -= p.Length
		}
		if err := pnSpace.history.Remove(p.PacketNumber); err != nil {
			return nil, err
		}
		if h.traceCallback != nil {
			frames := make([]wire.Frame, 0, len(p.Frames))
			for _, f := range p.Frames {
				frames = append(frames, f.Frame)
			}
			h.traceCallback(quictrace.Event{
				Time:            now,
				EventType:       quictrace.PacketLost,
				EncryptionLevel: p.EncryptionLevel,
				PacketNumber:    p.PacketNumber,
				PacketSize:      p.Length,
				Frames:          frames,
				TransportState:  h.GetStats(),
			})
		}
	}
	return lostPackets, nil
}

func (h *sentPacketHandler) OnLossDetectionTimeout() error {
	// When all outstanding are acknowledged, the alarm is canceled in
	// setLossDetectionTimer. This doesn't reset the timer in the session though.
	// When OnAlarm is called, we therefore need to make sure that there are
	// actually packets outstanding.
	if h.hasOutstandingPackets() || !h.peerCompletedAddressValidation {
		if err := h.onVerifiedLossDetectionTimeout(); err != nil {
			return err
		}
	}
	h.setLossDetectionTimer()
	return nil
}

func (h *sentPacketHandler) onVerifiedLossDetectionTimeout() error {
	earliestLossTime, encLevel := h.getEarliestLossTimeAndSpace()
	if !earliestLossTime.IsZero() {
		if h.logger.Debug() {
			h.logger.Debugf("Loss detection alarm fired in loss timer mode. Loss time: %s", earliestLossTime)
		}
		if h.qlogger != nil {
			h.qlogger.LossTimerExpired(qlog.TimerTypeACK, encLevel)
		}
		// Early retransmit or time loss detection
		priorInFlight := h.bytesInFlight
		lostPackets, err := h.detectAndRemoveLostPackets(time.Now(), encLevel)
		if err != nil {
			return err
		}
		for _, p := range lostPackets {
			h.congestion.OnPacketLost(p.PacketNumber, p.Length, priorInFlight)
		}
	}

	// PTO
	_, encLevel = h.getEarliestSentTimeAndSpace()
	if h.logger.Debug() {
		h.logger.Debugf("Loss detection alarm for %s fired in PTO mode. PTO count: %d", encLevel, h.ptoCount)
	}
	h.ptoCount++
	if h.qlogger != nil {
		h.qlogger.LossTimerExpired(qlog.TimerTypePTO, encLevel)
		h.qlogger.UpdatedPTOCount(h.ptoCount)
	}
	h.numProbesToSend += 2
	switch encLevel {
	case protocol.EncryptionInitial:
		h.ptoMode = SendPTOInitial
	case protocol.EncryptionHandshake:
		h.ptoMode = SendPTOHandshake
	case protocol.Encryption1RTT:
		h.ptoMode = SendPTOAppData
	default:
		return fmt.Errorf("TPO timer in unexpected encryption level: %s", encLevel)
	}
	return nil
}

func (h *sentPacketHandler) GetLossDetectionTimeout() time.Time {
	return h.alarm
}

func (h *sentPacketHandler) PeekPacketNumber(encLevel protocol.EncryptionLevel) (protocol.PacketNumber, protocol.PacketNumberLen) {
	pnSpace := h.getPacketNumberSpace(encLevel)

	var lowestUnacked protocol.PacketNumber
	if p := pnSpace.history.FirstOutstanding(); p != nil {
		lowestUnacked = p.PacketNumber
	} else {
		lowestUnacked = pnSpace.largestAcked + 1
	}

	pn := pnSpace.pns.Peek()
	return pn, protocol.GetPacketNumberLengthForHeader(pn, lowestUnacked)
}

func (h *sentPacketHandler) PopPacketNumber(encLevel protocol.EncryptionLevel) protocol.PacketNumber {
	return h.getPacketNumberSpace(encLevel).pns.Pop()
}

func (h *sentPacketHandler) SendMode() SendMode {
	numTrackedPackets := h.appDataPackets.history.Len()
	if h.initialPackets != nil {
		numTrackedPackets += h.initialPackets.history.Len()
	}
	if h.handshakePackets != nil {
		numTrackedPackets += h.handshakePackets.history.Len()
	}

	// Don't send any packets if we're keeping track of the maximum number of packets.
	// Note that since MaxOutstandingSentPackets is smaller than MaxTrackedSentPackets,
	// we will stop sending out new data when reaching MaxOutstandingSentPackets,
	// but still allow sending of retransmissions and ACKs.
	if numTrackedPackets >= protocol.MaxTrackedSentPackets {
		if h.logger.Debug() {
			h.logger.Debugf("Limited by the number of tracked packets: tracking %d packets, maximum %d", numTrackedPackets, protocol.MaxTrackedSentPackets)
		}
		return SendNone
	}
	if h.numProbesToSend > 0 {
		return h.ptoMode
	}
	// Only send ACKs if we're congestion limited.
	if !h.congestion.CanSend(h.bytesInFlight) {
		if h.logger.Debug() {
			h.logger.Debugf("Congestion limited: bytes in flight %d, window %d", h.bytesInFlight, h.congestion.GetCongestionWindow())
		}
		return SendAck
	}
	if numTrackedPackets >= protocol.MaxOutstandingSentPackets {
		if h.logger.Debug() {
			h.logger.Debugf("Max outstanding limited: tracking %d packets, maximum: %d", numTrackedPackets, protocol.MaxOutstandingSentPackets)
		}
		return SendAck
	}
	return SendAny
}

func (h *sentPacketHandler) TimeUntilSend() time.Time {
	return h.nextSendTime
}

func (h *sentPacketHandler) ShouldSendNumPackets() int {
	if h.numProbesToSend > 0 {
		// RTO probes should not be paced, but must be sent immediately.
		return h.numProbesToSend
	}
	delay := h.congestion.TimeUntilSend(h.bytesInFlight)
	if delay == 0 || delay > protocol.MinPacingDelay {
		return 1
	}
	return int(math.Ceil(float64(protocol.MinPacingDelay) / float64(delay)))
}

func (h *sentPacketHandler) QueueProbePacket(encLevel protocol.EncryptionLevel) bool {
	pnSpace := h.getPacketNumberSpace(encLevel)
	p := pnSpace.history.FirstOutstanding()
	if p == nil {
		return false
	}
	h.queueFramesForRetransmission(p)
	// TODO: don't remove the packet here
	// Keep track of acknowledged frames instead.
	if p.includedInBytesInFlight {
		h.bytesInFlight -= p.Length
	}
	if err := pnSpace.history.Remove(p.PacketNumber); err != nil {
		// should never happen. We just got this packet from the history.
		panic(err)
	}
	return true
}

func (h *sentPacketHandler) queueFramesForRetransmission(p *Packet) {
	for _, f := range p.Frames {
		f.OnLost(f.Frame)
	}
}

func (h *sentPacketHandler) ResetForRetry() error {
	h.bytesInFlight = 0
	h.initialPackets.history.Iterate(func(p *Packet) (bool, error) {
		h.queueFramesForRetransmission(p)
		return true, nil
	})
	// All application data packets sent at this point are 0-RTT packets.
	// In the case of a Retry, we can assume that the server dropped all of them.
	h.appDataPackets.history.Iterate(func(p *Packet) (bool, error) {
		h.queueFramesForRetransmission(p)
		return true, nil
	})

	h.initialPackets = newPacketNumberSpace(h.initialPackets.pns.Pop())
	h.appDataPackets = newPacketNumberSpace(h.appDataPackets.pns.Pop())
	oldAlarm := h.alarm
	h.alarm = time.Time{}
	if h.qlogger != nil {
		h.qlogger.UpdatedPTOCount(0)
		if !oldAlarm.IsZero() {
			h.qlogger.LossTimerCanceled()
		}
	}
	h.ptoCount = 0
	return nil
}

func (h *sentPacketHandler) SetHandshakeComplete() {
	h.handshakeComplete = true
	// We don't send PTOs for application data packets before the handshake completes.
	// Make sure the timer is armed now, if necessary.
	h.setLossDetectionTimer()
}

func (h *sentPacketHandler) GetStats() *quictrace.TransportState {
	return &quictrace.TransportState{
		MinRTT:           h.rttStats.MinRTT(),
		SmoothedRTT:      h.rttStats.SmoothedRTT(),
		LatestRTT:        h.rttStats.LatestRTT(),
		BytesInFlight:    h.bytesInFlight,
		CongestionWindow: h.congestion.GetCongestionWindow(),
		InSlowStart:      h.congestion.InSlowStart(),
		InRecovery:       h.congestion.InRecovery(),
	}
}
