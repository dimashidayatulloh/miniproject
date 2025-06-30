package middleware

import (
	"net/http"
	"strings"

	"github.com/dimashidayatulloh/miniproject/pkg/jwt"
)

func JWTAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }
        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
            return
        }

        token := bearerToken[1]
        claims, err := jwt.ValidateToken(token)
        if err != nil {
            http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
            return
        }
        r.Header.Set("user_id", string(rune(claims.UserID)))
        next.ServeHTTP(w, r)
    })
}