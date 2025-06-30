package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/gorilla/mux"
)

type FotoProdukHandler struct {
	service *app.FotoProdukService
}

func NewFotoProdukHandler(service *app.FotoProdukService) *FotoProdukHandler {
	return &FotoProdukHandler{service}
}

func (h *FotoProdukHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.FotoProduk
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateFotoProduk(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FotoProdukHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	foto, err := h.service.GetFotoProdukByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foto)
}

func (h *FotoProdukHandler) GetAllByProduk(w http.ResponseWriter, r *http.Request) {
	idProduk, err := strconv.Atoi(mux.Vars(r)["id_produk"])
	if err != nil {
		http.Error(w, "Invalid id_produk", http.StatusBadRequest)
		return
	}
	fotos, err := h.service.GetAllFotoProdukByProduk(idProduk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fotos)
}