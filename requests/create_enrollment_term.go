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

// CreateEnrollmentTerm Create a new enrollment term for the specified account.
// https://canvas.instructure.com/doc/api/enrollment_terms.html
//
// Path Parameters:
// # AccountID (Required) ID
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
type CreateEnrollmentTerm struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
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

func (t *CreateEnrollmentTerm) GetMethod() string {
	return "POST"
}

func (t *CreateEnrollmentTerm) GetURLPath() string {
	path := "accounts/{account_id}/terms"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateEnrollmentTerm) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateEnrollmentTerm) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateEnrollmentTerm) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateEnrollmentTerm) Do(c *canvasapi.Canvas) (*models.EnrollmentTerm, error) {
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
