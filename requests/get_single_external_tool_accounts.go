package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSingleExternalToolAccounts Returns the specified external tool.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ExternalToolID (Required) ID
//
type GetSingleExternalToolAccounts struct {
	Path struct {
		AccountID      string `json:"account_id" url:"account_id,omitempty"`             //  (Required)
		ExternalToolID string `json:"external_tool_id" url:"external_tool_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleExternalToolAccounts) GetMethod() string {
	return "GET"
}

func (t *GetSingleExternalToolAccounts) GetURLPath() string {
	path := "accounts/{account_id}/external_tools/{external_tool_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{external_tool_id}", fmt.Sprintf("%v", t.Path.ExternalToolID))
	return path
}

func (t *GetSingleExternalToolAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleExternalToolAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleExternalToolAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleExternalToolAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ExternalToolID == "" {
		errs = append(errs, "'ExternalToolID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleExternalToolAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
