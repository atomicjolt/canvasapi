package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// RemoveFeatureFlagAccounts Remove feature flag for a given Account, Course, or User.  (Note that the flag must
// be defined on the Account, Course, or User directly.)  The object will then inherit
// the feature flags from a higher account, if any exist.  If this flag was 'on' or 'off',
// then lower-level account flags that were masked by this one will apply again.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.Feature (Required) ID
//
type RemoveFeatureFlagAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		Feature   string `json:"feature" url:"feature,omitempty"`       //  (Required)
	} `json:"path"`
}

func (t *RemoveFeatureFlagAccounts) GetMethod() string {
	return "DELETE"
}

func (t *RemoveFeatureFlagAccounts) GetURLPath() string {
	path := "accounts/{account_id}/features/flags/{feature}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{feature}", fmt.Sprintf("%v", t.Path.Feature))
	return path
}

func (t *RemoveFeatureFlagAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *RemoveFeatureFlagAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveFeatureFlagAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveFeatureFlagAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.Feature == "" {
		errs = append(errs, "'Path.Feature' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveFeatureFlagAccounts) Do(c *canvasapi.Canvas) (*models.FeatureFlag, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.FeatureFlag{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
