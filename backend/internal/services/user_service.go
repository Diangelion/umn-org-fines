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

func (s *UserService) RegisterUser(user *models.UserRegistration) error {
    // Hash the password
    hashedPassword, errHash := utils.HashPassword(user.Password)
    if errHash != nil {
        return errHash
    }
    user.Password = hashedPassword

    // Save the user
    return s.repo.CreateUser(user)
}

func (s *UserService) LoginUser(user *models.UserLogin) error {
    if errCheck := s.repo.CheckCredential(user); errCheck != nil {
        return errCheck
    }
    return nil
}
