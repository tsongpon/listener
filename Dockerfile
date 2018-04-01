FROM golang

ADD . /go/src/github.com/tsongpon/listener

RUN go get -u github.com/Masterminds/glide
WORKDIR /go/src/github.com/tsongpon/listener

RUN glide install
RUN go install

ENV GO_ENV_PORT=5000
EXPOSE $GO_ENV_PORT

ENTRYPOINT /go/bin/listener