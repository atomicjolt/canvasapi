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

// IndexOfReports Shows all reports that have been run for the account of a specific type.
// https://canvas.instructure.com/doc/api/account_reports.html
//
// Path Parameters:
// # AccountID (Required) ID
// # Report (Required) ID
//
type IndexOfReports struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		Report    string `json:"report" url:"report,omitempty"`         //  (Required)
	} `json:"path"`
}

func (t *IndexOfReports) GetMethod() string {
	return "GET"
}

func (t *IndexOfReports) GetURLPath() string {
	path := "accounts/{account_id}/reports/{report}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{report}", fmt.Sprintf("%v", t.Path.Report))
	return path
}

func (t *IndexOfReports) GetQuery() (string, error) {
	return "", nil
}

func (t *IndexOfReports) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *IndexOfReports) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *IndexOfReports) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.Report == "" {
		errs = append(errs, "'Report' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *IndexOfReports) Do(c *canvasapi.Canvas) ([]*models.Report, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Report{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
