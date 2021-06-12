package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// AddToolToRceFavorites Add the specified editor_button external tool to a preferred location in the RCE
// for courses in the given account and its subaccounts (if the subaccounts
// haven't set their own RCE Favorites). Cannot set more than 2 RCE Favorites.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
type AddToolToRceFavorites struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`
}

func (t *AddToolToRceFavorites) GetMethod() string {
	return "POST"
}

func (t *AddToolToRceFavorites) GetURLPath() string {
	path := "accounts/{account_id}/external_tools/rce_favorites/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AddToolToRceFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *AddToolToRceFavorites) GetBody() (string, error) {
	return "", nil
}

func (t *AddToolToRceFavorites) HasErrors() error {
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

func (t *AddToolToRceFavorites) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
