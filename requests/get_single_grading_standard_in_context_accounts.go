package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleGradingStandardInContextAccounts Returns a grading standard for the given context that is visible to the user.
// https://canvas.instructure.com/doc/api/grading_standards.html
//
// Path Parameters:
// # AccountID (Required) ID
// # GradingStandardID (Required) ID
//
type GetSingleGradingStandardInContextAccounts struct {
	Path struct {
		AccountID         string `json:"account_id"`          //  (Required)
		GradingStandardID string `json:"grading_standard_id"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleGradingStandardInContextAccounts) GetMethod() string {
	return "GET"
}

func (t *GetSingleGradingStandardInContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/grading_standards/{grading_standard_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{grading_standard_id}", fmt.Sprintf("%v", t.Path.GradingStandardID))
	return path
}

func (t *GetSingleGradingStandardInContextAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleGradingStandardInContextAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleGradingStandardInContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.GradingStandardID == "" {
		errs = append(errs, "'GradingStandardID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGradingStandardInContextAccounts) Do(c *canvasapi.Canvas) (*models.GradingStandard, error) {
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
