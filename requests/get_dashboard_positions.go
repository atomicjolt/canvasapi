package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetDashboardPositions Returns all dashboard positions that have been saved for a user.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
//
type GetDashboardPositions struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetDashboardPositions) GetMethod() string {
	return "GET"
}

func (t *GetDashboardPositions) GetURLPath() string {
	path := "users/{id}/dashboard_positions"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetDashboardPositions) GetQuery() (string, error) {
	return "", nil
}

func (t *GetDashboardPositions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetDashboardPositions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetDashboardPositions) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetDashboardPositions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
