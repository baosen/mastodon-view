package view_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/baosen/mastodon_view/view"
)

func TestRootPath(t *testing.T) {
	// Setup server.
	go func() {
		view.Serve()
	}()

	// Test if you're able to get / from the server.
	var err error
	retries := 0
	for ; retries < 30; retries++ {
		resp, err := http.Get("http://localhost:8081")
		if err != nil {
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		defer resp.Body.Close()
		break
	}
	if 30 <= retries {
		t.Fatalf("failed to connect to server: %v", err)
	}
}
