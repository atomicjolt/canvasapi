package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// UpdateDashboardPositions Updates the dashboard positions for a user for a given context.  This allows
// positions for the dashboard cards and elsewhere to be customized on a per
// user basis.
//
// The asset string parameter should be in the format 'context_id', for example
// 'course_42'
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type UpdateDashboardPositions struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *UpdateDashboardPositions) GetMethod() string {
	return "PUT"
}

func (t *UpdateDashboardPositions) GetURLPath() string {
	path := "users/{id}/dashboard_positions"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateDashboardPositions) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateDashboardPositions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UpdateDashboardPositions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UpdateDashboardPositions) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateDashboardPositions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
