package view_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/baosen/mastodon_view/view"
)

func TestRootPath(t *testing.T) {
	// Setup server.
	const port = ":8081"
	go func() {
		view.Serve("viewtest", port)
	}()

	// Test if you're able to get / from the server.
	var err error
	retries := 0
	for ; retries < 30; retries++ {
		resp, err := http.Get(fmt.Sprintf("http://localhost%s", port))
		if err != nil {
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		defer resp.Body.Close()

		// TODO: Assert <body>.

		break
	}
	if 30 <= retries {
		t.Fatalf("failed to connect to server: %v", err)
	}
}
