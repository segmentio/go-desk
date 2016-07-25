package resource

import (
	. "github.com/talbright/go-desk/types"
)

type Case struct {
	ID              *int64                 `json:"id,omitempty"`
	ExternalID      *string                `json:"external_id,omitempty"`
	Blurb           *string                `json:"blurb,omitempty"`
	Subject         *string                `json:"subject,omitempty"`
	Priority        *int                   `json:"priority,omitempty"`
	Description     *string                `json:"description,omitempty"`
	Status          *string                `json:"status,omitempty"`
	Type            *string                `json:"type,omitempty"`
	Language        *string                `json:"language,omitempty"`
	CreatedAt       *Timestamp             `json:"created_at,omitempty"`
	UpdatedAt       *Timestamp             `json:"updated_at,omitempty"`
	ChangedAt       *Timestamp             `json:"changed_at,omitempty"`
	LastActiveAt    *Timestamp             `json:"active_at,omitempty"`
	ReceivedAt      *Timestamp             `json:"received_at,omitempty"`
	LockedUntil     *Timestamp             `json:"locked_until,omitempty"`
	FirstOpenedAt   *Timestamp             `json:"first_opened_at,omitempty"`
	LastOpenedAt    *Timestamp             `json:"opened_at,omitempty"`
	FirstResolvedAt *Timestamp             `json:"first_resolved_at,omitempty"`
	LastResolvedAt  *Timestamp             `json:"resolved_at,omitempty"`
	CustomFields    map[string]interface{} `json:"custom_fields,omitempty"`
	Resource
}

func NewCase() *Case {
	caze := &Case{}
	caze.InitializeResource(caze)
	return caze
}

func (c Case) String() string {
	return Stringify(c)
}
