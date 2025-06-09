FROM ghcr.io/pkg6/gruploader:0.1.2

COPY sh/*.sh /sh/
RUN find /sh/ -type f -name "*.sh" -exec chmod +x {} \;

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
