package handlers

import (
	"encoding/json"
	"errors"
	"kasir-api/dto"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	report, err := h.service.GetTodayReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var mostSoldProduct *dto.ProductSoldResponse
	if report.MostSoldProduct != nil {
		mostSoldProduct = &dto.ProductSoldResponse{
			Name:         report.MostSoldProduct.Name,
			QuantitySold: report.MostSoldProduct.QuantitySold,
		}
	}

	response := dto.ReportResponse{
		TotalRevenue:      report.TotalRevenue,
		TotalTransactions: report.TotalTransactions,
		MostSoldProduct:   mostSoldProduct,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	var req dto.ReportRequest

	// Support query params for GET requests:
	// /api/report?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
	if r.Method == http.MethodGet {
		req.StartDate = r.URL.Query().Get("start_date")
		req.EndDate = r.URL.Query().Get("end_date")
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	}

	if req.StartDate == "" || req.EndDate == "" {
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	report, err := h.service.GetReport(req.StartDate, req.EndDate)
	if err != nil {
		if errors.Is(err, errors.New("report not found")) || err.Error() == "report not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var mostSoldProduct *dto.ProductSoldResponse
	if report.MostSoldProduct != nil {
		mostSoldProduct = &dto.ProductSoldResponse{
			Name:         report.MostSoldProduct.Name,
			QuantitySold: report.MostSoldProduct.QuantitySold,
		}
	}

	response := dto.ReportResponse{
		TotalRevenue:      report.TotalRevenue,
		TotalTransactions: report.TotalTransactions,
		MostSoldProduct:   mostSoldProduct,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
