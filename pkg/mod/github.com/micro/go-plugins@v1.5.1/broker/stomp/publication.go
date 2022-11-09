package stomp

import (
	"github.com/go-stomp/stomp"
	"github.com/micro/go-micro/broker"
)

type publication struct {
	// msg is the actual STOMP message
	msg *stomp.Message
	// m is the broker message
	m *broker.Message
	// Link to the broken (for ack)
	broker *rbroker
	// Topic
	topic string
}

func (p *publication) Ack() error {
	return p.broker.stompConn.Ack(p.msg)
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.m
}
