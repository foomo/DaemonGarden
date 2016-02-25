FROM scratch

COPY bin/daemon-garden /usr/sbin/daemon-garden

RUN mkdir -p /var/log/daemonGarden

ENV DAEMON_GARDEN_ADDR=0.0.0.0:8081
ENV DAEMON_GARDEN_LOG_DIR=/var/log/daemonGarden

ENTRYPOINT ["/usr/sbin/daemon-garden"]

CMD ["-address=$DAEMON_GARDEN_ADDR", "-logDir=$DAEMON_GARDEN_LOG_DIR"]
