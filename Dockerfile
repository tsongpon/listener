FROM golang:1.10

ADD . /go/src/github.com/tsongpon/listener
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/tsongpon/listener

RUN dep ensure
RUN go install

ENTRYPOINT /go/bin/listener

EXPOSE 5000