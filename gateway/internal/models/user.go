package models

type UserRegistration struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UserLogin struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
