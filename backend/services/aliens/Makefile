
docker-build:
	 docker build -t planetarium/aliens .

protobuf-compile:
	protoc -I=. --go_out=. --go-grpc_out=. ./protobuf/*.proto

migrate-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgres://admin:admin@localhost:5432/aliens?sslmode=disable goose -dir=db/migration up