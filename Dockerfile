FROM ghcr.io/pkg6/gruploader:latest

COPY sh/*.sh /sh/
RUN find /sh/ -type f -name "*.sh" -exec chmod +x {} \;

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
