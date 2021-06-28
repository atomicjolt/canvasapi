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

// GetCurrentImportingSISImport Returns the SIS imports that are currently processing for an account. If no
// imports are running, will return an empty array.
//
// Example:
//   curl https://<canvas>/api/v1/accounts/<account_id>/sis_imports/importing \
//     -H 'Authorization: Bearer <token>'
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetCurrentImportingSISImport struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetCurrentImportingSISImport) GetMethod() string {
	return "GET"
}

func (t *GetCurrentImportingSISImport) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports/importing"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetCurrentImportingSISImport) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCurrentImportingSISImport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCurrentImportingSISImport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCurrentImportingSISImport) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCurrentImportingSISImport) Do(c *canvasapi.Canvas) (*models.SISImport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.SISImport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
