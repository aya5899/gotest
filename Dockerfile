FROM golang:1.17.5

ENV test=TEST 
WORKDIR /go/src/github.com/aya5899/gotest
COPY . .
ENTRYPOINT go test -v ./...