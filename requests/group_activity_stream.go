package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GroupActivityStream Returns the current user's group-specific activity stream, paginated.
//
// For full documentation, see the API documentation for the user activity
// stream, in the user api.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
//
type GroupActivityStream struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GroupActivityStream) GetMethod() string {
	return "GET"
}

func (t *GroupActivityStream) GetURLPath() string {
	path := "groups/{group_id}/activity_stream"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GroupActivityStream) GetQuery() (string, error) {
	return "", nil
}

func (t *GroupActivityStream) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GroupActivityStream) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GroupActivityStream) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GroupActivityStream) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
