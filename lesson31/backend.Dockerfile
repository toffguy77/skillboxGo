FROM golang:latest AS build
WORKDIR /http-server
COPY . .
ENV GO111MODULE "on"
RUN go build -o /out/http_server cmd/server/main.go
RUN chmod +x /out/http_server

FROM alpine:latest AS bin
COPY --from=build /out/http-server /server/http-server
EXPOSE 8080
CMD ["/server/http-server"]