FROM scratch
COPY mastodon_puller /mastodon_puller
ENTRYPOINT ["/mastodon_puller"]
