package service

import (
	"encoding/json"
	"net/http"
	"net/url"

	. "github.com/segmentio/go-desk/resource"
)

type ArticleService struct {
	client *Client
}

// List articles with filtering and pagination.
// See Desk API: http://dev.desk.com/API/articles/#list
func (c *ArticleService) List(params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewArticle())
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

func (c *ArticleService) unravelPage(page *Page) error {
	articles := new([]Article)
	err := json.Unmarshal(*page.Embedded.RawEntries, &articles)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*articles))
	for i, v := range *articles {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
