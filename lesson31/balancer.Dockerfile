FROM golang:1.19 AS builder
RUN apt update & apt upgrade -yy

WORKDIR /usr/src/app

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV CGO_ENABLED=0

COPY . .
RUN go build -o friends-balancer cmd/friends-balancer/main.go

FROM debian:buster-slim
RUN apt update & apt upgrade -yy

COPY --from=builder /usr/src/app/friends-balancer /go/bin/friends-balancer

EXPOSE 8080
ENV PATH="/go/bin:${PATH}"
ENTRYPOINT ["friends-balancer"]
