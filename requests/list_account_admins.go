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

// ListAccountAdmins A paginated list of the admins in the account
// https://canvas.instructure.com/doc/api/admins.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.UserID (Optional) Scope the results to those with user IDs equal to any of the IDs specified here.
//
type ListAccountAdmins struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		UserID []int `json:"user_id" url:"user_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListAccountAdmins) GetMethod() string {
	return "GET"
}

func (t *ListAccountAdmins) GetURLPath() string {
	path := "accounts/{account_id}/admins"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListAccountAdmins) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListAccountAdmins) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAccountAdmins) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAccountAdmins) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAccountAdmins) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Admin, *canvasapi.PagedResource, error) {
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
	ret := []*models.Admin{}
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
