package dto

type ReportRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ReportResponse struct {
	TotalRevenue      int                  `json:"total_revenue"`
	TotalTransactions int                  `json:"total_transactions"`
	MostSoldProduct   *ProductSoldResponse `json:"most_sold_product"`
}

type ProductSoldResponse struct {
	Name         string `json:"name"`
	QuantitySold int    `json:"quantity_sold"`
}
