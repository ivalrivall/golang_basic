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
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPass          string `mapstructure:"DB_PASS"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBSSLMode       string `mapstructure:"DB_SSLMODE"`
	DBMaxOpenConns  int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns  int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBConnMaxLifeMs int    `mapstructure:"DB_CONN_MAX_LIFETIME_MS"`
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
		Port:            viper.GetString("PORT"),
		DBUser:          strings.TrimSpace(viper.GetString("DB_USER")),
		DBPass:          strings.TrimSpace(viper.GetString("DB_PASS")),
		DBHost:          strings.TrimSpace(viper.GetString("DB_HOST")),
		DBPort:          strings.TrimSpace(viper.GetString("DB_PORT")),
		DBName:          strings.TrimSpace(viper.GetString("DB_NAME")),
		DBSSLMode:       strings.TrimSpace(viper.GetString("DB_SSLMODE")),
		DBMaxOpenConns:  viper.GetInt("DB_MAX_OPEN_CONNS"),
		DBMaxIdleConns:  viper.GetInt("DB_MAX_IDLE_CONNS"),
		DBConnMaxLifeMs: viper.GetInt("DB_CONN_MAX_LIFETIME_MS"),
	}

	if config.DBSSLMode == "" {
		config.DBSSLMode = "require"
	}

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)

	poolConfig := database.PoolConfig{
		MaxOpenConns:    config.DBMaxOpenConns,
		MaxIdleConns:    config.DBMaxIdleConns,
		ConnMaxLifetime: time.Duration(config.DBConnMaxLifeMs) * time.Millisecond,
	}

	db, err := database.InitDB(dsn, poolConfig)
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

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// Route kategori berbasis ID - menangani GET/PUT/DELETE kategori tertentu
	// Pola: /api/categories/{id} - contoh /api/categories/1
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Endpoint checkout transaksi
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)

	// Endpoint laporan transaksi
	http.HandleFunc("/api/report/today", reportHandler.GetTodayReport)
	http.HandleFunc("/api/report", reportHandler.GetReport)

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
	fmt.Println("Server running di localhost:" + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Failed running server:", err)
	}
}
