package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCustomColors Returns all custom colors that have been saved for a user.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
//
type GetCustomColors struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *GetCustomColors) GetMethod() string {
	return "GET"
}

func (t *GetCustomColors) GetURLPath() string {
	path := "users/{id}/colors"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetCustomColors) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCustomColors) GetBody() (string, error) {
	return "", nil
}

func (t *GetCustomColors) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCustomColors) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
