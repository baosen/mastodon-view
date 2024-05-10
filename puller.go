package main

import (
	"github.com/baosen/mastodon_view/streamer"
)

// Read a stream from a Mastodon-server and serve it over gRPC.
func main() {
	streamer.Stream(".env")
}
