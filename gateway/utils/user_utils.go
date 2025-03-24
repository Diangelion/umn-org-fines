package utils

import (
	"errors"
	"gateway/internal/models"
)

// Parse form data and return an error if parsing fails
func CombineProfileAndOrganizations(user *models.User, userRes models.Response, orgRes models.Response) error {
	user.Profile.Name = userRes.Data["name"].(string)
	user.Profile.Email = userRes.Data["email"].(string)
	user.Profile.ProfilePhoto = userRes.Data["profile_photo"].(string)
	user.Profile.CoverPhoto = userRes.Data["cover_photo"].(string)

	orgs, ok := orgRes.Data["list"].([]interface{})
	if !ok {
		return errors.New("Expected []interface{} for 'list'")
	}

	// Allocate correct slice type
	user.Organizations = make([]models.ListOrganization, len(orgs))

	// Convert each element
	for i, org := range orgs {
		orgMap, ok := org.(map[string]interface{})
		if !ok {
			return errors.New("Expected map[string]interface{} for org item")
		}

		user.Organizations[i] = models.ListOrganization{
			CreateOrganization: models.CreateOrganization{ // ✅ Embed CreateOrganization properly
				OrganizationPhoto:        orgMap["organization_photo"].(string),
				OrganizationTitle:        orgMap["organization_title"].(string),
				OrganizationStartDate:    orgMap["organization_start_date"].(string),
				OrganizationEndDate:      orgMap["organization_end_date"].(string),
				OrganizationDescriptions: orgMap["organization_descriptions"].(string),
			},
			ProfilePageLeftValue: float32(i) * 0.2,
			ProfilePageZValue: 0 - i,
		}

	}

	user.TotalOrganizations = len(user.Organizations)

	return nil
}