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

// DeleteReport Deletes a generated report instance.
// https://canvas.instructure.com/doc/api/account_reports.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.Report (Required) ID
// # Path.ID (Required) ID
//
type DeleteReport struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		Report    string `json:"report" url:"report,omitempty"`         //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *DeleteReport) GetMethod() string {
	return "DELETE"
}

func (t *DeleteReport) GetURLPath() string {
	path := "accounts/{account_id}/reports/{report}/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{report}", fmt.Sprintf("%v", t.Path.Report))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteReport) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteReport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteReport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteReport) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.Report == "" {
		errs = append(errs, "'Path.Report' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteReport) Do(c *canvasapi.Canvas) (*models.Report, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Report{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
