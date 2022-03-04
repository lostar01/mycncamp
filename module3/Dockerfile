FROM scratch

MAINTAINER "lostar01"
COPY bin/amd64/http_server /app/http_server

EXPOSE 8080
WORKDIR /app
ENTRYPOINT [ "/app/http_server" ]
