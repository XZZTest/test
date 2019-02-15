FROM xiezhunzhi/alpine

COPY server /server

ENTRYPOINT ["/server"]
