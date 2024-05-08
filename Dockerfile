FROM scratch

COPY mastodon_puller /mastodon_puller
COPY mastodon_view1 /mastodon_view1

CMD ["$1"]
