package product

import (
	"context"
	"fmt"
	"sync"

	pb "go-grpc-server/protofiles" // Importe o pacote gerado pelo protoc

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

// repositorio de produtos
var products map[string]*pb.Product

// ProductService é a estrutura que implementa a interface gerada pelo protoc
type ProductService struct {
	pb.UnimplementedProductServiceServer            // Necessario para manter a compatibilidade com a interface gerada pelo protoc
	mu                                   sync.Mutex // Mutex para sincronizar o acesso ao mapa
}

func NewProductService() *ProductService {
	products = make(map[string]*pb.Product)
	return &ProductService{}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	productID := uuid.New().String()
	product := &pb.Product{
		ProductId: productID,
		Name:      req.Name,
		Price:     req.Price,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	products[productID] = product
	fmt.Print("Product created: ", product)
	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	product, ok := products[req.ProductId]
	if !ok {
		fmt.Print("Product not found")
		return nil, nil
	}
	return product, nil
}

func (s *ProductService) ListProducts(context.Context, *emptypb.Empty) (*pb.ListProductsResponse, error) {
	fmt.Print("Listing products")
	var productsList []*pb.Product
	for _, product := range products {
		productsList = append(productsList, product)
	}
	return &pb.ListProductsResponse{Products: productsList}, nil
}
