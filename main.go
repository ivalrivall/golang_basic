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
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
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
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// Category routes with ID parameter - handles GET/PUT/DELETE for specific categories
	// Pattern: /api/categories/{id} - matches URLs like /api/categories/1, /api/categories/2
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Health check endpoint - returns basic API status
	// Usage: GET http://localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Start HTTP server on port 8080
	// This is a blocking call that runs indefinitely until interrupted
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}
