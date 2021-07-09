package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.IncludePast (Optional) Include past and dismissed global announcements.
//
type IndexOfActiveGlobalNotificationForUser struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		IncludePast bool `json:"include_past" url:"include_past,omitempty"` //  (Optional)
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
	return v.Encode(), nil
}

func (t *IndexOfActiveGlobalNotificationForUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *IndexOfActiveGlobalNotificationForUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *IndexOfActiveGlobalNotificationForUser) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *IndexOfActiveGlobalNotificationForUser) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.AccountNotification, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.AccountNotification{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
