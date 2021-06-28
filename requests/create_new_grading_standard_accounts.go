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

// CreateNewGradingStandardAccounts Create a new grading standard
//
// If grading_scheme_entry arguments are omitted, then a default grading scheme
// will be set. The default scheme is as follows:
//      "A" : 94,
//      "A-" : 90,
//      "B+" : 87,
//      "B" : 84,
//      "B-" : 80,
//      "C+" : 77,
//      "C" : 74,
//      "C-" : 70,
//      "D+" : 67,
//      "D" : 64,
//      "D-" : 61,
//      "F" : 0,
// https://canvas.instructure.com/doc/api/grading_standards.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # Title (Required) The title for the Grading Standard.
// # GradingSchemeEntry (Required) The name for an entry value within a GradingStandard that describes the range of the value
//    e.g. A-
// # GradingSchemeEntry (Required) The value for the name of the entry within a GradingStandard.
//    The entry represents the lower bound of the range for the entry.
//    This range includes the value up to the next entry in the GradingStandard,
//    or 100 if there is no upper bound. The lowest value will have a lower bound range of 0.
//    e.g. 93
//
type CreateNewGradingStandardAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Title              string `json:"title" url:"title,omitempty"` //  (Required)
		GradingSchemeEntry struct {
			Name  []string `json:"name" url:"name,omitempty"`   //  (Required)
			Value []int64  `json:"value" url:"value,omitempty"` //  (Required)
		} `json:"grading_scheme_entry" url:"grading_scheme_entry,omitempty"`
	} `json:"form"`
}

func (t *CreateNewGradingStandardAccounts) GetMethod() string {
	return "POST"
}

func (t *CreateNewGradingStandardAccounts) GetURLPath() string {
	path := "accounts/{account_id}/grading_standards"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateNewGradingStandardAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateNewGradingStandardAccounts) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateNewGradingStandardAccounts) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateNewGradingStandardAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.Title == "" {
		errs = append(errs, "'Title' is required")
	}
	if t.Form.GradingSchemeEntry.Name == nil {
		errs = append(errs, "'GradingSchemeEntry' is required")
	}
	if t.Form.GradingSchemeEntry.Value == nil {
		errs = append(errs, "'GradingSchemeEntry' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateNewGradingStandardAccounts) Do(c *canvasapi.Canvas) (*models.GradingStandard, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.GradingStandard{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
