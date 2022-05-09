FROM golang:1.16-alpine As builder
# The latest alpine images don't have some tools like (`git` and `bash`).
RUN apk update
RUN	apk add --no-cache git ca-certificates tzdata bash

# Set the time zone
RUN cp /usr/share/zoneinfo/Africa/Tunis /etc/localtime
RUN echo "Africa/Tunis" >  /etc/timezone

# Set the Current Working Directory
WORKDIR /go/src/fast-track-test/

# Copy go mod and sum files
COPY server/go.mod .
COPY server/go.sum .

# Download all dependancies.
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY server/ .

RUN go get -v github.com/pilu/fresh


ENTRYPOINT fresh