package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListUserParticipants A paginated list of users that are (or may be) participating in this
// appointment group.  Refer to the Users API for the response fields. Returns
// no results for appointment groups with the "Group" participant_type.
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # RegistrationStatus (Optional) . Must be one of all, registered, registeredLimits results to the a given participation status, defaults to "all"
//
type ListUserParticipants struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Query struct {
		RegistrationStatus string `json:"registration_status"` //  (Optional) . Must be one of all, registered, registered
	} `json:"query"`
}

func (t *ListUserParticipants) GetMethod() string {
	return "GET"
}

func (t *ListUserParticipants) GetURLPath() string {
	path := "appointment_groups/{id}/users"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListUserParticipants) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListUserParticipants) GetBody() (string, error) {
	return "", nil
}

func (t *ListUserParticipants) HasErrors() error {
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

func (t *ListUserParticipants) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
