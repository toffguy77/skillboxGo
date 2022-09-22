FROM golang:alpine AS build
WORKDIR /src
COPY . .
ENV SERVER_PORT "8080"
ENV SERVER_HOST "localhost"
ENV BUILD_NAME "load-babalncer"
ENV GO111MODULE "on"
RUN go build -o /out/$BUILD_NAME cmd/server/main.go

FROM scratch AS bin
COPY --from=build /out/$BUILD_NAME /
EXPOSE 8080
CMD ["go", "run", "bin/$BUILD_NAME"]