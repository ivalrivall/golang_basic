package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProducts menangani request collection produk.
// Endpoint:
//   - GET /api/product (ambil semua produk)
//   - POST /api/product (buat produk baru)
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type categoryResponse struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	type productResponse struct {
		ID       int               `json:"id"`
		Name     string            `json:"name"`
		Price    int               `json:"price"`
		Stock    int               `json:"stock"`
		Category *categoryResponse `json:"category,omitempty"`
	}

	responses := make([]productResponse, 0, len(products))
	for _, product := range products {
		item := productResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		if product.Category != nil {
			cat := product.Category
			item.Category = &categoryResponse{
				Name:        cat.Name,
				Description: cat.Description,
			}
		}
		responses = append(responses, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type categoryResponse struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	type productResponse struct {
		ID       int               `json:"id"`
		Name     string            `json:"name"`
		Price    int               `json:"price"`
		Stock    int               `json:"stock"`
		Category *categoryResponse `json:"category,omitempty"`
	}

	response := productResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	if product.Category != nil {
		cat := product.Category
		response.Category = &categoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// HandleProductByID menangani request produk berdasarkan ID.
// Endpoint:
//   - GET /api/product/{id}
//   - PUT /api/product/{id}
//   - DELETE /api/product/{id}
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetByID mengambil produk berdasarkan ID.
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	type categoryResponse struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	type productResponse struct {
		ID       int               `json:"id"`
		Name     string            `json:"name"`
		Price    int               `json:"price"`
		Stock    int               `json:"stock"`
		Category *categoryResponse `json:"category,omitempty"`
	}

	response := productResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	if product.Category != nil {
		cat := product.Category
		response.Category = &categoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	type categoryResponse struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	type productResponse struct {
		ID       int               `json:"id"`
		Name     string            `json:"name"`
		Price    int               `json:"price"`
		Stock    int               `json:"stock"`
		Category *categoryResponse `json:"category,omitempty"`
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if product.CategoryID == 0 {
		if product.Category != nil && product.Category.ID != 0 {
			product.CategoryID = product.Category.ID
		} else {
			existingProduct, err := h.service.GetByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			product.CategoryID = existingProduct.CategoryID
		}
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedProduct, err := h.service.GetByID(product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := productResponse{
		ID:    updatedProduct.ID,
		Name:  updatedProduct.Name,
		Price: updatedProduct.Price,
		Stock: updatedProduct.Stock,
	}

	if updatedProduct.Category != nil {
		cat := updatedProduct.Category
		response.Category = &categoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete menghapus produk berdasarkan ID.
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}
