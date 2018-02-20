FROM centos:latest

ARG BINARY=./metrics

RUN mkdir /metrics-api

ADD ${BINARY} /metrics-api/metrics-api

CMD ["/metrics-api/metrics-api"]