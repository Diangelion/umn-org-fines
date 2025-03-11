package models

type CreateOrganization struct {
	OrganizationPhoto        string `form:"organizationPhoto"`
	OrganizationTitle        string `form:"organizationTitle"`
	OrganizationStartDate    string `form:"organizationStartDate"`
	OrganizationEndDate      string `form:"organizationEndDate"`
	OrganizationDescriptions string `form:"organizationDescriptions"`
}