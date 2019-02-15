FROM alpine:latest

COPY server /server

ENTRYPOINT ["/server"]
