package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # EnrollmentTerm (Optional) The name of the term.
// # EnrollmentTerm (Optional) The day/time the term starts.
//    Accepts times in ISO 8601 format, e.g. 2015-01-10T18:48:00Z.
// # EnrollmentTerm (Optional) The day/time the term ends.
//    Accepts times in ISO 8601 format, e.g. 2015-01-10T18:48:00Z.
// # EnrollmentTerm (Optional) The unique SIS identifier for the term.
// # EnrollmentTerm (Optional) The day/time the term starts, overridden for the given enrollment type.
//    *enrollment_type* can be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, or DesignerEnrollment
// # EnrollmentTerm (Optional) The day/time the term ends, overridden for the given enrollment type.
//    *enrollment_type* can be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, or DesignerEnrollment
//
type UpdateEnrollmentTerm struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`

	Form struct {
		EnrollmentTerm struct {
			Name      string    `json:"name"`        //  (Optional)
			StartAt   time.Time `json:"start_at"`    //  (Optional)
			EndAt     time.Time `json:"end_at"`      //  (Optional)
			SISTermID string    `json:"sis_term_id"` //  (Optional)
			Overrides struct {
				EnrollmentType struct {
					StartAt time.Time `json:"start_at"` //  (Optional)
					EndAt   time.Time `json:"end_at"`   //  (Optional)
				} `json:"enrollment_type"`
			} `json:"overrides"`
		} `json:"enrollment_term"`
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

func (t *UpdateEnrollmentTerm) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateEnrollmentTerm) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
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
