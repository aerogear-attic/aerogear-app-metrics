FROM golang:latest

RUN mkdir -p /aerogear-metrics-api

ADD . /aerogear-metrics-api

WORKDIR /aerogear-metrics-api

RUN go build aerogear_metrics_api.go

CMD ["./aerogear_metrics_api"]