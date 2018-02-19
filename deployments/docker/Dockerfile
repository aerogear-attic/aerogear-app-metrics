FROM golang:latest

# Output location for make build
ENV BINARY=/metrics-api/metrics-api

RUN mkdir -p /go/src/github.com/aerogear/aerogear-metrics-api/
RUN mkdir /metrics-api

COPY . /go/src/github.com/aerogear/aerogear-metrics-api/

WORKDIR /go/src/github.com/aerogear/aerogear-metrics-api/

RUN go get github.com/golang/dep/cmd/dep
RUN make build

CMD ["/metrics-api/metrics-api"]
