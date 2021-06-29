package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListLinkedOutcomesAccounts A paginated list of the immediate OutcomeLink children of the outcome group.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.OutcomeStyle (Optional) The detail level of the outcomes. Defaults to "abbrev".
//    Specify "full" for more information.
//
type ListLinkedOutcomesAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Query struct {
		OutcomeStyle string `json:"outcome_style" url:"outcome_style,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListLinkedOutcomesAccounts) GetMethod() string {
	return "GET"
}

func (t *ListLinkedOutcomesAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_groups/{id}/outcomes"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListLinkedOutcomesAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListLinkedOutcomesAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLinkedOutcomesAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLinkedOutcomesAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLinkedOutcomesAccounts) Do(c *canvasapi.Canvas) ([]*models.OutcomeLink, error) {
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
