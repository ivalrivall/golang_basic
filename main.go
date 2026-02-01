// Package main mengimplementasikan REST API sederhana untuk sistem kasir.
// Aplikasi ini menyediakan operasi CRUD untuk produk dan kategori menggunakan
// database PostgreSQL sebagai penyimpanan data.
//
// Fitur Utama:
//   - Desain RESTful menggunakan metode HTTP standar
//   - Request/response berbasis JSON
//   - Pengambilan parameter ID dari URL
//   - Status code dan error response yang sesuai
//   - Integrasi PostgreSQL untuk persistensi data
//
// Endpoint API:
//
// Health Check:
//   - GET /health - Mengembalikan status dasar API
//
// Manajemen Produk (/api/product):
//   - GET /api/product - Ambil semua produk
//   - POST /api/product - Buat produk baru (ID otomatis)
//   - GET /api/product/{id} - Ambil produk berdasarkan ID
//   - PUT /api/product/{id} - Perbarui produk berdasarkan ID
//   - DELETE /api/product/{id} - Hapus produk berdasarkan ID
//
// Manajemen Kategori (/api/categories):
//   - GET /api/categories - Ambil semua kategori
//   - POST /api/categories - Buat kategori baru (ID otomatis)
//   - GET /api/categories/{id} - Ambil kategori berdasarkan ID
//   - PUT /api/categories/{id} - Perbarui kategori berdasarkan ID
//   - DELETE /api/categories/{id} - Hapus kategori berdasarkan ID
//
// Contoh penggunaan:
//
//	Jalankan server: go run main.go
//	Server listen di http://localhost:8080
//	Gunakan curl, Postman, atau browser untuk menguji endpoint
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

// main menginisialisasi HTTP server dan menyiapkan seluruh routing API.
// Fungsi ini menjadi entry point aplikasi dan mengonfigurasi:
//   - Endpoint health check untuk monitoring
//   - Endpoint manajemen produk (CRUD)
//   - Endpoint manajemen kategori (CRUD)
//   - HTTP server yang berjalan pada port 8080
//
// Routing menggunakan http.ServeMux bawaan Go.
// Route dengan trailing slash menangani operasi berbasis ID, sedangkan
// route tanpa slash menangani operasi koleksi (list/create).

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

	// Route kategori berbasis ID - menangani GET/PUT/DELETE kategori tertentu
	// Pola: /api/categories/{id} - contoh /api/categories/1
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Endpoint health check - mengembalikan status dasar API
	// Contoh: GET http://localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Menjalankan HTTP server pada port 8080
	// Pemanggilan ini bersifat blocking sampai proses dihentikan
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}
