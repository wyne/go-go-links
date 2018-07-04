FROM golang:1.8

RUN go get github.com/pilu/fresh

ADD . /go/src/api
WORKDIR /go/src/api

RUN go build api.go
CMD ["./api"]

EXPOSE 8080

