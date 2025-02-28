package models

type UserRegistration struct {
	Name            string `form:"name"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

type ForwardUserRegistration struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UserLogin struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UserEdit struct {
	Name         string `form:"name"`
	Email        string `form:"email"`
	ProfilePhoto string `form:"profile-photo"`
	CoverPhoto   string `form:"cover-photo"`
}