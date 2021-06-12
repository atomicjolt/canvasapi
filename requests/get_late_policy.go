package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetLatePolicy Returns the late policy for a course.
// https://canvas.instructure.com/doc/api/late_policy.html
//
// Path Parameters:
// # ID (Required) ID
//
type GetLatePolicy struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *GetLatePolicy) GetMethod() string {
	return "GET"
}

func (t *GetLatePolicy) GetURLPath() string {
	path := "courses/{id}/late_policy"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetLatePolicy) GetQuery() (string, error) {
	return "", nil
}

func (t *GetLatePolicy) GetBody() (string, error) {
	return "", nil
}

func (t *GetLatePolicy) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetLatePolicy) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
