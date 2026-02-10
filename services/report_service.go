package service

import (
	"errors"
	"time"

	"kasir.api/domain"
	"kasir.api/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetSalesReport(startDate, endDate string) (domain.SalesReport, error) {
	return s.repo.GetSalesReport(startDate, endDate)

	// validasi format tanggal
	_, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return domain.SalesReport{}, errors.New("invalid start_date format, use YYYY-MM-DD")
	}

	_, err = time.Parse("2006-01-02", endDate)
	if err != nil {
		return domain.SalesReport{}, errors.New("invalid end_date format, use YYYY-MM-DD")
	}

	return s.repo.GetSalesReport(startDate, endDate)
}
