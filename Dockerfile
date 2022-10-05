FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN mkdir -p ./configs

COPY server ./server
COPY configs/dev.yaml ./configs

RUN chmod +x /app/server

EXPOSE 15000

CMD ["/app/server"]