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

// IndexOfActiveGlobalNotificationForUser Returns a list of all global notifications in the account for the current user
// Any notifications that have been closed by the user will not be returned, unless
// a include_past parameter is passed in as true.
// https://canvas.instructure.com/doc/api/account_notifications.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # IncludePast (Optional) Include past and dismissed global announcements.
//
type IndexOfActiveGlobalNotificationForUser struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		IncludePast bool `json:"include_past"` //  (Optional)
	} `json:"query"`
}

func (t *IndexOfActiveGlobalNotificationForUser) GetMethod() string {
	return "GET"
}

func (t *IndexOfActiveGlobalNotificationForUser) GetURLPath() string {
	path := "accounts/{account_id}/account_notifications"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *IndexOfActiveGlobalNotificationForUser) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *IndexOfActiveGlobalNotificationForUser) GetBody() (string, error) {
	return "", nil
}

func (t *IndexOfActiveGlobalNotificationForUser) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *IndexOfActiveGlobalNotificationForUser) Do(c *canvasapi.Canvas) ([]*models.AccountNotification, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.AccountNotification{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
