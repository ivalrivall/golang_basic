package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(startDate string, endDate string) (*models.Report, error) {
	return s.repo.GetReport(startDate, endDate)
}

func (s *ReportService) GetTodayReport() (*models.Report, error) {
	return s.repo.GetTodayReport()
}
