package main

import (
	"log"
	"net"

	"github.com/kenmobility/product-microservice/config"
	"github.com/kenmobility/product-microservice/db"
	"github.com/kenmobility/product-microservice/handlers"
	"github.com/kenmobility/product-microservice/models"
	"github.com/kenmobility/product-microservice/pb"
	"github.com/kenmobility/product-microservice/repository"
	"github.com/kenmobility/product-microservice/service"
	"google.golang.org/grpc"
)

func main() {
	// load env variables
	config := config.LoadConfig("")

	db, err := db.ConnectPostgresDb(*config)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Run migrations
	models.Migrate(db)

	// Set up repository, service, and handler
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register gRPC services
	pb.RegisterProductServiceServer(grpcServer, productHandler)

	log.Println("gRPC server running on port :50051")

	// Start gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
