package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// EditExternalToolAccounts Update the specified external tool. Uses same parameters as create
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ExternalToolID (Required) ID
//
type EditExternalToolAccounts struct {
	Path struct {
		AccountID      string `json:"account_id"`       //  (Required)
		ExternalToolID string `json:"external_tool_id"` //  (Required)
	} `json:"path"`
}

func (t *EditExternalToolAccounts) GetMethod() string {
	return "PUT"
}

func (t *EditExternalToolAccounts) GetURLPath() string {
	path := "accounts/{account_id}/external_tools/{external_tool_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{external_tool_id}", fmt.Sprintf("%v", t.Path.ExternalToolID))
	return path
}

func (t *EditExternalToolAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *EditExternalToolAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *EditExternalToolAccounts) HasErrors() error {
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

func (t *EditExternalToolAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
