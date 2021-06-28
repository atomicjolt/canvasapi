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

// AbortSISImport Abort a SIS import that has not completed.
//
// Aborting a sis batch that is running can take some time for every process to
// see the abort event. Subsequent sis batches begin to process 10 minutes
// after the abort to allow each process to clean up properly.
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
type AbortSISImport struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *AbortSISImport) GetMethod() string {
	return "PUT"
}

func (t *AbortSISImport) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports/{id}/abort"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AbortSISImport) GetQuery() (string, error) {
	return "", nil
}

func (t *AbortSISImport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *AbortSISImport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *AbortSISImport) HasErrors() error {
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

func (t *AbortSISImport) Do(c *canvasapi.Canvas) (*models.SISImport, error) {
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
