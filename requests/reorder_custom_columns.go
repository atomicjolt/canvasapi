package requests

import (
	"fmt"
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
// # CourseID (Required) ID
//
// Form Parameters:
// # Order (Required) no description
//
type ReorderCustomColumns struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Order []int64 `json:"order"` //  (Required)
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

func (t *ReorderCustomColumns) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ReorderCustomColumns) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Order == nil {
		errs = append(errs, "'Order' is required")
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