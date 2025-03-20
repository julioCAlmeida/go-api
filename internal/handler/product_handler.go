package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/julioCAlmeida/go-api/internal/model"
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

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest) //400
		return
	}

	product, err := h.productService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //500
		return
	}

	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound) //404
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h * ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // 400
		return
	}

	createdProduct, err := h.productService.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Comtent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	idStr := chi.URLParam(r, "id")

	id, errId := strconv.Atoi(idStr)
	if errId != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest) //400
		return
	}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // 400
		return
	}

	product.ID = id

	uploadProduct, err := h.productService.Update(id, product)
	if uploadProduct == nil {
		http.Error(w, "Product not found", http.StatusNotFound) // 404
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Comtent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest) //400
		return
	}

	err = h.productService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //500
		return
	}

	w.Header().Set("Comtent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}