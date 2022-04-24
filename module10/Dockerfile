FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLE=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
#RUN GOOS=linux go build -installsuffix cgo -o http_server http_server.go
RUN go mod tidy && \
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o . *.go

FROM scratch
#FROM busybox
MAINTAINER "lostar01"
COPY --from=build /httpserver/http_server_metrics /app/http_server_metrics

EXPOSE 8080
WORKDIR /app
ENTRYPOINT [ "/app/http_server_metrics" ]
