package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	"time"
)

const reportDateLayout = "2006-01-02"

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetTodayReport() (*models.Report, error) {
	today := time.Now().Format(reportDateLayout)
	startDate := today
	endDate := today
	return repo.GetReport(startDate, endDate)
}

func (repo *ReportRepository) GetReport(startDate string, endDate string) (*models.Report, error) {
	totalRevenue, err := repo.getTotalRevenue(startDate, endDate)
	if err != nil {
		return nil, err
	}

	totalTransactions, err := repo.getTotalTransactions(startDate, endDate)
	if err != nil {
		return nil, err
	}

	mostSoldProductQuantity, mostSoldProductName, err := repo.getMostSoldProduct(startDate, endDate)
	if err != nil {
		return nil, err
	}

	var mostSoldProduct *models.ProductSold
	if mostSoldProductName != "" {
		mostSoldProduct = &models.ProductSold{
			Name:         mostSoldProductName,
			QuantitySold: mostSoldProductQuantity,
		}
	}

	return &models.Report{
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
		MostSoldProduct:   mostSoldProduct,
	}, nil
}

func (repo *ReportRepository) getTotalRevenue(startDate string, endDate string) (int, error) {
	var totalRevenue int
	err := repo.db.QueryRow(
		"SELECT COALESCE(SUM(total_amount), 0) FROM transactions WHERE DATE(created_at) BETWEEN $1 AND $2",
		startDate,
		endDate,
	).Scan(&totalRevenue)
	if err != nil {
		return 0, err
	}

	return totalRevenue, nil
}

func (repo *ReportRepository) getTotalTransactions(startDate string, endDate string) (int, error) {
	var totalTransactions int
	err := repo.db.QueryRow(
		"SELECT COUNT(*) FROM transactions WHERE DATE(created_at) BETWEEN $1 AND $2",
		startDate,
		endDate,
	).Scan(&totalTransactions)
	if err != nil {
		return 0, err
	}

	return totalTransactions, nil
}

func (repo *ReportRepository) getMostSoldProduct(startDate string, endDate string) (int, string, error) {
	var mostSoldProductQuantity sql.NullInt64
	var mostSoldProductName sql.NullString
	err := repo.db.QueryRow(
		`SELECT COALESCE(SUM(td.quantity), 0) AS quantity_sold, p.name
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY td.product_id, p.name
		ORDER BY quantity_sold DESC
		LIMIT 1`,
		startDate,
		endDate,
	).Scan(&mostSoldProductQuantity, &mostSoldProductName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return 0, "", nil
	}

	return int(mostSoldProductQuantity.Int64), mostSoldProductName.String, nil
}
