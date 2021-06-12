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

// GetAllOutcomeLinksForContextAccounts
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # OutcomeStyle (Optional) The detail level of the outcomes. Defaults to "abbrev".
//    Specify "full" for more information.
// # OutcomeGroupStyle (Optional) The detail level of the outcome groups. Defaults to "abbrev".
//    Specify "full" for more information.
//
type GetAllOutcomeLinksForContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		OutcomeStyle      string `json:"outcome_style"`       //  (Optional)
		OutcomeGroupStyle string `json:"outcome_group_style"` //  (Optional)
	} `json:"query"`
}

func (t *GetAllOutcomeLinksForContextAccounts) GetMethod() string {
	return "GET"
}

func (t *GetAllOutcomeLinksForContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_group_links"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetAllOutcomeLinksForContextAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetAllOutcomeLinksForContextAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *GetAllOutcomeLinksForContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllOutcomeLinksForContextAccounts) Do(c *canvasapi.Canvas) ([]*models.OutcomeLink, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.OutcomeLink{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}