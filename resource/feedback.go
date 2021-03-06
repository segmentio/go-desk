package resource

import (
	. "github.com/segmentio/go-desk/types"
)

type Feedback struct {
	ID                 *int64     `json:"id,omitempty"`
	Rating             *string    `json:"rating,omitempty"`
	RatingType         *string    `json:"rating_type,omitempty"`
	AdditionalFeedback *string    `json:"additional_feedback,omitempty"`
	CreatedAt          *Timestamp `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp `json:"updated_at,omitempty"`
	Case               *string    `json:"case,omitempty"`
	User               *string    `json:"user,omitempty"`
	Links              struct {
		Self struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"self"`
		Case struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"case"`
		Customer struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"customer"`
		Reply struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"reply"`
		User struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"user"`
	} `json:"_links,omitempty"`
	Resource
}

func NewFeedback() *Feedback {
	feedback := &Feedback{}
	feedback.InitializeResource(feedback)
	return feedback
}

func (c Feedback) String() string {
	return Stringify(c)
}
