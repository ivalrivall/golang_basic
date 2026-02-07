package dto

type ProductCategoryResponse struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductResponse struct {
	ID       int                      `json:"id"`
	Name     string                   `json:"name"`
	Price    int                      `json:"price"`
	Stock    int                      `json:"stock"`
	Category *ProductCategoryResponse `json:"category,omitempty"`
}
