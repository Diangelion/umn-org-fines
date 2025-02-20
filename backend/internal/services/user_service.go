package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/utils"
	"errors"
	"log"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo}
}

func (s *UserService) RegisterUser(user *models.UserRegistration) error {
    // Hash the password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        log.Println("RegisterUser | Hash password error: ", err)
        return errors.New("Failed to hash password.")
    }
    user.Password = hashedPassword // Store in user model
    
    // Create the user
    return s.repo.CreateUser(user)
}

func (s *UserService) LoginUser(user *models.UserLogin) (string, error) {
    userId, err := s.repo.CheckCredential(user)
    if err != nil {
        return "", err
    }
    return userId, nil
}
