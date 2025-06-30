package handler

import (
	"encoding/json"
	"net/http"
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