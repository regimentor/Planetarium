package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/regimentor/planetarium/backend/services/aliens/internal"

	"github.com/regimentor/planetarium/backend/services/aliens/internal/postgres"

	"github.com/joho/godotenv"

	shared "github.com/regimentor/planetarium/backend/services/shared/pkg"

	"github.com/regimentor/planetarium/backend/services/aliens/internal/services"
	"github.com/regimentor/planetarium/backend/services/aliens/pkg/api"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	log.Printf("Hello, Aliens!")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	psqlUser := os.Getenv("DB_USERNAME")
	psqlPass := os.Getenv("DB_PASSWORD")
	psqlDb := os.Getenv("DB_DATABASE")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// TODO: defer listener.Close()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pgCon, err := shared.NewConnection(ctx, psqlUser, psqlPass, psqlDb, "localhost")
	if err != nil {
		log.Fatalf("failed to create connection to database due err: %v", err)
	}

	aliensStorage := postgres.NewAliensStorage(pgCon)
	relationshipsStorage := postgres.NewRelationshipsStorage(pgCon)

	aliensRepository := internal.NewAliensRepository(aliensStorage)
	relationshipsRepository := internal.NewRelationshipsRepository(relationshipsStorage)

	managerService := services.NewManagerService(aliensRepository)
	relationService := services.NewRelationService(relationshipsRepository)

	api.RegisterManagerServiceServer(s, managerService)
	api.RegisterRelationServiceServer(s, relationService)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
