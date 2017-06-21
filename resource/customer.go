package resource

import (
	. "github.com/segmentio/go-desk/types"
)

type Customer struct {
	ID                     *int64                 `json:"id,omitempty"`
	ExternalID             *string                `json:"external_id,omitempty"`
	FirstName              *string                `json:"first_name,omitempty"`
	LastName               *string                `json:"last_name,omitempty"`
	Company                *string                `json:"company,omitempty"`
	Title                  *string                `json:"title,omitempty"`
	Avatar                 *string                `json:"avatar,omitempty"`
	Background             *string                `json:"background,omitempty"`
	Language               *string                `json:"language,omitempty"`
	LockedUntil            *Timestamp             `json:"locked_until,omitempty"`
	CreatedAt              *Timestamp             `json:"created_at,omitempty"`
	UpdatedAt              *Timestamp             `json:"updated_at,omitempty"`
	CustomFields           map[string]interface{} `json:"custom_fields,omitempty"`
	CanAccessPrivatePortal *bool                  `json:"access_private_portal,omitempty"`
	CanAccessCompanyCases  *bool                  `json:"access_company_cases,omitempty"`
	Emails                 []map[string]string    `json:"emails,omitempty"`
	PhoneNumbers           []map[string]string    `json:"phone_numbers,omitempty"`
	Addresses              []map[string]string    `json:"addresses,omitempty"`
	Links                  struct {
		Self struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"self"`
		Cases struct {
			Class string `json:"class"`
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"cases"`
		Company struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"company"`
		FacebookUser interface{} `json:"facebook_user"`
		TwitterUser  interface{} `json:"twitter_user"`
		LockedBy     interface{} `json:"locked_by"`
		Labels       struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"labels"`
	} `json:"_links,omitempty"`
	Resource
}

func NewCustomer() *Customer {
	customer := &Customer{}
	customer.InitializeResource(customer)
	return customer
}

func (c Customer) String() string {
	return Stringify(c)
}

func (c *Customer) AddEmail(email string, emailType string) {
	c.Emails = c.AddToSlice(c.Emails, email, emailType)
}

func (c *Customer) AddAddress(address string, addressType string) {
	c.Addresses = c.AddToSlice(c.Addresses, address, addressType)
}

func (c *Customer) AddPhoneNumber(phone string, phoneType string) {
	c.PhoneNumbers = c.AddToSlice(c.PhoneNumbers, phone, phoneType)
}

func (c *Customer) AddToSlice(slice []map[string]string, value string, valueType string) []map[string]string {
	pair := make(map[string]string)
	pair["value"] = value
	pair["type"] = valueType
	if slice == nil {
		slice = make([]map[string]string, 1, 2)
		slice[0] = pair
	} else {
		slice = append(slice, pair)
	}
	return slice
}
