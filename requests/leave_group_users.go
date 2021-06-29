package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// LeaveGroupUsers Leave a group if you are allowed to leave (some groups, such as sets of
// course groups created by teachers, cannot be left). You may also use 'self'
// in place of a membership_id.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.UserID (Required) ID
//
type LeaveGroupUsers struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		UserID  string `json:"user_id" url:"user_id,omitempty"`   //  (Required)
	} `json:"path"`
}

func (t *LeaveGroupUsers) GetMethod() string {
	return "DELETE"
}

func (t *LeaveGroupUsers) GetURLPath() string {
	path := "groups/{group_id}/users/{user_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *LeaveGroupUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *LeaveGroupUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *LeaveGroupUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *LeaveGroupUsers) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *LeaveGroupUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
