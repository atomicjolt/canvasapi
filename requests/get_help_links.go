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

// GetHelpLinks Returns the help links for that account
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetHelpLinks struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetHelpLinks) GetMethod() string {
	return "GET"
}

func (t *GetHelpLinks) GetURLPath() string {
	path := "accounts/{account_id}/help_links"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetHelpLinks) GetQuery() (string, error) {
	return "", nil
}

func (t *GetHelpLinks) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetHelpLinks) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetHelpLinks) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetHelpLinks) Do(c *canvasapi.Canvas) (*models.HelpLinks, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.HelpLinks{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
