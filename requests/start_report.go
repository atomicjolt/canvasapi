package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// StartReport Generates a report instance for the account. Note that "report" in the
// request must match one of the available report names. To fetch a list of
// available report names and parameters for each report (including whether or
// not those parameters are required), see
// {api:AccountReportsController#available_reports List Available Reports}.
// https://canvas.instructure.com/doc/api/account_reports.html
//
// Path Parameters:
// # AccountID (Required) ID
// # Report (Required) ID
//
// Form Parameters:
// # Parameters (Optional) The parameters will vary for each report. To fetch a list
//    of available parameters for each report, see {api:AccountReportsController#available_reports List Available Reports}.
//    A few example parameters have been provided below. Note that the example
//    parameters provided below may not be valid for every report.
// # Parameters (Optional) The id of the course to report on.
//    Note: this parameter has been listed to serve as an example and may not be
//    valid for every report.
// # Parameters (Optional) If true, user data will be included. If
//    false, user data will be omitted. Note: this parameter has been listed to
//    serve as an example and may not be valid for every report.
//
type StartReport struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		Report    string `json:"report" url:"report,omitempty"`         //  (Required)
	} `json:"path"`

	Form struct {
		Parameters string `json:"parameters" url:"parameters,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *StartReport) GetMethod() string {
	return "POST"
}

func (t *StartReport) GetURLPath() string {
	path := "accounts/{account_id}/reports/{report}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{report}", fmt.Sprintf("%v", t.Path.Report))
	return path
}

func (t *StartReport) GetQuery() (string, error) {
	return "", nil
}

func (t *StartReport) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *StartReport) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *StartReport) HasErrors() error {
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

func (t *StartReport) Do(c *canvasapi.Canvas) (*models.Report, error) {
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
