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

// GetAllOutcomeGroupsForContextAccounts
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type GetAllOutcomeGroupsForContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetAllOutcomeGroupsForContextAccounts) GetMethod() string {
	return "GET"
}

func (t *GetAllOutcomeGroupsForContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_groups"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetAllOutcomeGroupsForContextAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAllOutcomeGroupsForContextAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllOutcomeGroupsForContextAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllOutcomeGroupsForContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllOutcomeGroupsForContextAccounts) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.OutcomeGroup, *canvasapi.PagedResource, error) {
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
	ret := []*models.OutcomeGroup{}
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
