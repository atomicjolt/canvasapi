package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GroupActivityStreamSummary Returns a summary of the current user's group-specific activity stream.
//
// For full documentation, see the API documentation for the user activity
// stream summary, in the user api.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
//
type GroupActivityStreamSummary struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`
}

func (t *GroupActivityStreamSummary) GetMethod() string {
	return "GET"
}

func (t *GroupActivityStreamSummary) GetURLPath() string {
	path := "groups/{group_id}/activity_stream/summary"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GroupActivityStreamSummary) GetQuery() (string, error) {
	return "", nil
}

func (t *GroupActivityStreamSummary) GetBody() (string, error) {
	return "", nil
}

func (t *GroupActivityStreamSummary) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GroupActivityStreamSummary) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
