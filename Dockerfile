FROM golang

ADD . /go/src/github.com/tsongpon/listener

RUN go get -u github.com/Masterminds/glide
WORKDIR /go/src/github.com/tsongpon/listener

RUN glide install
RUN go install

ENTRYPOINT /go/bin/listener

EXPOSE 5000