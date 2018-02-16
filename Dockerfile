FROM golang:latest

RUN mkdir -p /go/src/github.com/aerogear/aerogear-metrics-api/
RUN mkdir /metrics-api

ADD . /go/src/github.com/aerogear/aerogear-metrics-api/

WORKDIR /go/src/github.com/aerogear/aerogear-metrics-api/

RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go build -o /metrics-api/metrics-api ./cmd/metrics-api/metrics-api.go

RUN ls /metrics-api

CMD ["/metrics-api/metrics-api"]