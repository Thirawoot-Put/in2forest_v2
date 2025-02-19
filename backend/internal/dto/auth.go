package dto

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
}

type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
