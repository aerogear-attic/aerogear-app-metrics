FROM centos:7.4.1708
ARG BINARY=./aerogear-app-metrics
EXPOSE 3000

COPY ${BINARY} /opt/aerogear-app-metrics
ENTRYPOINT ["/opt/aerogear-app-metrics"]
