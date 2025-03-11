package models

type RegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserId struct {
	UserId string `json:"user_id"`
}

type EditUser struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	ProfilePhoto string `json:"profilephoto"`
	CoverPhoto   string `json:"coverphoto"`
}