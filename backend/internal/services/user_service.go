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

func (s *UserService) RegisterUserService(user *models.RegisterUser) error {
    // Hash the password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        log.Println("RegisterUserService | Hash password error: ", err)
        return errors.New("Failed to hash password.")
    }
    user.Password = hashedPassword // Store in user model

    // Create the user
    return s.repo.RegisterUserRepository(user)
}

func (s *UserService) LoginUserService(user *models.LoginUser) (string, error) {
    return s.repo.LoginUserRepository(user)
}

func (s *UserService) GetUserService(userId string) (*models.EditUser, error) {
    return s.repo.GetUserRepository(userId)
}

func (s *UserService) EditUserService(user *models.EditUser, userId string) error {
    return s.repo.EditUserRepository(user, userId)
}
