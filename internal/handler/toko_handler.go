package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/pkg/jwt"
)

type TokoHandler struct {
	service *app.TokoService
}

func NewTokoHandler(service *app.TokoService) *TokoHandler {
	return &TokoHandler{service}
}

// Get toko milik user yang sedang login (GET /toko/me)
func (h *TokoHandler) GetMyToko(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := bearerToken[1]
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	toko, err := h.service.GetTokoByUserID(claims.UserID)
	if err != nil {
		http.Error(w, "Toko tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toko)
}

// Update toko milik user yang sedang login (PUT /toko/me)
func (h *TokoHandler) UpdateMyToko(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := bearerToken[1]
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input domain.Toko
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateToko(claims.UserID, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TokoHandler) GetAllPaginatedFiltered(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	nama := r.URL.Query().Get("nama")
	urlFoto := r.URL.Query().Get("url_foto")
	userID := 0
	if s := r.URL.Query().Get("id_user"); s != "" {
		userID, _ = strconv.Atoi(s)
	}

	tokos, total, err := h.service.GetAllTokoPaginatedFiltered(page, limit, nama, urlFoto, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":        tokos,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}