package resource

import (
	. "github.com/segmentio/go-desk/types"
)

type Company struct {
	ID           *int64                 `json:"id,omitempty"`
	ExternalID   *string                `json:"external_id,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Domains      []string               `json:"domains,omitempty"`
	CreatedAt    *Timestamp             `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp             `json:"updated_at,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Links        struct {
		Self struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"self"`
		Customers struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"customers"`
		Cases struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"cases"`
		Labels struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"labels"`
	} `json:"_links,omitempty"`
	Resource
}

func NewCompany() *Company {
	company := &Company{}
	company.InitializeResource(company)
	return company
}

func (c Company) String() string {
	return Stringify(c)
}

func (c *Company) AddDomain(domain string) {
	c.Domains = append(c.Domains, domain)
}
