FROM scratch

COPY mastodon_puller /mastodon_puller

COPY mastodon_view1 /mastodon_view1
COPY index.html /index.html

CMD ["$1"]
