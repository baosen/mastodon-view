# Mastodon View

__Mastodon View__ is a service that continuously displays a stream of updates from a Mastodon server.

## Architecture

The puller reads from a Mastodon server that pushes updates to the backend's view using gRPC which then pushes it further down to the view's frontend to be displayed.

```plaintext
puller (streaming app) <->(gRPC) view1 (backend-for-frontend) <->(websocket) view1's index.html (frontend)
                       <->(gRPC) view2 (backend-for-frontend) <->(websocket) view2's index.html (frontend)
```

## Requirements

This application requires [`just`](https://just.systems/) and [`go`](https://go.dev/) installed on your computer to build and run it.

## How to run it

Execute `just`, which will run all necessary applications to run __Mastodon View__. Then browse to `localhost:8081` to view the first view (view1) of the app. Browse to `localhost:8082` to view the second view (view2) of the app.
