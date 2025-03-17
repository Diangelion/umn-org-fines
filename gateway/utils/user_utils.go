package utils

import (
	"gateway/internal/models"
)

// Parse form data and return an error if parsing fails
func CombineProfileAndOrganizations(user *models.User, userRes models.Response, orgRes models.Response) error {
	user.Profile.Name = userRes.Data["name"].(string)
	user.Profile.Email = userRes.Data["email"].(string)
	user.Profile.ProfilePhoto = userRes.Data["profile_photo"].(string)
	user.Profile.CoverPhoto = userRes.Data["cover_photo"].(string)
	user.Organizations = orgRes.Data["list"].([]models.CreateOrganization)
	user.TotalOrganizations = len(orgRes.Data["list"].([]models.CreateOrganization))
	return nil
}