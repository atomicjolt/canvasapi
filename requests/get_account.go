package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetAccount Retrieve information on an individual account, given by local or global ID.
// https://canvas.instructure.com/doc/api/accounts_(lti).html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetAccount struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *GetAccount) GetMethod() string {
	return "GET"
}

func (t *GetAccount) GetURLPath() string {
	path := "/lti/accounts/{account_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAccount) GetBody() (string, error) {
	return "", nil
}

func (t *GetAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAccount) Do(c *canvasapi.Canvas) (*models.Account, error) {
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
