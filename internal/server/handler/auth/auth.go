package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/it-chep/my_optium_bot.git/internal/pkg/jwt"
)

// LoginHandler авторизация пользователя
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !jwt.CheckCredentials(loginReq.Username, loginReq.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateJWT(loginReq.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CheckValidHandler авторизация пользователя
func CheckValidHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		http.Error(w, "Bearer token required", http.StatusUnauthorized)
		return
	}

	if !jwt.Valid(tokenString) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
}
