package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteSubAccount Cannot delete an account with active courses or active sub_accounts.
// Cannot delete a root_account
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type DeleteSubAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *DeleteSubAccount) GetMethod() string {
	return "DELETE"
}

func (t *DeleteSubAccount) GetURLPath() string {
	path := "accounts/{account_id}/sub_accounts/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteSubAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteSubAccount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteSubAccount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteSubAccount) HasErrors() error {
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

func (t *DeleteSubAccount) Do(c *canvasapi.Canvas) (*models.Account, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Account{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
