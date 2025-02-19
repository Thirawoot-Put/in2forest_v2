package dto

type AuthResponse struct {
	AccessToken string
}

type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
