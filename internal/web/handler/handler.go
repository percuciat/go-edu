package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"goProject/internal/service"
	"goProject/pkg/logger"

	"go.uber.org/zap"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {

	return &Handler{
		//logger: logger,
		service: service,
	}
}

func (h *Handler) HandleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getItems(w, r)
	case http.MethodPost:
		h.createItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleItem(w http.ResponseWriter, r *http.Request) {
	// convert the path to int by id
	id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
	logger.New().Info("HandleItem", zap.Int("id", id))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getItem(w, r, id)
	case http.MethodPut:
		h.updateItem(w, r, id)
	case http.MethodDelete:
		h.deleteItem(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {
	items := h.service.GetItems()
	logger.New().Info("HandleItem", zap.Any("items", items))
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) getItem(w http.ResponseWriter, r *http.Request, id int) {
	item, err := h.service.GetItem(id)
	if err != nil {
		http.Error(w, "Error getting item", http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	var item service.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdItem := h.service.CreateItem(item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdItem)
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var item service.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedItem, err := h.service.UpdateItem(id, item)
	if err != nil {
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}
	if updatedItem == nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedItem)
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.DeleteItem(id); err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/web/templates/layout.html")
}
