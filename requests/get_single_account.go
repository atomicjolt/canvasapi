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

// GetSingleAccount Retrieve information on an individual account, given by id or sis
// sis_account_id.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type GetSingleAccount struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleAccount) GetMethod() string {
	return "GET"
}

func (t *GetSingleAccount) GetURLPath() string {
	path := "accounts/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleAccount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleAccount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleAccount) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleAccount) Do(c *canvasapi.Canvas) (*models.Account, error) {
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
