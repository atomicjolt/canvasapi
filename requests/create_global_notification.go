package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateGlobalNotification Create and return a new global notification for an account.
// https://canvas.instructure.com/doc/api/account_notifications.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # AccountNotification (Required) The subject of the notification.
// # AccountNotification (Required) The message body of the notification.
// # AccountNotification (Required) The start date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # AccountNotification (Required) The end date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # AccountNotification (Optional) . Must be one of warning, information, question, error, calendarThe icon to display with the notification.
//    Note: Defaults to warning.
// # AccountNotificationRoles (Optional) The role(s) to send global notification to.  Note:  ommitting this field will send to everyone
//    Example:
//      account_notification_roles: ["StudentEnrollment", "TeacherEnrollment"]
//
type CreateGlobalNotification struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		AccountNotification struct {
			Subject string    `json:"subject"`  //  (Required)
			Message string    `json:"message"`  //  (Required)
			StartAt time.Time `json:"start_at"` //  (Required)
			EndAt   time.Time `json:"end_at"`   //  (Required)
			Icon    string    `json:"icon"`     //  (Optional) . Must be one of warning, information, question, error, calendar
		} `json:"account_notification"`

		AccountNotificationRoles []string `json:"account_notification_roles"` //  (Optional)
	} `json:"form"`
}

func (t *CreateGlobalNotification) GetMethod() string {
	return "POST"
}

func (t *CreateGlobalNotification) GetURLPath() string {
	path := "accounts/{account_id}/account_notifications"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateGlobalNotification) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateGlobalNotification) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateGlobalNotification) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.AccountNotification.Subject == "" {
		errs = append(errs, "'AccountNotification' is required")
	}
	if t.Form.AccountNotification.Message == "" {
		errs = append(errs, "'AccountNotification' is required")
	}
	if t.Form.AccountNotification.StartAt.IsZero() {
		errs = append(errs, "'AccountNotification' is required")
	}
	if t.Form.AccountNotification.EndAt.IsZero() {
		errs = append(errs, "'AccountNotification' is required")
	}
	if !string_utils.Include([]string{"warning", "information", "question", "error", "calendar"}, t.Form.AccountNotification.Icon) {
		errs = append(errs, "AccountNotification must be one of warning, information, question, error, calendar")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateGlobalNotification) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
