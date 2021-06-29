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
	"github.com/atomicjolt/string_utils"
)

// ListEnrollmentTerms An object with a paginated list of all of the terms in the account.
// https://canvas.instructure.com/doc/api/enrollment_terms.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.WorkflowState (Optional) . Must be one of active, deleted, allIf set, only returns terms that are in the given state.
//    Defaults to 'active'.
// # Query.Include (Optional) . Must be one of overridesArray of additional information to include.
//
//    "overrides":: term start/end dates overridden for different enrollment types
//
type ListEnrollmentTerms struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		WorkflowState []string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Optional) . Must be one of active, deleted, all
		Include       []string `json:"include" url:"include,omitempty"`               //  (Optional) . Must be one of overrides
	} `json:"query"`
}

func (t *ListEnrollmentTerms) GetMethod() string {
	return "GET"
}

func (t *ListEnrollmentTerms) GetURLPath() string {
	path := "accounts/{account_id}/terms"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListEnrollmentTerms) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListEnrollmentTerms) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEnrollmentTerms) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEnrollmentTerms) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	for _, v := range t.Query.WorkflowState {
		if v != "" && !string_utils.Include([]string{"active", "deleted", "all"}, v) {
			errs = append(errs, "WorkflowState must be one of active, deleted, all")
		}
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"overrides"}, v) {
			errs = append(errs, "Include must be one of overrides")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEnrollmentTerms) Do(c *canvasapi.Canvas) (*models.EnrollmentTermsList, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.EnrollmentTermsList{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
