package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

func (h *FotoProdukHandler) Upload(w http.ResponseWriter, r *http.Request) {
	// Limit upload max 5MB
	r.ParseMultipartForm(5 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	idProdukStr := r.FormValue("id_produk")
	idProduk, err := strconv.Atoi(idProdukStr)
	if err != nil || idProduk <= 0 {
		http.Error(w, "invalid id_produk", http.StatusBadRequest)
		return
	}

	// Pastikan folder upload ada
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	// Nama unik
	filename := fmt.Sprintf("%d_%d_%s", idProduk, time.Now().UnixNano(), handler.Filename)
	filepath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}

	// Simpan ke database
	now := time.Now()
	foto := &domain.FotoProduk{
		IdProduk:  idProduk,
		URL:       "/uploads/" + filename, // path untuk diakses client
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if err := h.service.CreateFotoProduk(foto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foto)
}

func (h *FotoProdukHandler) GetAllByProdukPaginatedFiltered(w http.ResponseWriter, r *http.Request) {
	idProduk, err := strconv.Atoi(mux.Vars(r)["id_produk"])
	if err != nil {
		http.Error(w, "Invalid id_produk", http.StatusBadRequest)
		return
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	url := r.URL.Query().Get("url")

	fotos, total, err := h.service.GetAllFotoProdukByProdukPaginatedFiltered(idProduk, page, limit, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":        fotos,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}