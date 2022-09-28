FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN mkdir -p ./config

COPY server ./server
COPY config/dev.yaml ./config

RUN chmod +x /app/server

EXPOSE 15000

CMD ["/app/server"]