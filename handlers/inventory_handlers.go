package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir.api/libs"
	"kasir.api/models"
	"kasir.api/services"
)

type InventoryHandler struct {
	service *services.InventoryService
}

func NewInventoryHandler(service *services.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

// HandleInventory - GET /api/inventory
func (h *InventoryHandler) HandleInventory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		libs.HandleError(http.StatusMethodNotAllowed, w, "Method not allowed")
	}
}

func (h *InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	inventory, err := h.service.GetAll()
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, inventory)
}

func (h *InventoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateInventoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid request body")
		return
	}

	product, err := h.service.Create(req)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusCreated, w, product)
}

// HandleInventoryByID - GET/PUT/DELETE /api/inventory/{id}
func (h *InventoryHandler) HandleInventoryByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		libs.HandleError(http.StatusMethodNotAllowed, w, "Method not allowed")
	}
}

// GetByID - GET /api/inventory/{id}
func (h *InventoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/inventory/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid inventory ID")
		return
	}

	inventory, err := h.service.GetByID(id)
	if err != nil {
		libs.HandleError(http.StatusNotFound, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, inventory)
}

func (h *InventoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/inventory/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid inventory ID")
		return
	}

	var req models.UpdateInventoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid request body")
		return
	}

	inventory, err := h.service.Update(id, req)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, inventory)
}

// Delete - DELETE /api/inventory/{id}
func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/inventory/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid inventory ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, nil, "Inventory deleted successfully")
}
