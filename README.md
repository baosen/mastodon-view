# Mastodon View

__Mastodon View__ is a service that continuously displays a stream of updates from a Mastodon server.

## Architecture

```plaintext
puller (streaming app) <-> view1 (backend-for-frontend) <-> view1's index.html (frontend)
                       <-> view2 (backend-for-frontend) <-> view2's index.html (frontend)
```

## Requirements

This application requires `just` and `go` installed on your computer to build and run it.

## How to run it

Execute `just`, which will run all necessary applications to run __Mastodon View__. Then browse to `localhost:8081` to view the first view (view1) of the app. Browse to `localhost:8082` to view the second view (view2) of the app.
