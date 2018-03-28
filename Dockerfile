# pulling a lightweight version of golang
FROM golang:1.10-alpine
MAINTAINER Songpon Imyen <t.songpon@gmail.com>
RUN apk --update add --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/listener
WORKDIR /go/src/listener

RUN go get listener

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/listener"]

# Document that the service listens on port 9000.
EXPOSE 5000