FROM golang:latest AS builder

# Output location for make build
RUN mkdir -p /go/src/github.com/aerogear/aerogear-metrics-api/
ENV BINARY=/opt/metrics

COPY . /go/src/github.com/aerogear/aerogear-metrics-api/

WORKDIR /go/src/github.com/aerogear/aerogear-metrics-api/

RUN go get github.com/golang/dep/cmd/dep
RUN make setup
RUN make build_binary_linux

FROM centos:latest AS local
COPY --from=builder /opt/metrics /opt/metrics
RUN chmod +x /opt/metrics
CMD ["/opt/metrics"]
