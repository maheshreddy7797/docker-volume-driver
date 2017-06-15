FROM alpine

ADD . /code

WORKDIR /code

RUN cd /code

CMD ["bin/sh"]
