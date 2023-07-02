FROM golang:buster


COPY . /app

WORKDIR /app


RUN go build -o ./build/main ./cmd/gateway

ENTRYPOINT ["/app/build/main"]