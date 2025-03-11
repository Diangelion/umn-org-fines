package models

type RegisterUser struct {
	Name            string `form:"name"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

type ForwardRegisterUser struct {
	Name     string
	Email    string
	Password string
}

type LoginUser struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type EditUser struct {
	Name         string `form:"name"`
	Email        string `form:"email"`
	ProfilePhoto string `form:"profile-photo"`
	CoverPhoto   string `form:"cover-photo"`
}