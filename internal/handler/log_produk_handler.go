package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/gorilla/mux"
)

type LogProdukHandler struct {
	service *app.LogProdukService
}

func NewLogProdukHandler(service *app.LogProdukService) *LogProdukHandler {
	return &LogProdukHandler{service}
}

func (h *LogProdukHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.LogProduk
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateLogProduk(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *LogProdukHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	log, err := h.service.GetLogProdukByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(log)
}

func (h *LogProdukHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logs, err := h.service.GetAllLogProduk()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func (h *LogProdukHandler) GetAllPaginatedFiltered(w http.ResponseWriter, r *http.Request) {
	// Ambil query param
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	jenis := r.URL.Query().Get("jenis")
	keterangan := r.URL.Query().Get("keterangan")
	idProdukStr := r.URL.Query().Get("id_produk")
	idProduk := 0
	if idProdukStr != "" {
		idProduk, _ = strconv.Atoi(idProdukStr)
	}

	logs, total, err := h.service.GetAllLogProdukPaginatedFiltered(page, limit, jenis, keterangan, idProduk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":        logs,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}