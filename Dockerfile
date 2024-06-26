FROM gcr.io/distroless/static-debian12

# The service that pulls updates from Mastodon.
COPY puller /puller
COPY .env /.env

# The service that views the updates from the puller.
COPY view1 /view1
COPY view2 /view2
COPY index.html /index.html

# The first argument is used to specify the service.
CMD ["$1"]
