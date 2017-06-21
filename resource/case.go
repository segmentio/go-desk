package resource

import (
	. "github.com/segmentio/go-desk/types"
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
	Links           struct {
		Labels struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"labels"`
		AssignedGroup struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"assigned_group"`
		MacroPreview struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"macro_preview"`
		LockedBy  interface{} `json:"locked_by"`
		CaseLinks struct {
			Class string `json:"class"`
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"case_links"`
		Attachments struct {
			Class string `json:"class"`
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"attachments"`
		History struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"history"`
		Feedbacks interface{} `json:"feedbacks"`
		Customer  struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"customer"`
		Replies struct {
			Class string `json:"class"`
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"replies"`
		Draft struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"draft"`
		Notes struct {
			Class string `json:"class"`
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"notes"`
		AssignedUser struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"assigned_user"`
		Self struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"self"`
		Message struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"message"`
	} `json:"_links"`
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
