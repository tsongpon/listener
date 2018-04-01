FROM golang:1.10-alpine3.7

RUN apk add --no-cache curl
RUN apk add --no-cache git

ADD . /go/src/github.com/tsongpon/listener
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/tsongpon/listener

RUN dep ensure
RUN go install

ENTRYPOINT /go/bin/listener

EXPOSE 5000