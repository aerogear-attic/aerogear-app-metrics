FROM golang:1.10.0 AS builder

# Output location for make build
RUN mkdir -p /go/src/github.com/aerogear/aerogear-metrics-api/
ENV BINARY=/opt/metrics

# must be run before COPY so dep download is cached
# use docker build --no-cache to update
RUN go get github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/aerogear/aerogear-metrics-api/
WORKDIR /go/src/github.com/aerogear/aerogear-metrics-api/

RUN make setup

# Disabling cgo forces static linking of c libs, allowing the binary to run
# in `scratch`
RUN CGO_ENABLED=0 make build_binary_linux
RUN chmod +x /opt/metrics

FROM scratch AS local
EXPOSE 3000
COPY --from=builder /opt/metrics /opt/metrics
ENTRYPOINT ["/opt/metrics"]
