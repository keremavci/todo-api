package api

import (
	"net/http"
	"testing"
)

func TestGetStatus(t *testing.T) {
	server:=SetupTest()

	resp, err := http.Get(server.URL + "/api/status")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non 200 response %d", resp.StatusCode)
	}
}