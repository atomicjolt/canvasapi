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

// StatusOfReport Returns the status of a report.
// https://canvas.instructure.com/doc/api/account_reports.html
//
// Path Parameters:
// # AccountID (Required) ID
// # Report (Required) ID
// # ID (Required) ID
//
type StatusOfReport struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		Report    string `json:"report" url:"report,omitempty"`         //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *StatusOfReport) GetMethod() string {
	return "GET"
}

func (t *StatusOfReport) GetURLPath() string {
	path := "accounts/{account_id}/reports/{report}/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{report}", fmt.Sprintf("%v", t.Path.Report))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *StatusOfReport) GetQuery() (string, error) {
	return "", nil
}

func (t *StatusOfReport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *StatusOfReport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *StatusOfReport) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.Report == "" {
		errs = append(errs, "'Report' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *StatusOfReport) Do(c *canvasapi.Canvas) (*models.Report, error) {
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
