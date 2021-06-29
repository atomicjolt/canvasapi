package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListUserLoginsUsers Given a user ID, return a paginated list of that user's logins for the given account.
// https://canvas.instructure.com/doc/api/logins.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ListUserLoginsUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
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

func (t *ListUserLoginsUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListUserLoginsUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListUserLoginsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
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
