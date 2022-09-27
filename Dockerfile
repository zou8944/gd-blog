FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY server /app/

EXPOSE 15000

CMD ["/app/server"]