package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// SetFeatureFlagUsers Set a feature flag for a given Account, Course, or User. This call will fail if a parent account sets
// a feature flag for the same feature in any state other than "allowed".
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # UserID (Required) ID
// # Feature (Required) ID
//
// Form Parameters:
// # State (Optional) . Must be one of off, allowed, on"off":: The feature is not available for the course, user, or account and sub-accounts.
//    "allowed":: (valid only on accounts) The feature is off in the account, but may be enabled in
//                sub-accounts and courses by setting a feature flag on the sub-account or course.
//    "on":: The feature is turned on unconditionally for the user, course, or account and sub-accounts.
//
type SetFeatureFlagUsers struct {
	Path struct {
		UserID  string `json:"user_id" url:"user_id,omitempty"` //  (Required)
		Feature string `json:"feature" url:"feature,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		State string `json:"state" url:"state,omitempty"` //  (Optional) . Must be one of off, allowed, on
	} `json:"form"`
}

func (t *SetFeatureFlagUsers) GetMethod() string {
	return "PUT"
}

func (t *SetFeatureFlagUsers) GetURLPath() string {
	path := "users/{user_id}/features/flags/{feature}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{feature}", fmt.Sprintf("%v", t.Path.Feature))
	return path
}

func (t *SetFeatureFlagUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *SetFeatureFlagUsers) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SetFeatureFlagUsers) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SetFeatureFlagUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.Feature == "" {
		errs = append(errs, "'Feature' is required")
	}
	if t.Form.State != "" && !string_utils.Include([]string{"off", "allowed", "on"}, t.Form.State) {
		errs = append(errs, "State must be one of off, allowed, on")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetFeatureFlagUsers) Do(c *canvasapi.Canvas) (*models.FeatureFlag, error) {
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
