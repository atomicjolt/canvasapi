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

// EnrollmentByID Get an Enrollment object by Enrollment ID
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) The ID of the enrollment object
//
type EnrollmentByID struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        int64  `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *EnrollmentByID) GetMethod() string {
	return "GET"
}

func (t *EnrollmentByID) GetURLPath() string {
	path := "accounts/{account_id}/enrollments/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EnrollmentByID) GetQuery() (string, error) {
	return "", nil
}

func (t *EnrollmentByID) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *EnrollmentByID) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *EnrollmentByID) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EnrollmentByID) Do(c *canvasapi.Canvas) (*models.Enrollment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Enrollment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
