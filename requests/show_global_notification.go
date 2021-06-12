package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowGlobalNotification Returns a global notification for the current user
// A notification that has been closed by the user will not be returned
// https://canvas.instructure.com/doc/api/account_notifications.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
type ShowGlobalNotification struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`
}

func (t *ShowGlobalNotification) GetMethod() string {
	return "GET"
}

func (t *ShowGlobalNotification) GetURLPath() string {
	path := "accounts/{account_id}/account_notifications/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowGlobalNotification) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowGlobalNotification) GetBody() (string, error) {
	return "", nil
}

func (t *ShowGlobalNotification) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowGlobalNotification) Do(c *canvasapi.Canvas) (*models.AccountNotification, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AccountNotification{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
