package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateAccount Update an existing account.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Account (Optional) Updates the account name
// # Account (Optional) Updates the account sis_account_id
//    Must have manage_sis permission and must not be a root_account.
// # Account (Optional) The default time zone of the account. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # Account (Optional) The default course storage quota to be used, if not otherwise specified.
// # Account (Optional) The default user storage quota to be used, if not otherwise specified.
// # Account (Optional) The default group storage quota to be used, if not otherwise specified.
// # Account (Optional) The ID of a course to be used as a template for all newly created courses.
//    Empty means to inherit the setting from parent account, 0 means to not
//    use a template even if a parent account has one set. The course must be
//    marked as a template.
// # Account (Optional) Restrict students from viewing courses after end date
// # Account (Optional) Lock this setting for sub-accounts and courses
// # Account (Optional) Restrict students from viewing courses before start date
// # Account (Optional) Determines whether this account has Microsoft Teams Sync enabled or not.
//
//    Note that if you are altering Microsoft Teams sync settings you must enable
//    the Microsoft Group enrollment syncing feature flag. In addition, if you are enabling
//    Microsoft Teams sync, you must also specify a tenant and login attribute.
// # Account (Optional) The tenant this account should use when using Microsoft Teams Sync.
//    This should be an Azure Active Directory domain name.
// # Account (Optional) The attribute this account should use to lookup users when using Microsoft Teams Sync.
//    Must be one of sub, email, oid, or preferred_username.
// # Account (Optional) Lock this setting for sub-accounts and courses
// # Account (Optional) Disable comments on announcements
// # Account (Optional) Lock this setting for sub-accounts and courses
// # Account (Optional) Copyright and license information must be provided for files before they are published.
// # Account (Optional) Lock this setting for sub-accounts and courses
// # Account (Optional) Restrict students from viewing future enrollments in course list
// # Account (Optional) Lock this setting for sub-accounts and courses
// # Account (Optional) [DEPRECATED] Restrict instructors from changing mastery scale
// # Account (Optional) [DEPRECATED] Lock this setting for sub-accounts and courses
// # Account (Optional) [DEPRECATED] Restrict instructors from changing proficiency calculation method
// # Account (Optional) [DEPRECATED] Lock this setting for sub-accounts and courses
// # Account (Optional) Give this a set of keys and boolean values to enable or disable services matching the keys
//
type UpdateAccount struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Account struct {
			Name                       string `json:"name"`                           //  (Optional)
			SISAccountID               string `json:"sis_account_id"`                 //  (Optional)
			DefaultTimeZone            string `json:"default_time_zone"`              //  (Optional)
			DefaultStorageQuotaMb      int64  `json:"default_storage_quota_mb"`       //  (Optional)
			DefaultUserStorageQuotaMb  int64  `json:"default_user_storage_quota_mb"`  //  (Optional)
			DefaultGroupStorageQuotaMb int64  `json:"default_group_storage_quota_mb"` //  (Optional)
			CourseTemplateID           int64  `json:"course_template_id"`             //  (Optional)
			Settings                   struct {
				RestrictStudentPastView struct {
					Value  bool `json:"value"`  //  (Optional)
					Locked bool `json:"locked"` //  (Optional)
				} `json:"restrict_student_past_view"`

				RestrictStudentFutureView struct {
					Value  bool `json:"value"`  //  (Optional)
					Locked bool `json:"locked"` //  (Optional)
				} `json:"restrict_student_future_view"`

				MicrosoftSyncEnabled        bool   `json:"microsoft_sync_enabled"`         //  (Optional)
				MicrosoftSyncTenant         string `json:"microsoft_sync_tenant"`          //  (Optional)
				MicrosoftSyncLoginAttribute string `json:"microsoft_sync_login_attribute"` //  (Optional)
				LockAllAnnouncements        struct {
					Value  bool `json:"value"`  //  (Optional)
					Locked bool `json:"locked"` //  (Optional)
				} `json:"lock_all_announcements"`

				UsageRightsRequired struct {
					Value  bool `json:"value"`  //  (Optional)
					Locked bool `json:"locked"` //  (Optional)
				} `json:"usage_rights_required"`

				RestrictStudentFutureListing struct {
					Value  bool `json:"value"`  //  (Optional)
					Locked bool `json:"locked"` //  (Optional)
				} `json:"restrict_student_future_listing"`

				LockOutcomeProficiency struct {
					Value bool `json:"value"` //  (Optional)
				} `json:"lock_outcome_proficiency"`

				LockProficiencyCalculation struct {
					Value bool `json:"value"` //  (Optional)
				} `json:"lock_proficiency_calculation"`
			} `json:"settings"`

			LockOutcomeProficiency struct {
				Locked bool `json:"locked"` //  (Optional)
			} `json:"lock_outcome_proficiency"`

			LockProficiencyCalculation struct {
				Locked bool `json:"locked"` //  (Optional)
			} `json:"lock_proficiency_calculation"`

			Services string `json:"services"` //  (Optional)
		} `json:"account"`
	} `json:"form"`
}

func (t *UpdateAccount) GetMethod() string {
	return "PUT"
}

func (t *UpdateAccount) GetURLPath() string {
	path := "accounts/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateAccount) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateAccount) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateAccount) Do(c *canvasapi.Canvas) (*models.Account, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Account{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
