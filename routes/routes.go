package routes

import (
	"net/http"

	"kasir.api/handler"
)

func RegisterRoutes(h *handler.InventoryHandler, reportHandler *handler.ReportHandler, transactionHandler *handler.TransactionHandler) {

	http.HandleFunc("/inventories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			h.GetAll(w, r)
		case "POST":
			h.Create(w, r)
		case "PUT":
			h.Update(w, r)
		case "DELETE":
			h.Delete(w, r)
		}
	})

	// Endpoint report sales
	http.HandleFunc("/reports/sales", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		reportHandler.GetSalesReport(w, r)
	})

}
