package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ReorderCustomColumns Puts the given columns in the specified order
//
// <b>200 OK</b> is returned if successful
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.Order (Required) no description
//
type ReorderCustomColumns struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Order []string `json:"order" url:"order,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *ReorderCustomColumns) GetMethod() string {
	return "POST"
}

func (t *ReorderCustomColumns) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns/reorder"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ReorderCustomColumns) GetQuery() (string, error) {
	return "", nil
}

func (t *ReorderCustomColumns) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ReorderCustomColumns) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ReorderCustomColumns) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.Order == nil {
		errs = append(errs, "'Form.Order' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReorderCustomColumns) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
