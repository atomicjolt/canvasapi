package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetAllOutcomeGroupsForContextAccounts
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetAllOutcomeGroupsForContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
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

func (t *GetAllOutcomeGroupsForContextAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *GetAllOutcomeGroupsForContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllOutcomeGroupsForContextAccounts) Do(c *canvasapi.Canvas) ([]*models.OutcomeGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.OutcomeGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
