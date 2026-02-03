package routes

import (
	"net/http"

	"kasir.api/handler"
)

func RegisterRoutes(h *handler.InventoryHandler) {
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
}
