package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListExternalToolsAccounts Returns the paginated list of external tools for the current context.
// See the get request docs for a single tool for a list of properties on an external tool.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.SearchTerm (Optional) The partial name of the tools to match and return.
// # Query.Selectable (Optional) If true, then only tools that are meant to be selectable are returned
// # Query.IncludeParents (Optional) If true, then include tools installed in all accounts above the current context
//
type ListExternalToolsAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm     string `json:"search_term" url:"search_term,omitempty"`         //  (Optional)
		Selectable     bool   `json:"selectable" url:"selectable,omitempty"`           //  (Optional)
		IncludeParents bool   `json:"include_parents" url:"include_parents,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListExternalToolsAccounts) GetMethod() string {
	return "GET"
}

func (t *ListExternalToolsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/external_tools"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListExternalToolsAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListExternalToolsAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListExternalToolsAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListExternalToolsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListExternalToolsAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
