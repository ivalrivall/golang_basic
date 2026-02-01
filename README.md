# Golang REST API - Cashier System

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

This repository contains a REST API implementation for a cashier system built with Go. It demonstrates building CRUD operations for managing products and categories using Go's standard HTTP package with a PostgreSQL database. Perfect for learning Go web development and REST API design patterns.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Code](#running-the-code)
- [Configuration](#configuration)
- [Database Setup](#database-setup)
- [API Endpoints](#api-endpoints)
- [Request/Response Examples](#requestresponse-examples)
- [Data Models](#data-models)
- [Topics Covered](#topics-covered)
- [Learning Resources](#learning-resources)
- [Next Steps](#next-steps)

## Prerequisites

- Go 1.16 or later installed on your system
- Basic understanding of programming concepts

## Installation

Clone the repository:

```bash
git clone https://github.com/ivalrivall/golang_basic.git
cd golang_basic
```

## Running the Code

To run the main program:

```bash
go run main.go
```

### Hot Reload (Air)

Install Air (if you don't have it yet):

```bash
go install github.com/air-verse/air@latest
```

Run the API with hot reloading:

```bash
air
```

## Configuration

Create a `.env` file in the project root:

```bash
PORT=8080
DB_CONN=postgresql://<user>:<password>@<host>:5432/<database>
```

The application loads environment variables using `viper`. If `.env` exists, it will be read automatically.

## Database Setup

This API expects two tables: `products` and `categories`. Example schema (PostgreSQL):

```sql
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  price INT NOT NULL,
  stock INT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL
);
```

## API Endpoints

**Health Check:**
- `GET /health` - API health check endpoint

**Product Management:**
- `GET /api/product` - Retrieve all products
- `POST /api/product` - Create a new product
- `GET /api/product/{id}` - Retrieve a product by ID
- `PUT /api/product/{id}` - Update a product by ID
- `DELETE /api/product/{id}` - Delete a product by ID

**Category Management:**
- `GET /api/categories` - Retrieve all categories
- `POST /api/categories` - Create a new category
- `GET /api/categories/{id}` - Retrieve a category by ID
- `PUT /api/categories/{id}` - Update a category by ID
- `DELETE /api/categories/{id}` - Delete a category by ID

## Request/Response Examples

### Health Check

```bash
curl http://localhost:8080/health
```

```json
{
  "status": "OK",
  "message": "API Running"
}
```

### Create Product

```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{"name":"Rice","price":12000,"stock":50}'
```

```json
{
  "id": 1,
  "name": "Rice",
  "price": 12000,
  "stock": 50
}
```

### Get Product By ID

```bash
curl http://localhost:8080/api/product/1
```

```json
{
  "id": 1,
  "name": "Rice",
  "price": 12000,
  "stock": 50
}
```

### Create Category

```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Groceries","description":"Daily needs"}'
```

```json
{
  "id": 1,
  "name": "Groceries",
  "description": "Daily needs"
}
```

## Data Models

### Product

```json
{
  "id": 1,
  "name": "string",
  "price": 12000,
  "stock": 50
}
```

### Category

```json
{
  "id": 1,
  "name": "string",
  "description": "string"
}
```

## Topics Covered

The `main.go` file implements a REST API with the following features:

### Technical Concepts

- Struct definitions with JSON tags
- HTTP routing with Go's net/http package
- JSON encoding/decoding
- URL path parsing and parameter extraction
- Repository/service layering
- PostgreSQL persistence with `database/sql`
- CRUD operations implementation
- Error handling and HTTP status codes

## Learning Resources

### Tutorials
- [Dasar Pemrograman Golang](https://dasarpemrogramangolang.novalagung.com/)
- [Materi Pertemuan 1](https://docs.kodingworks.io/s/01e57b74-74e6-44df-ac02-7e30a2478528)

### Video Tutorials
- [Pertemuan 1: Introduction to Go](https://www.youtube.com/watch?v=HL1JU206V-4)

### Tasks

#### Pertemuan 1 Checklist

**API Implementation Tasks:**
- [x] Create Category model with ID, Name, Description fields
- [x] Implement GET /categories endpoint (get all categories)
- [x] Implement POST /categories endpoint (add new category)
- [x] Implement PUT /categories/{id} endpoint (update category)
- [x] Implement GET /categories/{id} endpoint (get category details)
- [x] Implement DELETE /categories/{id} endpoint (delete category)
- [x] Implement Product model with ID, Nama, Harga, Stok fields
- [x] Implement GET /api/produk endpoint (get all products)
- [x] Implement POST /api/produk endpoint (add new product)
- [x] Implement PUT /api/produk/{id} endpoint (update product)
- [x] Implement GET /api/produk/{id} endpoint (get product details)
- [x] Implement DELETE /api/produk/{id} endpoint (delete product)
- [x] Implement GET /health endpoint (health check)
- [x] Use JSON encoding/decoding for request/response handling
- [x] Handle HTTP methods (GET, POST, PUT, DELETE) appropriately
- [x] Implement URL path parsing for ID-based operations
- [x] Add proper HTTP status codes and error responses
- [ ] Deploy API to cloud platform (Railway or Zeabur)
- [ ] Submit project via form with email, GitHub link, and deployment link
- [ ] Complete exercises: [Pertemuan 1](https://docs.kodingworks.io/s/17137b9a-ed7a-4950-ba9e-eb11299531c2)

## Next Steps

- Constants: https://dasarpemrogramangolang.novalagung.com/A-konstanta.html
