all: test run

build: proto
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo puller.go
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo view1.go
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo view2.go
    docker build . -t mastodon_view

test:
    cd view/ && go test

run: build
    docker compose up

proto:
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mastodon/mastodon.proto
