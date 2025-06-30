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

type CategoryHandler struct {
	service *app.CategoryService
}

func NewCategoryHandler(service *app.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

// Helper: ambil userID dan isAdmin dari JWT
func getIsAdminFromJWT(r *http.Request) (int, bool, bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, false, false
	}
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return 0, false, false
	}
	tokenString := bearerToken[1]
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		return 0, false, false
	}
	return claims.UserID, claims.IsAdmin, true
}

// POST /category (admin only)
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	_, isAdmin, ok := getIsAdminFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var input domain.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateCategory(isAdmin, &input); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// PUT /category/{id} (admin only)
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	_, isAdmin, ok := getIsAdminFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	var input domain.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateCategory(isAdmin, id, &input); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /category/{id} (admin only)
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_, isAdmin, ok := getIsAdminFromJWT(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteCategory(isAdmin, id); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GET /category (public)
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	cats, err := h.service.GetAllCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
}

// GET /category?page=1&limit=10&nama=kopi
func (h *CategoryHandler) GetAllPaginatedFiltered(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	nama := r.URL.Query().Get("nama")

	cats, total, err := h.service.GetAllCategoryPaginatedFiltered(page, limit, nama)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":        cats,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}