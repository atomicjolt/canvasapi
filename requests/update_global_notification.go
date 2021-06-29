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

// UpdateGlobalNotification Update global notification for an account.
// https://canvas.instructure.com/doc/api/account_notifications.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.AccountNotification.Subject (Optional) The subject of the notification.
// # Form.AccountNotification.Message (Optional) The message body of the notification.
// # Form.AccountNotification.StartAt (Optional) The start date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # Form.AccountNotification.EndAt (Optional) The end date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # Form.AccountNotification.Icon (Optional) . Must be one of warning, information, question, error, calendarThe icon to display with the notification.
// # Form.AccountNotificationRoles (Optional) The role(s) to send global notification to.  Note:  ommitting this field will send to everyone
//    Example:
//      account_notification_roles: ["StudentEnrollment", "TeacherEnrollment"]
//
type UpdateGlobalNotification struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Form struct {
		AccountNotification struct {
			Subject string    `json:"subject" url:"subject,omitempty"`   //  (Optional)
			Message string    `json:"message" url:"message,omitempty"`   //  (Optional)
			StartAt time.Time `json:"start_at" url:"start_at,omitempty"` //  (Optional)
			EndAt   time.Time `json:"end_at" url:"end_at,omitempty"`     //  (Optional)
			Icon    string    `json:"icon" url:"icon,omitempty"`         //  (Optional) . Must be one of warning, information, question, error, calendar
		} `json:"account_notification" url:"account_notification,omitempty"`

		AccountNotificationRoles []string `json:"account_notification_roles" url:"account_notification_roles,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateGlobalNotification) GetMethod() string {
	return "PUT"
}

func (t *UpdateGlobalNotification) GetURLPath() string {
	path := "accounts/{account_id}/account_notifications/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateGlobalNotification) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateGlobalNotification) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateGlobalNotification) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateGlobalNotification) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.AccountNotification.Icon != "" && !string_utils.Include([]string{"warning", "information", "question", "error", "calendar"}, t.Form.AccountNotification.Icon) {
		errs = append(errs, "AccountNotification must be one of warning, information, question, error, calendar")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateGlobalNotification) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
