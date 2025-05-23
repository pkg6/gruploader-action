FROM ubuntu:24.10
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl wget jq zip unzip git \
    && rm -rf /var/lib/apt/lists/*
COPY entrypoint.sh /entrypoint.sh
COPY sh/*.sh /sh/
RUN find /sh/ -type f -name "*.sh" -exec chmod +x {} \;
ENTRYPOINT ["/entrypoint.sh"]