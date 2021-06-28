package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListEnabledFeaturesUsers A paginated list of all features that are enabled on a given Account, Course, or User.
// Only the feature names are returned.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # UserID (Required) ID
//
type ListEnabledFeaturesUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListEnabledFeaturesUsers) GetMethod() string {
	return "GET"
}

func (t *ListEnabledFeaturesUsers) GetURLPath() string {
	path := "users/{user_id}/features/enabled"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListEnabledFeaturesUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListEnabledFeaturesUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEnabledFeaturesUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEnabledFeaturesUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEnabledFeaturesUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
