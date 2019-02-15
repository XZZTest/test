FROM docker.io/alpine

COPY server /server

ENTRYPOINT ["/server"]
