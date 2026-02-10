package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kasir.api/domain"
	service "kasir.api/services"
)

type InventoryHandler struct {
	service *service.InventoryService
}

func NewInventoryHandler(s *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{s}
}

func (h *InventoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var inv domain.Inventories
	json.NewDecoder(r.Body).Decode(&inv)

	if err := h.service.Create(&inv); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(inv)
}

func (h *InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println(name)
	data, err := h.service.GetAll(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *InventoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var inv domain.Inventories
	json.NewDecoder(r.Body).Decode(&inv)

	if err := h.service.Update(id, &inv); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(inv)
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message":"deleted"}`))
}
