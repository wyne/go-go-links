FROM golang:1.8

RUN go get github.com/pilu/fresh
RUN go get github.com/go-redis/redis

ADD . /go/src/api
WORKDIR /go/src/api

RUN go build api.go
CMD ["./api"]

EXPOSE 8080

