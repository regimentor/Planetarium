FROM golang:buster


COPY . /app

WORKDIR /app


RUN go build -o ./build/main ./cmd/aliens

ENTRYPOINT ["/app/build/main"]