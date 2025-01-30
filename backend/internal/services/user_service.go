package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
    // Add any business logic here, like password hashing
    return s.repo.CreateUser(user)
}
