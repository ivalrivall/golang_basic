# Golang REST API - Cashier System

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

This repository contains a REST API implementation for a cashier system built with Go. It demonstrates building CRUD operations for managing products and categories using Go's standard HTTP package with in-memory storage. Perfect for learning Go web development and REST API design patterns.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Code](#running-the-code)
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

## Topics Covered

The `main.go` file implements a REST API with the following features:

### API Endpoints

**Health Check:**
- `GET /health` - API health check endpoint

**Product Management:**
- `GET /api/produk` - Retrieve all products
- `POST /api/produk` - Create a new product
- `GET /api/produk/{id}` - Retrieve a product by ID
- `PUT /api/produk/{id}` - Update a product by ID
- `DELETE /api/produk/{id}` - Delete a product by ID

**Category Management:**
- `GET /api/categories` - Retrieve all categories
- `POST /api/categories` - Create a new category
- `GET /api/categories/{id}` - Retrieve a category by ID
- `PUT /api/categories/{id}` - Update a category by ID
- `DELETE /api/categories/{id}` - Delete a category by ID

### Technical Concepts

- Struct definitions with JSON tags
- HTTP routing with Go's net/http package
- JSON encoding/decoding
- URL path parsing and parameter extraction
- In-memory data storage with slices
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
