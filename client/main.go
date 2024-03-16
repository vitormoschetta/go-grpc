package main

import (
	"log"
	"net/http"

	"go-grpc-client/product"

	pb "go-grpc-client/protofiles" // Importe o pacote gerado pelo protoc

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	grpcServerAddress = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpClient := pb.NewProductServiceClient(conn)
	handler := product.NewProductHandler(grpClient)

	r := mux.NewRouter()
	r.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handler.GetProduct).Methods("GET")
	r.HandleFunc("/products", handler.ListProducts).Methods("GET")
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
