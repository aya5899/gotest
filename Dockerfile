FROM golang:1.17.5

ENV test=TEST
ENV test2=test2
WORKDIR /go/src/github.com/aya5899/gotest
COPY . .
ENTRYPOINT go test -v ./...