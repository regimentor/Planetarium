FROM golang:buster


COPY . /app

WORKDIR /app


RUN go build -o ./build/main ./cmd/stars

ENTRYPOINT ["/app/build/main"]