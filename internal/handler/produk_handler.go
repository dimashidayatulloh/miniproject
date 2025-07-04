package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/gorilla/mux"
)

type ProdukHandler struct {
	service *app.ProdukService
}

func NewProdukHandler(service *app.ProdukService) *ProdukHandler {
	return &ProdukHandler{service}
}

func (h *ProdukHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.Produk
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateProduk(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProdukHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	var input domain.Produk
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateProduk(id, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProdukHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteProduk(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProdukHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	produk, err := h.service.GetProdukByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func (h *ProdukHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	produks, err := h.service.GetAllProduk()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produks)
}

func (h *ProdukHandler) GetByToko(w http.ResponseWriter, r *http.Request) {
	idToko, err := strconv.Atoi(mux.Vars(r)["id_toko"])
	if err != nil {
		http.Error(w, "Invalid id_toko", http.StatusBadRequest)
		return
	}
	produks, err := h.service.GetProdukByToko(idToko)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produks)
}

func (h *ProdukHandler) GetAllPaginated(w http.ResponseWriter, r *http.Request) {
    // Pagination
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if limit < 1 {
        limit = 10
    }

    // Filter
    nama := r.URL.Query().Get("nama")
    kategori := r.URL.Query().Get("kategori")
    hargaMinStr := r.URL.Query().Get("harga_min")
    hargaMaxStr := r.URL.Query().Get("harga_max")
    var hargaMin, hargaMax int
    if hargaMinStr != "" {
        hargaMin, _ = strconv.Atoi(hargaMinStr)
    }
    if hargaMaxStr != "" {
        hargaMax, _ = strconv.Atoi(hargaMaxStr)
    }

    produks, total, err := h.service.GetAllProdukFiltered(page, limit, nama, kategori, hargaMin, hargaMax)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    totalPages := int((total + int64(limit) - 1) / int64(limit))
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "data":        produks,
        "page":        page,
        "limit":       limit,
        "total":       total,
        "total_pages": totalPages,
    })
}