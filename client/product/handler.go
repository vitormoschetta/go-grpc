package product

import (
	"encoding/json"
	pb "go-grpc-client/protofiles" // Importe o pacote gerado pelo protoc
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductHandler struct {
	client pb.ProductServiceClient
}

func NewProductHandler(productGrpcClient pb.ProductServiceClient) *ProductHandler {
	return &ProductHandler{
		client: productGrpcClient,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createReq pb.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	createResp, err := h.client.CreateProduct(r.Context(), &createReq)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createResp)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/products/"):]
	getReq := pb.GetProductRequest{
		ProductId: productID,
	}
	if err := json.NewDecoder(r.Body).Decode(&getReq); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	getResp, err := h.client.GetProduct(r.Context(), &getReq)
	if err != nil {
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getResp)
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	var emptypb emptypb.Empty
	listResp, err := h.client.ListProducts(r.Context(), &emptypb)
	if err != nil {
		http.Error(w, "Failed to list products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listResp)
}
