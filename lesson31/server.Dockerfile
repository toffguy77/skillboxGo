FROM golang:1.19 AS builder

WORKDIR /usr/src/app

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV CGO_ENABLED=0

COPY . .
RUN go build -o friends-server cmd/friends-server/main.go

FROM debian:buster-slim
COPY --from=builder /usr/src/app/friends-server /go/bin/friends-server
EXPOSE 54321
ENV PATH="/go/bin:${PATH}"
