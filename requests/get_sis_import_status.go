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

// GetSISImportStatus Get the status of an already created SIS import.
//
//   Examples:
//     curl https://<canvas>/api/v1/accounts/<account_id>/sis_imports/<sis_import_id> \
//         -H 'Authorization: Bearer <token>'
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type GetSISImportStatus struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *GetSISImportStatus) GetMethod() string {
	return "GET"
}

func (t *GetSISImportStatus) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSISImportStatus) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSISImportStatus) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSISImportStatus) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSISImportStatus) HasErrors() error {
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

func (t *GetSISImportStatus) Do(c *canvasapi.Canvas) (*models.SISImport, error) {
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
