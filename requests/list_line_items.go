package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListLineItems
// https://canvas.instructure.com/doc/api/line_items.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Tag (Optional) If specified only Line Items with this tag will be included.
// # ResourceID (Optional) If specified only Line Items with this resource_id will be included.
// # ResourceLinkID (Optional) If specified only Line Items attached to the specified resource_link_id will be included.
// # Limit (Optional) May be used to limit the number of Line Items returned in a page
//
type ListLineItems struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Tag            string `json:"tag" url:"tag,omitempty"`                           //  (Optional)
		ResourceID     string `json:"resource_id" url:"resource_id,omitempty"`           //  (Optional)
		ResourceLinkID string `json:"resource_link_id" url:"resource_link_id,omitempty"` //  (Optional)
		Limit          string `json:"limit" url:"limit,omitempty"`                       //  (Optional)
	} `json:"query"`
}

func (t *ListLineItems) GetMethod() string {
	return "GET"
}

func (t *ListLineItems) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListLineItems) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListLineItems) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLineItems) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLineItems) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLineItems) Do(c *canvasapi.Canvas) (*models.LineItem, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.LineItem{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
