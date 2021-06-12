package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListStudentGroupParticipants A paginated list of student groups that are (or may be) participating in
// this appointment group. Refer to the Groups API for the response fields.
// Returns no results for appointment groups with the "User" participant_type.
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # RegistrationStatus (Optional) . Must be one of all, registered, registeredLimits results to the a given participation status, defaults to "all"
//
type ListStudentGroupParticipants struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Query struct {
		RegistrationStatus string `json:"registration_status"` //  (Optional) . Must be one of all, registered, registered
	} `json:"query"`
}

func (t *ListStudentGroupParticipants) GetMethod() string {
	return "GET"
}

func (t *ListStudentGroupParticipants) GetURLPath() string {
	path := "appointment_groups/{id}/groups"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListStudentGroupParticipants) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListStudentGroupParticipants) GetBody() (string, error) {
	return "", nil
}

func (t *ListStudentGroupParticipants) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"all", "registered", "registered"}, t.Query.RegistrationStatus) {
		errs = append(errs, "RegistrationStatus must be one of all, registered, registered")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListStudentGroupParticipants) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}