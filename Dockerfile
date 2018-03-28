# pulling a lightweight version of golang
FROM golang:1.8-alpine
MAINTAINER Songpon Imyen <t.songpon@gmail.com>
RUN apk --update add --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/listener
WORKDIR /go/src/listener

#ENV REDPLANET_DB_HOST localhost
#ENV TOKEN some_token
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get listener

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/listener"]

# Document that the service listens on port 9000.
EXPOSE 5000