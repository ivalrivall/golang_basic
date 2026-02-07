package models

type Report struct {
	TotalRevenue      int          `json:"total_revenue"`
	TotalTransactions int          `json:"total_transactions"`
	MostSoldProduct   *ProductSold `json:"most_sold_product"`
}

type ProductSold struct {
	Name         string `json:"name"`
	QuantitySold int    `json:"quantity_sold"`
}
