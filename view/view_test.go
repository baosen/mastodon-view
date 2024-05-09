package view_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/baosen/mastodon_view/view"
)

func TestRootPath(t *testing.T) {
	go func() {
		view.Serve()
	}()

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
		fmt.Printf("failed to connect to server: %v", err)
	}
}
