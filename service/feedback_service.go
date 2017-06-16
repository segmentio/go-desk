package service

import (
	"encoding/json"
	"net/http"
	"net/url"

	. "github.com/segmentio/go-desk/resource"
)

type FeedbackService struct {
	client *Client
}

// List feedbacks with filtering and pagination.
// See Desk API: http://dev.desk.com/API/feedbacks/#list
func (c *FeedbackService) List(params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewFeedback())
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (c *FeedbackService) unravelPage(page *Page) error {
	feedbacks := new([]Feedback)
	err := json.Unmarshal(*page.Embedded.RawEntries, &feedbacks)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*feedbacks))
	for i, v := range *feedbacks {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
