package models

type GetListOrganization struct {
	List []string `json:"list"`
}

type CreateOrganization struct {
	OrganizationPhoto        string `form:"organizationPhoto"`
	OrganizationTitle        string `form:"organizationTitle"`
	OrganizationStartDate    string `form:"organizationStartDate"`
	OrganizationEndDate      string `form:"organizationEndDate"`
	OrganizationDescriptions string `form:"organizationDescriptions"`
}