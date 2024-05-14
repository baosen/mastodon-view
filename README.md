# Mastodon View

__Mastodon View__ is a service that continuously displays a stream of updates from a Mastodon server.

## Architecture

The puller reads from a Mastodon server that publishes updates to a subscriber in a view's backend using gRPC which then pushes it further down to the view's frontend to be displayed.

```plaintext
puller (streaming app) <->(gRPC) view1 (backend-for-frontend) <->(websocket) view1's index.html (frontend)
                       <->(gRPC) view2 (backend-for-frontend) <->(websocket) view2's index.html (frontend)
```

The services are built from distroless base images that only contains the necessary stuff (TLS, timezone etc...) to run the Go applications.

## Requirements

This application requires [`just`](https://just.systems/), [`go`](https://go.dev/), [`grpc`](https://grpc.io/docs/languages/go/quickstart/) and docker installed on your computer to build and run it. On Mac, you can use [`colima`](https://github.com/abiosoft/colima).

## How to run it

Make a `.env`-file that contains the credentials to authenticate to _mastodon.social_ at the root of this directory for example:

```bash
MASTODON_CLIENT_ID=your_id
MASTODON_CLIENT_SECRET=your_very_secret_key
MASTODON_ACCESS_TOKEN=your_very_secret_access_token
```

Execute `just`, which will run all necessary applications to run __Mastodon View__. Then browse to `localhost:8081` to view the first view (view1) of the app. Browse to `localhost:8082` to view the second view (view2) of the app.
