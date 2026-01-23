// Package main implements a comprehensive REST API for a cashier system using Go's standard library.
// This implementation demonstrates building a full-featured web API with CRUD operations for both products and categories.
// All data is stored in memory using Go slices, making it suitable for learning and development purposes.
// In a production environment, this would be replaced with a proper database system.
//
// Key Features:
//   - RESTful API design following standard HTTP methods
//   - JSON-based request/response handling
//   - URL path parameter extraction for resource-specific operations
//   - Proper HTTP status codes and error responses
//   - In-memory data persistence (not suitable for production)
//
// API Endpoints:
//
// Health Check:
//   - GET /health - Returns API status and basic information
//
// Product Management (/api/produk):
//   - GET /api/produk - Retrieve all products from inventory
//   - POST /api/produk - Create a new product with auto-incremented ID
//   - GET /api/produk/{id} - Retrieve a specific product by its ID
//   - PUT /api/produk/{id} - Update an existing product's information
//   - DELETE /api/produk/{id} - Remove a product from inventory
//
// Category Management (/api/categories):
//   - GET /api/categories - Retrieve all product categories
//   - POST /api/categories - Create a new category with auto-incremented ID
//   - GET /api/categories/{id} - Retrieve a specific category by its ID
//   - PUT /api/categories/{id} - Update an existing category's information
//   - DELETE /api/categories/{id} - Remove a category from the system
//
// Usage Example:
//
//	Start the server: go run main.go
//	Server listens on http://localhost:8080
//	Use tools like curl, Postman, or browser to test endpoints
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Produk represents a product entity in the cashier system.
// This struct defines the data model for inventory items and includes all necessary
// fields for managing products in a retail environment. The struct uses JSON tags
// to control serialization/deserialization when communicating with API clients.
type Produk struct {
	ID    int    `json:"id"`    // Unique identifier for the product (auto-incremented)
	Nama  string `json:"nama"`  // Product name in Indonesian language
	Harga int    `json:"harga"` // Product price in Indonesian Rupiah (IDR)
	Stok  int    `json:"stok"`  // Current stock quantity available for sale
}

// Category represents a product category in the cashier system.
// Categories are used to organize and group products for better inventory management
// and user experience. Each category has a unique identifier, name, and description.
type Category struct {
	ID          int    `json:"id"`          // Unique identifier for the category (auto-incremented)
	Name        string `json:"name"`        // Category name (e.g., "Makanan", "Minuman")
	Description string `json:"description"` // Detailed description of what this category contains
}

// produk holds all products in memory.
// Note: This is temporary in-memory storage. In a production system,
// this should be replaced with a proper database for persistence and concurrency.
var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan dalam kemasan"},
	{ID: 2, Name: "Minuman", Description: "Minuman dalam kemasan"},
}

// main initializes the HTTP server and sets up all API routes.
// This function is the entry point of the application and configures:
//   - Health check endpoint for monitoring
//   - Product management endpoints (CRUD operations)
//   - Category management endpoints (CRUD operations)
//   - HTTP server listening on port 8080
//
// The routing uses Go's built-in http.ServeMux for pattern matching.
// Routes with trailing slashes handle ID-specific operations, while routes
// without slashes handle collection-level operations (list/create).
func main() {
	// Health check endpoint - returns basic API status
	// Usage: GET http://localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Product routes with ID parameter - handles GET/PUT/DELETE for specific products
	// Pattern: /api/produk/{id} - matches URLs like /api/produk/1, /api/produk/2
	// Routes HTTP methods to appropriate handler functions
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProdukByID(w, r)
		case "PUT":
			updateProduk(w, r)
		case "DELETE":
			deleteProduk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Product collection routes - handles GET (list all) and POST (create new)
	// Pattern: /api/produk - exact match for collection operations
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Return all products as JSON array
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		case "POST":
			// Parse JSON from request body into new product struct
			var produkBaru Produk
			if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
				http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
				return
			}

			// Generate new ID and add to in-memory storage
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			// Return created product with 201 Created status
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Category routes with ID parameter - handles GET/PUT/DELETE for specific categories
	// Pattern: /api/categories/{id} - matches URLs like /api/categories/1, /api/categories/2
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCategoryByID(w, r)
		case "PUT":
			updateCategory(w, r)
		case "DELETE":
			deleteCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Category collection routes - handles GET (list all) and POST (create new)
	// Pattern: /api/categories - exact match for collection operations
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Return all categories as JSON array
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		case "POST":
			// Parse JSON from request body into new category struct
			var categoryBaru Category
			if err := json.NewDecoder(r.Body).Decode(&categoryBaru); err != nil {
				http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
				return
			}

			// Generate new ID and add to in-memory storage
			categoryBaru.ID = len(categories) + 1
			categories = append(categories, categoryBaru)

			// Return created category with 201 Created status
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(categoryBaru)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start HTTP server on port 8080
	// This is a blocking call that runs indefinitely until interrupted
	http.ListenAndServe(":8080", nil)
}

// getProdukByID retrieves a single product by its ID from the URL path.
// It parses the ID from /api/produk/{id}, searches the in-memory store,
// and returns the product as JSON if found, or a 404 error if not found.
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/produk/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// Cari produk dengan ID tersebut
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// getCategoryByID retrieves a single product by its ID from the URL path.
// It parses the ID from /api/categories/{id}, searches the in-memory store,
// and returns the product as JSON if found, or a 404 error if not found.
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// Cari category dengan ID tersebut
	for _, p := range categories {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Category belum ada", http.StatusNotFound)
}

// updateProduk updates an existing product by its ID.
// It parses the ID from /api/produk/{id}, reads the updated product data from the request body,
// finds the product in the store, and replaces it with the new data.
// Returns the updated product as JSON if successful, or a 404 error if the product is not found.
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// updateCategory updates an existing product by its ID.
// It parses the ID from /api/categories/{id}, reads the updated product data from the request body,
// finds the product in the store, and replaces it with the new data.
// Returns the updated product as JSON if successful, or a 404 error if the product is not found.
func updateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop categories, cari id, ganti sesuai data dari request
	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}

	http.Error(w, "Category belum ada", http.StatusNotFound)
}

// deleteProduk removes a product from the store by its ID.
// It parses the ID from /api/produk/{id}, finds the product in the slice,
// and creates a new slice without that product using slice operations.
// Returns a success message if deleted, or a 404 error if the product is not found.
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range produk {
		if p.ID == id {
			// Create new slice by appending elements before and after the target index
			// This effectively removes the product at index i from the slice
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// deleteCategory removes a product from the store by its ID.
// It parses the ID from /api/categories/{id}, finds the product in the slice,
// and creates a new slice without that product using slice operations.
// Returns a success message if deleted, or a 404 error if the product is not found.
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// loop categories cari ID, dapet index yang mau dihapus
	for i, p := range categories {
		if p.ID == id {
			// Create new slice by appending elements before and after the target index
			// This effectively removes the product at index i from the slice
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Category belum ada", http.StatusNotFound)
}
