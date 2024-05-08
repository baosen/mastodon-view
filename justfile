all: run

build: mastodon_proto
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo mastodon_puller.go
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo mastodon_view1.go
    docker build . -t mastodon_view

run: build
    docker compose up

mastodon_puller:
    go run mastodon_puller.go

mastodon_proto:
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mastodon/mastodon.proto

mastodon_view1:
    go run mastodon_view1.go

mastodon_view2:
    go run mastodon_view2.go
