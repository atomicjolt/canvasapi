package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListUserLoginsUsers Given a user ID, return a paginated list of that user's logins for the given account.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # UserID (Required) ID
//
type ListUserLoginsUsers struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`
}

func (t *ListUserLoginsUsers) GetMethod() string {
	return "GET"
}

func (t *ListUserLoginsUsers) GetURLPath() string {
	path := "users/{user_id}/logins"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListUserLoginsUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListUserLoginsUsers) GetBody() (string, error) {
	return "", nil
}

func (t *ListUserLoginsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUserLoginsUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
