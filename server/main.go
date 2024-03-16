package main

import (
	"log"
	"net"

	"go-grpc-server/product"
	pb "go-grpc-server/protofiles" // Importe o pacote gerado pelo protoc

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service := product.NewProductService()
	reflection.Register(server) // necessario para a descoberta de serviços
	pb.RegisterProductServiceServer(server, service)

	log.Println("Server running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
