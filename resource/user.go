package resource

import (
	. "github.com/segmentio/go-desk/types"
)

type User struct {
	ID             *int64     `json:"id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	PublicName     *string    `json:"public_name,omitempty"`
	Email          *string    `json:"email,omitempty"`
	EmailVerified  *bool      `json:"email_verified,omitempty"`
	Available      *bool      `json:"available,omitempty"`
	Avatar         *string    `json:"avatar,omitempty"`
	Level          *string    `json:"level,omitempty"`
	CreatedAt      *Timestamp `json:"created_at,omitempty"`
	UpdatedAt      *Timestamp `json:"updated_at,omitempty"`
	CurrentLoginAt *Timestamp `json:"current_login_at,omitempty"`
	LastLoginAt    *Timestamp `json:"last_login_at,omitempty"`
	Links          struct {
		Self struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"self"`
		Preferences struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"preferences"`
		Searches struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"searches"`
		Groups struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"groups"`
		Macros struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"macros"`
		Filters struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"filters"`
		IntegrationUrls struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"integration_urls"`
		CaseLayout struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"case_layout"`
		PhoneIntegration interface{} `json:"phone_integration"`
	} `json:"_links,omitempty"`
	Resource
}

func NewUser() *User {
	user := &User{}
	user.InitializeResource(user)
	return user
}

func (c User) String() string {
	return Stringify(c)
}
