FROM alpine

RUN apk update

RUN mkdir -p /run/docker/plugins /mnt/state /tmp/mntdir1

COPY bin/docker-volume-driver /bin/docker-volume-driver

CMD ["/bin/sh"]
