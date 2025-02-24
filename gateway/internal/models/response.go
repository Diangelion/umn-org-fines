package models

// Global alert
type Alert struct {
	Title   string
	Message string
}

// Auth page (e.g. login, register)
type AuthPage struct {
	BaseURL string
}

// Logged in
type AuthorizationToken struct {
	AccessToken  string
	RefreshToken string
}