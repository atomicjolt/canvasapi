package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListGradingStandardsAvailableInContextAccounts Returns the paginated list of grading standards for the given context that are visible to the user.
// https://canvas.instructure.com/doc/api/grading_standards.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
type ListGradingStandardsAvailableInContextAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGradingStandardsAvailableInContextAccounts) GetMethod() string {
	return "GET"
}

func (t *ListGradingStandardsAvailableInContextAccounts) GetURLPath() string {
	path := "accounts/{account_id}/grading_standards"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListGradingStandardsAvailableInContextAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGradingStandardsAvailableInContextAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGradingStandardsAvailableInContextAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGradingStandardsAvailableInContextAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGradingStandardsAvailableInContextAccounts) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.GradingStandard, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.GradingStandard{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
