# Golang Basic Learning

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

This repository contains basic examples and tutorials for learning Go programming language. The code demonstrates fundamental concepts including variable declarations, data types, and basic operations.

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

The `main.go` file demonstrates:

- Variable declarations (manifest typing and type inference)
- Multi-variable declarations
- Underscore variable usage
- Pointer variables with `new()`
- Numeric data types (uint8, int, float64)
- Boolean data type
- String literals and multi-line strings
- Printing with `fmt.Printf` and `fmt.Println`

## Learning Resources

### Tutorials
- [Dasar Pemrograman Golang](https://dasarpemrogramangolang.novalagung.com/)
- [Materi Pertemuan 1](https://docs.kodingworks.io/s/01e57b74-74e6-44df-ac02-7e30a2478528)

### Video Tutorials
- [Pertemuan 1: Introduction to Go](https://www.youtube.com/watch?v=HL1JU206V-4)

### Tasks

#### Pertemuan 1 Checklist

**Basic Concepts:**
- [ ] Understand variable declarations (manifest typing vs type inference)
- [ ] Practice declaring variables with `var` keyword
- [ ] Learn short variable declarations with `:=`
- [ ] Implement multi-variable declarations
- [ ] Use underscore `_` for unused variables
- [ ] Work with pointer variables using `new()`
- [ ] Format numeric types (uint8, int, float64) with `%d` and `%f`
- [ ] Format boolean values with `%t`
- [ ] Practice string literals and multi-line strings

**API Implementation Tasks:**
- [ ] Create Category model with ID, Name, Description fields
- [ ] Implement GET /categories endpoint (get all categories)
- [ ] Implement POST /categories endpoint (add new category)
- [ ] Implement PUT /categories/{id} endpoint (update category)
- [ ] Implement GET /categories/{id} endpoint (get category details)
- [ ] Implement DELETE /categories/{id} endpoint (delete category)
- [ ] Deploy API to cloud platform (Railway or Zeabur)
- [ ] Submit project via form with email, GitHub link, and deployment link
- [ ] Complete exercises: [Pertemuan 1](https://docs.kodingworks.io/s/17137b9a-ed7a-4950-ba9e-eb11299531c2)

## Next Steps

- Constants: https://dasarpemrogramangolang.novalagung.com/A-konstanta.html
