FROM golang:alpine

WORKDIR /app

COPY server ./server

RUN chmod +x /app/server

EXPOSE 15000

CMD ["/app/server"]