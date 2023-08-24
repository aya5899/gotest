FROM golang:1.17.5

ENV test=TEST
ENV test2=test2
ENV test3=TEST3
ENV test4=TEST4
WORKDIR /go/src/github.com/aya5899/gotest
COPY . .
ENTRYPOINT go test -v ./...