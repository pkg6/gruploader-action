FROM golang:1.23-alpine AS builder

ADD gruploader /app


RUN cd /app && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go build -v -trimpath -ldflags "-s -w -buildid=" -o gruploader && \
    ./gruploader --help


FROM ubuntu:24.10

COPY --from=builder /app/gruploader /usr/local/bin/gruploader
RUN chmod +x /usr/local/bin/gruploader
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl wget jq zip unzip \
    && rm -rf /var/lib/apt/lists/*

COPY entrypoint.sh /entrypoint.sh
COPY sh/*.sh /sh/

ENTRYPOINT ["/entrypoint.sh"]