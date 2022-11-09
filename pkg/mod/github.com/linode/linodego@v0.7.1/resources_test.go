package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"golang.org/x/oauth2"
)

func TestResourceEndpoint(t *testing.T) {
	apiKey := "MYFAKEAPIKEY"

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})
	oc := oauth2.NewClient(context.Background(), tokenSource)
	client := NewClient(oc)

	r := client.Resource("images")
	e, err := r.Endpoint()
	if err != nil {
		t.Error("Got error when querying for images endpoint")
	}

	if e != imagesEndpoint {
		t.Errorf("Images endpoint did not match '%s'", imagesEndpoint)
	}
}
func TestResourceTemplatedendpointWithID(t *testing.T) {
	apiKey := "MYFAKEAPIKEY"

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})
	oc := oauth2.NewClient(context.Background(), tokenSource)
	client := NewClient(oc)

	backupID := 1234255
	e, err := client.InstanceSnapshots.endpointWithID(backupID)
	if err != nil {
		t.Error("Got error when getting endpoint with id for backups")
	}
	if e != fmt.Sprintf("linode/instances/%d/backups", backupID) {
		t.Errorf("Backups endpoint did not contain backup ID '%d'", backupID)
	}
}

func TestUnmarshalTimeRemaining(t *testing.T) {
	if *unmarshalTimeRemaining(json.RawMessage("\"1:23\"")) != 83 {
		t.Errorf("Error parsing duration style time_remaining")
	}
	if unmarshalTimeRemaining(json.RawMessage("null")) != nil {
		t.Errorf("Error parsing null time_remaining")
	}
	if *unmarshalTimeRemaining(json.RawMessage("0")) != 0 {
		t.Errorf("Error parsing int style time_remaining")
	}
}
