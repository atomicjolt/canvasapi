package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// RemoveToolFromRceFavorites Remove the specified external tool from a preferred location in the RCE
// for the given account
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type RemoveToolFromRceFavorites struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *RemoveToolFromRceFavorites) GetMethod() string {
	return "DELETE"
}

func (t *RemoveToolFromRceFavorites) GetURLPath() string {
	path := "accounts/{account_id}/external_tools/rce_favorites/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RemoveToolFromRceFavorites) GetQuery() (string, error) {
	return "", nil
}

func (t *RemoveToolFromRceFavorites) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveToolFromRceFavorites) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveToolFromRceFavorites) HasErrors() error {
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

func (t *RemoveToolFromRceFavorites) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
