FROM alpine

RUN apk update

RUN mkdir -p /run/docker/plugins /mnt/state /tmp/myexampledriver

COPY bin/myexampledriver /bin/myexampledriver

CMD ["/bin/sh"]
