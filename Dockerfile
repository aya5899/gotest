FROM golang:1.17.5

WORKDIR /go/src/github.com/aya5899/gotest
COPY . .
ENTRYPOINT go test -v ./...