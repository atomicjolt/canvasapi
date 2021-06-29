package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteExternalToolAccounts Remove the specified external tool
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ExternalToolID (Required) ID
//
type DeleteExternalToolAccounts struct {
	Path struct {
		AccountID      string `json:"account_id" url:"account_id,omitempty"`             //  (Required)
		ExternalToolID string `json:"external_tool_id" url:"external_tool_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteExternalToolAccounts) GetMethod() string {
	return "DELETE"
}

func (t *DeleteExternalToolAccounts) GetURLPath() string {
	path := "accounts/{account_id}/external_tools/{external_tool_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{external_tool_id}", fmt.Sprintf("%v", t.Path.ExternalToolID))
	return path
}

func (t *DeleteExternalToolAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteExternalToolAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteExternalToolAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteExternalToolAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.ExternalToolID == "" {
		errs = append(errs, "'Path.ExternalToolID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteExternalToolAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
