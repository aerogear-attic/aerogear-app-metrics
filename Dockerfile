FROM centos:7
ARG BINARY=./aerogear-app-metrics
EXPOSE 3000

COPY ${BINARY} /opt/aerogear-app-metrics
ENTRYPOINT ["/opt/aerogear-app-metrics"]
