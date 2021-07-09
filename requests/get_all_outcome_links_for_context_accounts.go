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

// GetAllOutcomeLinksForContextAccounts
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.OutcomeStyle (Optional) The detail level of the outcomes. Defaults to "abbrev".
//    Specify "full" for more information.
// # Query.OutcomeGroupStyle (Optional) The detail level of the outcome groups. Defaults to "abbrev".
//    Specify "full" for more information.
//
type GetAllOutcomeLinksForContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		OutcomeStyle      string `json:"outcome_style" url:"outcome_style,omitempty"`             //  (Optional)
		OutcomeGroupStyle string `json:"outcome_group_style" url:"outcome_group_style,omitempty"` //  (Optional)
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
	return v.Encode(), nil
}

func (t *GetAllOutcomeLinksForContextAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllOutcomeLinksForContextAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllOutcomeLinksForContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllOutcomeLinksForContextAccounts) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.OutcomeLink, *canvasapi.PagedResource, error) {
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
	ret := []*models.OutcomeLink{}
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
