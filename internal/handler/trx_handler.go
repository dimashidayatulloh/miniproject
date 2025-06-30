package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/gorilla/mux"
)

type TrxHandler struct {
	service *app.TrxService
}

func NewTrxHandler(service *app.TrxService) *TrxHandler {
	return &TrxHandler{service}
}

func (h *TrxHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserIDFromJWT(r) // gunakan dari package lain jika sudah ada
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var input app.TrxInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateTrx(userID, &input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TrxHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserIDFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	trxs, err := h.service.GetAllTrx(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trxs)
}

func (h *TrxHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserIDFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	trx, detail, err := h.service.GetTrxByID(userID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"trx":    trx,
		"detail": detail,
	})
}

func (h *TrxHandler) GetAllPaginatedFiltered(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserIDFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	page := 1
	limit := 10
	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	kodeInvoice := r.URL.Query().Get("kode_invoice")
	metode := r.URL.Query().Get("method_bayar")
	tanggal := r.URL.Query().Get("tanggal") // format YYYY-MM-DD
	minTotal, _ := strconv.Atoi(r.URL.Query().Get("min_total"))
	maxTotal, _ := strconv.Atoi(r.URL.Query().Get("max_total"))

	trxs, total, err := h.service.GetAllTrxPaginatedFiltered(userID, page, limit, kodeInvoice, metode, tanggal, minTotal, maxTotal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":        trxs,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}