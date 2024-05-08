# Mastodon View

__Mastodon View__ is a service that continuously displays a stream of updates from a Mastodon server.

## Architecture

```plaintext
streaming app <-> backend-for-frontend <-> frontend
              <-> backend-for-frontend <-> frontend
```

## Requirements

This application requires `just` and `go` installed on your computer to build and run it.

## How to run it

Execute `just`, which will run all necessary applications to setup __Mastodon View__.

## Building the container image

```
docker build . -t mastodon_puller
```

## Running the container image

```
docker run --rm -ti mastodon_puller
```
