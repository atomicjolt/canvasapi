package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ListEnabledFeaturesAccounts A paginated list of all features that are enabled on a given Account, Course, or User.
// Only the feature names are returned.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type ListEnabledFeaturesAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListEnabledFeaturesAccounts) GetMethod() string {
	return "GET"
}

func (t *ListEnabledFeaturesAccounts) GetURLPath() string {
	path := "accounts/{account_id}/features/enabled"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListEnabledFeaturesAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListEnabledFeaturesAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEnabledFeaturesAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEnabledFeaturesAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEnabledFeaturesAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
