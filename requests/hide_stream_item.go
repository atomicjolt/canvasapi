package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// HideStreamItem Hide the given stream item.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
//
type HideStreamItem struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *HideStreamItem) GetMethod() string {
	return "DELETE"
}

func (t *HideStreamItem) GetURLPath() string {
	path := "users/self/activity_stream/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *HideStreamItem) GetQuery() (string, error) {
	return "", nil
}

func (t *HideStreamItem) GetBody() (string, error) {
	return "", nil
}

func (t *HideStreamItem) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *HideStreamItem) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
