package auth

// LoginResponse — структура для ответа при логине
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginRequest — структура для получения запроса от пользователя
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
