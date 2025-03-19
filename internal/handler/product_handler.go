package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julioCAlmeida/go-api/internal/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductRepository(service service.ProductService) ProductHandler{
	return ProductHandler{productService: service}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
}
	json.NewEncoder(w).Encode(products)
}