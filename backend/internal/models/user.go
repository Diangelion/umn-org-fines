package models

type UserRegistration struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserId struct {
	UserId string `json:"user_id"`
}

type UserEdit struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	ProfilePhoto string `json:"profilephoto"`
	CoverPhoto   string `json:"coverphoto"`
}