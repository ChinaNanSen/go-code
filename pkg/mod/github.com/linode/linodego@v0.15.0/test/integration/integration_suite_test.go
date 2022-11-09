package integration

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"
	"k8s.io/client-go/transport"
)

var testingMode = recorder.ModeDisabled
var debugAPI = false
var validTestAPIKey = "NOTANAPIKEY"

var testingPollDuration = time.Duration(15000)
var testingMaxRetryTime = time.Duration(30) * time.Second

func init() {
	if apiToken, ok := os.LookupEnv("LINODE_TOKEN"); ok {
		validTestAPIKey = apiToken
	}

	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			debugAPI = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", debugAPI)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}

	if envFixtureMode, ok := os.LookupEnv("LINODE_FIXTURE_MODE"); ok {
		if envFixtureMode == "record" {
			log.Printf("[INFO] LINODE_FIXTURE_MODE %s will be used for tests", envFixtureMode)
			testingMode = recorder.ModeRecording
		} else if envFixtureMode == "play" {
			log.Printf("[INFO] LINODE_FIXTURE_MODE %s will be used for tests", envFixtureMode)
			testingMode = recorder.ModeReplaying
			testingPollDuration = 1
			testingMaxRetryTime = time.Duration(1) * time.Nanosecond
		}
	}
}

// testRecorder returns a go-vcr recorder and an associated function that the caller must defer
func testRecorder(t *testing.T, fixturesYaml string, testingMode recorder.Mode, realTransport http.RoundTripper) (r *recorder.Recorder, recordStopper func()) {
	if t != nil {
		t.Helper()
	}

	r, err := recorder.NewAsMode(fixturesYaml, testingMode, realTransport)
	if err != nil {
		log.Fatalln(err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	recordStopper = func() {
		r.Stop()
	}
	return
}

// createTestClient is a testing helper to creates a linodego.Client initialized using
// environment variables and configured to record or playback testing fixtures.
// The returned function should be deferred by the caller to ensure the fixture
// recording is properly closed.
func createTestClient(t *testing.T, fixturesYaml string) (*linodego.Client, func()) {
	var (
		c      linodego.Client
		apiKey *string
	)
	if t != nil {
		t.Helper()
	}

	apiKey = &validTestAPIKey

	var recordStopper func()
	var r http.RoundTripper

	if len(fixturesYaml) > 0 {
		r, recordStopper = testRecorder(t, fixturesYaml, testingMode, nil)
	} else {
		r = nil
		recordStopper = func() {}
	}

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *apiKey})
	oc := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
			Base:   r,
		},
	}

	c = linodego.NewClient(oc)
	c.SetDebug(debugAPI).
		SetPollDelay(testingPollDuration).
		SetRetryMaxWaitTime(testingMaxRetryTime)

	return &c, recordStopper
}

// transportRecordWrapper returns a tranport.WrapperFunc which provides the test
// recorder as an http.RoundTripper.
func transportRecorderWrapper(t *testing.T, fixtureYaml string) (transport.WrapperFunc, func()) {
	t.Helper()

	rec, teardown := testRecorder(t, fixtureYaml, testingMode, nil)
	return func(r http.RoundTripper) http.RoundTripper {
		rec.SetTransport(r)
		return rec
	}, teardown
}
