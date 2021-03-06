package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListGroupCategoriesForContextAccounts Returns a paginated list of group categories in a context
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type ListGroupCategoriesForContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGroupCategoriesForContextAccounts) GetMethod() string {
	return "GET"
}

func (t *ListGroupCategoriesForContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/group_categories"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListGroupCategoriesForContextAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGroupCategoriesForContextAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGroupCategoriesForContextAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGroupCategoriesForContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupCategoriesForContextAccounts) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.GroupCategory, *canvasapi.PagedResource, error) {
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
	ret := []*models.GroupCategory{}
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
