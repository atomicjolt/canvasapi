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

// GetOutcomeImportStatusAccounts Get the status of an already created Outcome import. Pass 'latest' for the outcome import id
// for the latest import.
//
//   Examples:
//     curl 'https://<canvas>/api/v1/accounts/<account_id>/outcome_imports/<outcome_import_id>' \
//         -H "Authorization: Bearer <token>"
//     curl 'https://<canvas>/api/v1/courses/<course_id>/outcome_imports/<outcome_import_id>' \
//         -H "Authorization: Bearer <token>"
// https://canvas.instructure.com/doc/api/outcome_imports.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type GetOutcomeImportStatusAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *GetOutcomeImportStatusAccounts) GetMethod() string {
	return "GET"
}

func (t *GetOutcomeImportStatusAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_imports/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetOutcomeImportStatusAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetOutcomeImportStatusAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetOutcomeImportStatusAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetOutcomeImportStatusAccounts) HasErrors() error {
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

func (t *GetOutcomeImportStatusAccounts) Do(c *canvasapi.Canvas) (*models.OutcomeImport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeImport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
