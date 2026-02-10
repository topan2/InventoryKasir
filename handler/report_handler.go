package handler

import (
	"encoding/json"
	"net/http"

	"kasir.api/repository"
	service "kasir.api/services"
)

type ReportHandler struct {
	Repo    *repository.ReportRepository
	service *service.ReportService
}

func NewReportHandler(repo *repository.ReportRepository, service *service.ReportService) *ReportHandler {
	return &ReportHandler{Repo: repo, service: service}
}

func (h *ReportHandler) DailySales(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	date := r.URL.Query().Get("date")
	if date == "" {
		http.Error(w, "date is required, format YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	report, _ := h.Repo.GetDailySales(date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) GetSalesReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		http.Error(w, "start_date and end_date are required, format YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	report, err := h.service.GetSalesReport(startDate, endDate)
	if err != nil {
		http.Error(w, "failed to generate report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
