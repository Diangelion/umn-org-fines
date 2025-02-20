package models

type DuplicateEmailError struct {
	Email string
}

func (email *DuplicateEmailError) Error() string {
	return "Email already exists."
}
