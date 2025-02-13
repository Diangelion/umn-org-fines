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
    if hashedPassword, errHash := utils.HashPassword(user.Password); errHash != nil {
        return errHash
    } else {
        user.Password = hashedPassword // Store in user model
    }

    // Create the user
    return s.repo.CreateUser(user)
}

func (s *UserService) LoginUser(user *models.UserLogin) error {
    if errCheck := s.repo.CheckCredential(user); errCheck != nil {
        return errCheck
    }
    return nil
}
