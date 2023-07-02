package main

import (
	"log"
	"net"

	"github.com/regimentor/planetarium/backend/services/aliens/internal/services"
	"github.com/regimentor/planetarium/backend/services/aliens/pkg/api"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	log.Printf("Hello, Aliens!")

	// TODO: defer listener.Close()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	managerService := services.NewManagerService()
	relationService := services.NewRelationService()

	api.RegisterManagerServiceServer(s, managerService)
	api.RegisterRelationServiceServer(s, relationService)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
