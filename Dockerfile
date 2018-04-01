FROM golang:1.10.1
MAINTAINER Songpon Imyen <t.songpon@gmail.com>

RUN apk --update add --no-cache git

ADD . /go/src/github.com/tsongpon/listener
WORKDIR /go/src/github.com/tsongpon/listener

RUN go get -u github.com/Masterminds/glide

RUN glide install
RUN go install

EXPOSE 5000

ENTRYPOINT /go/bin/listener