package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteAuthenticationProvider Delete the config
// https://canvas.instructure.com/doc/api/authentication_providers.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
type DeleteAuthenticationProvider struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`
}

func (t *DeleteAuthenticationProvider) GetMethod() string {
	return "DELETE"
}

func (t *DeleteAuthenticationProvider) GetURLPath() string {
	path := "accounts/{account_id}/authentication_providers/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteAuthenticationProvider) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteAuthenticationProvider) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteAuthenticationProvider) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteAuthenticationProvider) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
