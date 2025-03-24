package models

type CreateOrganization struct {
	OrganizationPhoto        string `form:"organizationPhoto" json:"organization_photo"`
	OrganizationTitle        string `form:"organizationTitle" json:"organization_title"`
	OrganizationStartDate    string `form:"organizationStartDate" json:"organization_start_date"`
	OrganizationEndDate      string `form:"organizationEndDate" json:"organization_end_date"`
	OrganizationDescriptions string `form:"organizationDescriptions" json:"organization_descriptions"`
}

type ListOrganization struct {
	CreateOrganization
	ProfilePageLeftValue float32
	ProfilePageZValue    int
}