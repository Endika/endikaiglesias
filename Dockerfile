FROM ubuntu
MAINTAINER Endika Iglesias

RUN apt-get update \
    && apt-get dist-upgrade -y \
    && apt-get install -y \
    golang-go

ADD . /home/endikaiglesias

EXPOSE 8002

CMD ["go", "run", "/home/endikaiglesias/run.go"]
