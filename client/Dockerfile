FROM golang:1.16.0-stretch

# Set the Current Working Directory
WORKDIR /go/src/client-cli/
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0
# Copy go mod and sum files
COPY client/go.mod .
COPY client/go.sum .

# Download all dependancies.
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY client/ .

#CMD ["tail", "-f", "/dev/null"]
#CMD go run main.go answer --id=1