FROM golang:buster


COPY . /app

WORKDIR /app


RUN go build -o ./build/main ./cmd/galaxies

ENTRYPOINT ["/app/build/main"]