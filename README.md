# Mastodon View

__Mastodon View__ is a service that continuously displays a stream of updates from a Mastodon server.

## Architecture

```plaintext
puller (streaming app) <-> view1 (bff) <-> view1's index.html (frontend)
                       <-> view2 (bff) <-> view2's index.html (frontend)
```

## Requirements

This application requires `just` and `go` installed on your computer to build and run it.

## How to run it

Execute `just`, which will run all necessary applications to run __Mastodon View__. Then browse to `localhost:8081` to view the app.

### Building the container image

```
just build
```

### Running the container image

```
just run
```
