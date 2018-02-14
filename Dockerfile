FROM golang:latest

RUN mkdir -p /aerogear-metrics-api

ADD . /aerogear-metrics-api

WORKDIR /aerogear-metrics-api

CMD ["./aerogear_metrics_api"]