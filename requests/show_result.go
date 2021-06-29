package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowResult Show existing Result of a line item.
// https://canvas.instructure.com/doc/api/result.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.LineItemID (Required) ID
// # Path.ID (Required) ID
//
type ShowResult struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`       //  (Required)
		LineItemID string `json:"line_item_id" url:"line_item_id,omitempty"` //  (Required)
		ID         string `json:"id" url:"id,omitempty"`                     //  (Required)
	} `json:"path"`
}

func (t *ShowResult) GetMethod() string {
	return "GET"
}

func (t *ShowResult) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items/{line_item_id}/results/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{line_item_id}", fmt.Sprintf("%v", t.Path.LineItemID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowResult) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowResult) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowResult) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowResult) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.LineItemID == "" {
		errs = append(errs, "'Path.LineItemID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowResult) Do(c *canvasapi.Canvas) (*models.Result, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Result{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
