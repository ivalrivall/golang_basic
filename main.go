// Package main implements a simple REST API for a cashier system.
// It provides CRUD operations for managing products in an in-memory store.
// This is a basic implementation intended for learning purposes.
//
// API Endpoints:
//   - GET /health - Health check endpoint
//   - GET /api/produk - Retrieve all products
//   - POST /api/produk - Create a new product
//   - GET /api/produk/{id} - Retrieve a product by ID
//   - PUT /api/produk/{id} - Update a product by ID
//   - DELETE /api/produk/{id} - Delete a product by ID
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Produk represents a product in the cashier system.
// It contains the basic information needed to manage inventory and pricing.
type Produk struct {
	ID    int    `json:"id"`    // Unique identifier for the product
	Nama  string `json:"nama"`  // Product name in Indonesian
	Harga int    `json:"harga"` // Price in Indonesian Rupiah (IDR)
	Stok  int    `json:"stok"`  // Available stock quantity
}

// produk holds all products in memory.
// Note: This is temporary in-memory storage. In a production system,
// this should be replaced with a proper database for persistence and concurrency.
var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

// main sets up HTTP routes and starts the server.
// It defines endpoints for health checks, product listing, creation, retrieval, update, and deletion.
// The server listens on port 8080.
func main() {
	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	// Route pattern with trailing slash to handle ID-based operations
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	// Handle both listing all products (GET) and creating new products (POST)
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)

		} else if r.Method == "POST" {
			// Parse product data from JSON request body
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// Assign auto-incrementing ID and add to in-memory store
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

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
