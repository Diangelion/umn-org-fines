package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/utils"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
    // Hash the password
    hashedPassword, errHash := utils.HashPassword(user.Password)
    if errHash != nil {
        return errHash
    }

    // Assign the hashed password back to the user
    user.Password = hashedPassword

    // Save the user
    return s.repo.CreateUser(user)
}
