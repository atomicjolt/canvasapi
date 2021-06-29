package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// EnableDisableOrClearExplicitCspSettingCourses Either explicitly sets CSP to be on or off for courses and sub-accounts,
// or clear the explicit settings to default to those set by a parent account
//
// Note: If "inherited" and "settings_locked" are both true for this account or course,
// then the CSP setting cannot be modified.
// https://canvas.instructure.com/doc/api/content_security_policy_settings.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.Status (Required) . Must be one of enabled, disabled, inheritedIf set to "enabled" for an account, CSP will be enabled for all its courses and sub-accounts (that
//    have not explicitly enabled or disabled it), using the allowed domains set on this account.
//    If set to "disabled", CSP will be disabled for this account or course and for all sub-accounts
//    that have not explicitly re-enabled it.
//    If set to "inherited", this account or course will reset to the default state where CSP settings
//    are inherited from the first parent account to have them explicitly set.
//
type EnableDisableOrClearExplicitCspSettingCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Status string `json:"status" url:"status,omitempty"` //  (Required) . Must be one of enabled, disabled, inherited
	} `json:"form"`
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) GetMethod() string {
	return "PUT"
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) GetURLPath() string {
	path := "courses/{course_id}/csp_settings"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.Status == "" {
		errs = append(errs, "'Form.Status' is required")
	}
	if t.Form.Status != "" && !string_utils.Include([]string{"enabled", "disabled", "inherited"}, t.Form.Status) {
		errs = append(errs, "Status must be one of enabled, disabled, inherited")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EnableDisableOrClearExplicitCspSettingCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
