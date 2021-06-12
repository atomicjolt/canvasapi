package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowLineItem Show existing Line Item
// https://canvas.instructure.com/doc/api/line_items.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type ShowLineItem struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *ShowLineItem) GetMethod() string {
	return "GET"
}

func (t *ShowLineItem) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowLineItem) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowLineItem) GetBody() (string, error) {
	return "", nil
}

func (t *ShowLineItem) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowLineItem) Do(c *canvasapi.Canvas) (*models.LineItem, error) {
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
