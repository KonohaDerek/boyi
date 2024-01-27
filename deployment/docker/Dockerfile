FROM golang:1.21.6-alpine AS builder
WORKDIR /server
ENV GO111MODULE=on

COPY . .
RUN ls
RUN apk add gcc g++ libc-dev
RUN GOOS=linux GOARCH=amd64 go build -mod=vendor -o main


FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl tzdata && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /server
COPY --from=builder /server/main /server/main

