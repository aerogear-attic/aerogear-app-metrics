FROM centos:7.4.1708
ARG BINARY=./metrics
EXPOSE 3000

COPY ${BINARY} /opt/metrics
ENTRYPOINT ["/opt/metrics"]
