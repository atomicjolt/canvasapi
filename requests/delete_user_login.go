package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteUserLogin Delete an existing login.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # UserID (Required) ID
// # ID (Required) ID
//
type DeleteUserLogin struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`
}

func (t *DeleteUserLogin) GetMethod() string {
	return "DELETE"
}

func (t *DeleteUserLogin) GetURLPath() string {
	path := "users/{user_id}/logins/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteUserLogin) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteUserLogin) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteUserLogin) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteUserLogin) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
