package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Tag            string `json:"tag"`              //  (Optional)
		ResourceID     string `json:"resource_id"`      //  (Optional)
		ResourceLinkID string `json:"resource_link_id"` //  (Optional)
		Limit          string `json:"limit"`            //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListLineItems) GetBody() (string, error) {
	return "", nil
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
