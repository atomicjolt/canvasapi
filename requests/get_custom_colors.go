package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCustomColors Returns all custom colors that have been saved for a user.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type GetCustomColors struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
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

func (t *GetCustomColors) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCustomColors) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCustomColors) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
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
