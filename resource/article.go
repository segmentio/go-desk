package resource

import (
	. "github.com/talbright/go-desk/types"
)

type Article struct {
	ID              *int64     `json:"id,omitempty"`
	Subject         *string    `json:"subject,omitempty"`
	Position        *int       `json:"position,omitempty"`
	Quickcode       *string    `json:"quickcode,omitempty"`
	PublishAt       *Timestamp `json:"publish_at,omitempty"`
	CreatedAt       *Timestamp `json:"created_at,omitempty"`
	UpdatedAt       *Timestamp `json:"updated_at,omitempty"`
	Rating          *int       `json:"rating,omitempty"`
	RatingCount     *int       `json:"rating_count,omitempty"`
	RatingScore     *int       `json:"rating_score,omitempty"`
	PublicUrl       *string    `json:"public_url,omitempty"`
	InSupportCenter *bool      `json:"in_support_center,omitempty"`
	InternalNotes   *string    `json:"internal_notes,omitempty"`
	Locale          *string    `json:"locale,omitempty"`
	Body            *string    `json:"body,omitempty"`
	BodyEmail       *string    `json:"body_email,omitempty"`
	BodyEmailAuto   *bool      `json:"body_email_auto,omitempty"`
	BodyChat        *string    `json:"body_chat,omitempty"`
	BodyChatAuto    *bool      `json:"body_chat_auto,omitempty"`
	BodyWebCallback *string    `json:"body_web_callback,omitempty"`
	BodyTwitter     *string    `json:"body_twitter,omitempty"`
	BodyQna         *string    `json:"body_qna,omitempty"`
	BodyPhone       *string    `json:"body_phone,omitempty"`
	BodyFacebook    *string    `json:"body_facebook,omitempty"`
	Resource
}

func NewArticle() *Article {
	article := &Article{}
	article.InitializeResource(article)
	return article
}

func (c Article) String() string {
	return Stringify(c)
}
