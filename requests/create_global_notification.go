package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
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
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.AccountNotification.Subject (Required) The subject of the notification.
// # Form.AccountNotification.Message (Required) The message body of the notification.
// # Form.AccountNotification.StartAt (Required) The start date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # Form.AccountNotification.EndAt (Required) The end date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # Form.AccountNotification.Icon (Optional) . Must be one of warning, information, question, error, calendarThe icon to display with the notification.
//    Note: Defaults to warning.
// # Form.AccountNotificationRoles (Optional) The role(s) to send global notification to.  Note:  ommitting this field will send to everyone
//    Example:
//      account_notification_roles: ["StudentEnrollment", "TeacherEnrollment"]
//
type CreateGlobalNotification struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		AccountNotification struct {
			Subject string    `json:"subject" url:"subject,omitempty"`   //  (Required)
			Message string    `json:"message" url:"message,omitempty"`   //  (Required)
			StartAt time.Time `json:"start_at" url:"start_at,omitempty"` //  (Required)
			EndAt   time.Time `json:"end_at" url:"end_at,omitempty"`     //  (Required)
			Icon    string    `json:"icon" url:"icon,omitempty"`         //  (Optional) . Must be one of warning, information, question, error, calendar
		} `json:"account_notification" url:"account_notification,omitempty"`

		AccountNotificationRoles []string `json:"account_notification_roles" url:"account_notification_roles,omitempty"` //  (Optional)
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

func (t *CreateGlobalNotification) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateGlobalNotification) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateGlobalNotification) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Form.AccountNotification.Subject == "" {
		errs = append(errs, "'Form.AccountNotification.Subject' is required")
	}
	if t.Form.AccountNotification.Message == "" {
		errs = append(errs, "'Form.AccountNotification.Message' is required")
	}
	if t.Form.AccountNotification.StartAt.IsZero() {
		errs = append(errs, "'Form.AccountNotification.StartAt' is required")
	}
	if t.Form.AccountNotification.EndAt.IsZero() {
		errs = append(errs, "'Form.AccountNotification.EndAt' is required")
	}
	if t.Form.AccountNotification.Icon != "" && !string_utils.Include([]string{"warning", "information", "question", "error", "calendar"}, t.Form.AccountNotification.Icon) {
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
