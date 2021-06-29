package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateEnrollmentTerm Update an existing enrollment term for the specified account.
// https://canvas.instructure.com/doc/api/enrollment_terms.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.EnrollmentTerm.Name (Optional) The name of the term.
// # Form.EnrollmentTerm.StartAt (Optional) The day/time the term starts.
//    Accepts times in ISO 8601 format, e.g. 2015-01-10T18:48:00Z.
// # Form.EnrollmentTerm.EndAt (Optional) The day/time the term ends.
//    Accepts times in ISO 8601 format, e.g. 2015-01-10T18:48:00Z.
// # Form.EnrollmentTerm.SISTermID (Optional) The unique SIS identifier for the term.
// # Form.EnrollmentTerm.Overrides.EnrollmentType.StartAt (Optional) The day/time the term starts, overridden for the given enrollment type.
//    *enrollment_type* can be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, or DesignerEnrollment
// # Form.EnrollmentTerm.Overrides.EnrollmentType.EndAt (Optional) The day/time the term ends, overridden for the given enrollment type.
//    *enrollment_type* can be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, or DesignerEnrollment
//
type UpdateEnrollmentTerm struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Form struct {
		EnrollmentTerm struct {
			Name      string    `json:"name" url:"name,omitempty"`               //  (Optional)
			StartAt   time.Time `json:"start_at" url:"start_at,omitempty"`       //  (Optional)
			EndAt     time.Time `json:"end_at" url:"end_at,omitempty"`           //  (Optional)
			SISTermID string    `json:"sis_term_id" url:"sis_term_id,omitempty"` //  (Optional)
			Overrides struct {
				EnrollmentType struct {
					StartAt time.Time `json:"start_at" url:"start_at,omitempty"` //  (Optional)
					EndAt   time.Time `json:"end_at" url:"end_at,omitempty"`     //  (Optional)
				} `json:"enrollment_type" url:"enrollment_type,omitempty"`
			} `json:"overrides" url:"overrides,omitempty"`
		} `json:"enrollment_term" url:"enrollment_term,omitempty"`
	} `json:"form"`
}

func (t *UpdateEnrollmentTerm) GetMethod() string {
	return "PUT"
}

func (t *UpdateEnrollmentTerm) GetURLPath() string {
	path := "accounts/{account_id}/terms/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateEnrollmentTerm) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateEnrollmentTerm) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateEnrollmentTerm) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateEnrollmentTerm) HasErrors() error {
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

func (t *UpdateEnrollmentTerm) Do(c *canvasapi.Canvas) (*models.EnrollmentTerm, error) {
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
