FROM harbor1-c3-bj.yottachain.net/yt-common/alpine:3
LABEL maintainer="yuanye@yottachain.io"
LABEL desc="dnrpc service"
LABEL src="https://github.com/yottachain/dnrpc.git"

WORKDIR /app
COPY ./yotta-miner-tracker /app/yotta-miner-tracker

EXPOSE 8080

ENTRYPOINT ["/app/yotta-miner-tracker"]
