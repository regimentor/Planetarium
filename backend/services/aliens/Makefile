table = need_argument_table

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://admin:admin@localhost:5432/aliens?sslmode=disable

docker-build:
	 docker build -t planetarium/aliens .

protobuf-compile:
	protoc -I=. --go_out=. --go-grpc_out=. ./protobuf/*.proto

migrate-up:
	goose -dir=db/migration up

# make migrate-create table=TableName
migrate-create:
	goose -dir=db/migration create $(table) sql