# Golang REST API - Cashier System

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

REST API sistem kasir berbasis Go + PostgreSQL dengan arsitektur berlapis (handler → service → repository). Selain CRUD produk dan kategori, project ini sudah mendukung checkout transaksi multi-item dan endpoint laporan penjualan.

## Fitur

- Health check endpoint
- CRUD Category
- CRUD Product (dengan relasi category)
- Filter product by name (`GET /api/product?name=...`)
- Checkout transaksi multi item (`POST /api/checkout`)
- Laporan transaksi harian (`GET /api/report/today`)
- Laporan transaksi rentang tanggal (`GET /api/report?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`)

## Prerequisites

- Go 1.25+ 
- PostgreSQL

## Installation

```bash
git clone https://github.com/ivalrivall/golang_basic.git
cd golang_basic
go mod tidy
```

## Configuration

Copy env template:

```bash
cp .env.example .env
```

Isi `.env` sesuai environment kamu:

```env
PORT=8080
DB_USER=postgres
DB_PASS=password123
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
DB_SSLMODE=require
DB_MAX_OPEN_CONNS=10
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME_MS=300000
```

## Database Setup (PostgreSQL)

Contoh minimal schema:

```sql
CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  price INT NOT NULL,
  stock INT NOT NULL,
  category_id INT REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  total_amount INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transaction_details (
  id SERIAL PRIMARY KEY,
  transaction_id INT NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
  product_id INT NOT NULL REFERENCES products(id),
  quantity INT NOT NULL,
  subtotal INT NOT NULL
);
```

## Running the API

```bash
go run main.go
```

Hot reload (opsional):

```bash
go install github.com/air-verse/air@latest
air
```

## API Endpoints

### Health
- `GET /health`

### Categories
- `GET /api/categories`
- `POST /api/categories`
- `GET /api/categories/{id}`
- `PUT /api/categories/{id}`
- `DELETE /api/categories/{id}`

### Products
- `GET /api/product`
- `GET /api/product?name=keyword`
- `POST /api/product`
- `GET /api/product/{id}`
- `PUT /api/product/{id}`
- `DELETE /api/product/{id}`

### Transactions
- `POST /api/checkout`

### Reports
- `GET /api/report/today`
- `GET /api/report?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`

## Request Examples

### Health Check

```bash
curl http://localhost:8080/health
```

### Create Category

```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Groceries","description":"Daily needs"}'
```

### Create Product

```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{"name":"Rice","price":12000,"stock":50,"category_id":1}'
```

### Checkout

```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 2, "quantity": 1}
    ]
  }'
```

### Report by Date Range

```bash
curl "http://localhost:8080/api/report?start_date=2026-01-01&end_date=2026-01-31"
```

## Project Structure

```text
.
├── main.go
├── database/
├── dto/
├── handlers/
├── models/
├── repositories/
└── services/
```

## Notes

- Konfigurasi env dibaca menggunakan `viper`.
- Driver PostgreSQL menggunakan `pgx/v5` lewat `database/sql` stdlib adapter.
- `database.InitDB` menerapkan setting connection pool dari env.
