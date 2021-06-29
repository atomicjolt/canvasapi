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

// RetrieveEnrollmentTerm Retrieves the details for an enrollment term in the account. Includes overrides by default.
// https://canvas.instructure.com/doc/api/enrollment_terms.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type RetrieveEnrollmentTerm struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *RetrieveEnrollmentTerm) GetMethod() string {
	return "GET"
}

func (t *RetrieveEnrollmentTerm) GetURLPath() string {
	path := "accounts/{account_id}/terms/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RetrieveEnrollmentTerm) GetQuery() (string, error) {
	return "", nil
}

func (t *RetrieveEnrollmentTerm) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RetrieveEnrollmentTerm) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RetrieveEnrollmentTerm) HasErrors() error {
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

func (t *RetrieveEnrollmentTerm) Do(c *canvasapi.Canvas) (*models.EnrollmentTerm, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.EnrollmentTerm{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
