package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type OrganizationService struct {
    repo *repositories.OrganizationRepository
}

func NewOrganizationService(repo *repositories.OrganizationRepository) *OrganizationService {
    return &OrganizationService{repo}
}

func (s *OrganizationService) CreateOrganizationService(org *models.CreateOrganization, userId string) error {
    return s.repo.CreateOrganizationRepository(org, userId)
}