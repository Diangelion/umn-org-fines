package models

type GetListOrganization struct {
	List []CreateOrganization `json:"list"`
}

type CreateOrganization struct {
	OrganizationPhoto        string `form:"organizationPhoto", json:"organization_photo"`
	OrganizationName         string `form:"organizationName", json:"organization_name"`
	OrganizationStartDate    string `form:"organizationStartDate", json:"organization_start_date"`
	OrganizationEndDate      string `form:"organizationEndDate", json:"organization_end_date"`
	OrganizationDescriptions string `form:"organizationDescriptions", json:"organization_descriptions"`
}