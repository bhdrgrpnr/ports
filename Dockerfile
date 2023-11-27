FROM golang:1.16-alpine

ENV GO111MODULE=auto

ENV WAIT_HOSTS_TIMEOUT=300
ENV WAIT_SLEEP_INTERVAL=5
ENV WAIT_HOST_CONNECT_TIMEOUT=30

RUN apk add curl git build-base bash

RUN curl -Lo /wait https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait && \
  chmod +x /wait

RUN go get github.com/oxequa/realize
