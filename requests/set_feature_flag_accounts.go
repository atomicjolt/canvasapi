package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// SetFeatureFlagAccounts Set a feature flag for a given Account, Course, or User. This call will fail if a parent account sets
// a feature flag for the same feature in any state other than "allowed".
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # AccountID (Required) ID
// # Feature (Required) ID
//
// Form Parameters:
// # State (Optional) . Must be one of off, allowed, on"off":: The feature is not available for the course, user, or account and sub-accounts.
//    "allowed":: (valid only on accounts) The feature is off in the account, but may be enabled in
//                sub-accounts and courses by setting a feature flag on the sub-account or course.
//    "on":: The feature is turned on unconditionally for the user, course, or account and sub-accounts.
//
type SetFeatureFlagAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		Feature   string `json:"feature"`    //  (Required)
	} `json:"path"`

	Form struct {
		State string `json:"state"` //  (Optional) . Must be one of off, allowed, on
	} `json:"form"`
}

func (t *SetFeatureFlagAccounts) GetMethod() string {
	return "PUT"
}

func (t *SetFeatureFlagAccounts) GetURLPath() string {
	path := "accounts/{account_id}/features/flags/{feature}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{feature}", fmt.Sprintf("%v", t.Path.Feature))
	return path
}

func (t *SetFeatureFlagAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *SetFeatureFlagAccounts) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *SetFeatureFlagAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.Feature == "" {
		errs = append(errs, "'Feature' is required")
	}
	if !string_utils.Include([]string{"off", "allowed", "on"}, t.Form.State) {
		errs = append(errs, "State must be one of off, allowed, on")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetFeatureFlagAccounts) Do(c *canvasapi.Canvas) (*models.FeatureFlag, error) {
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
