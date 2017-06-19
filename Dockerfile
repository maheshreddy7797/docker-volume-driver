FROM alpine

RUN apk update

RUN mkdir -p /run/docker/plugins /mnt/state

COPY systemd/ /lib/systemd/system/

COPY bin/myexampledriver /bin/myexampledriver

CMD ["/bin/sh"]
