package auth

// User модель пользователя
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequest запрос на авторизацию
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse ответ с токеном
type AuthResponse struct {
	Token string `json:"token"`
}
