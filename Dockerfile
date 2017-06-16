FROM alpine

ADD . /code

WORKDIR /code

COPY systemd/ /lib/systemd/system/

COPY myexampledriver /bin

CMD ["/bin/myexampledriver"]
