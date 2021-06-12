package requests

import (
	"fmt"
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
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # AccountNotification (Optional) The subject of the notification.
// # AccountNotification (Optional) The message body of the notification.
// # AccountNotification (Optional) The start date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # AccountNotification (Optional) The end date and time of the notification in ISO8601 format.
//    e.g. 2014-01-01T01:00Z
// # AccountNotification (Optional) . Must be one of warning, information, question, error, calendarThe icon to display with the notification.
// # AccountNotificationRoles (Optional) The role(s) to send global notification to.  Note:  ommitting this field will send to everyone
//    Example:
//      account_notification_roles: ["StudentEnrollment", "TeacherEnrollment"]
//
type UpdateGlobalNotification struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`

	Form struct {
		AccountNotification struct {
			Subject string    `json:"subject"`  //  (Optional)
			Message string    `json:"message"`  //  (Optional)
			StartAt time.Time `json:"start_at"` //  (Optional)
			EndAt   time.Time `json:"end_at"`   //  (Optional)
			Icon    string    `json:"icon"`     //  (Optional) . Must be one of warning, information, question, error, calendar
		} `json:"account_notification"`

		AccountNotificationRoles []string `json:"account_notification_roles"` //  (Optional)
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

func (t *UpdateGlobalNotification) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateGlobalNotification) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"warning", "information", "question", "error", "calendar"}, t.Form.AccountNotification.Icon) {
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
