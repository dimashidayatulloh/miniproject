package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/pkg/jwt"
	"github.com/gorilla/mux"
)

type AlamatHandler struct {
	service *app.AlamatService
}

func NewAlamatHandler(service *app.AlamatService) *AlamatHandler {
	return &AlamatHandler{service}
}

// POST /alamat
func (h *AlamatHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserIDFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var input domain.Alamat
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.service.CreateAlamat(userID, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GET /alamat
func (h *AlamatHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserIDFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	alamat, err := h.service.GetAllAlamatByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alamat)
}

// PUT /alamat/{id}
func (h *AlamatHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	var input domain.Alamat
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateAlamat(userID, id, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /alamat/{id}
func (h *AlamatHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	if err := h.service.DeleteAlamat(userID, id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Helper untuk ambil userID dari JWT
func getUserIDFromJWT(r *http.Request) (int, bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, false
	}
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return 0, false
	}
	tokenString := bearerToken[1]
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		return 0, false
	}
	return claims.UserID, true
}