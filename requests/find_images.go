package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// FindImages Find public domain images for use in courses and user content.  If you select an image using this API, please use the {api:InternetImageController#image_selection Confirm image selection API} to indicate photo usage to the server.
// https://canvas.instructure.com/doc/api/image_search.html
//
// Query Parameters:
// # Query (Required) Search terms used for matching images (e.g. "cats").
//
type FindImages struct {
	Query struct {
		Query string `json:"query" url:"query,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *FindImages) GetMethod() string {
	return "GET"
}

func (t *FindImages) GetURLPath() string {
	return ""
}

func (t *FindImages) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *FindImages) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *FindImages) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *FindImages) HasErrors() error {
	errs := []string{}
	if t.Query.Query == "" {
		errs = append(errs, "'Query' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *FindImages) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
