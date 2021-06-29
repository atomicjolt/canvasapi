package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSingleAppointmentGroup Returns information for a single appointment group
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of child_events, appointments, all_context_codesArray of additional information to include. See include[] argument of
//    "List appointment groups" action.
//
//    "child_events":: reservations of time slots time slots
//    "appointments":: will always be returned
//    "all_context_codes":: all context codes associated with this appointment group
//
type GetSingleAppointmentGroup struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of child_events, appointments, all_context_codes
	} `json:"query"`
}

func (t *GetSingleAppointmentGroup) GetMethod() string {
	return "GET"
}

func (t *GetSingleAppointmentGroup) GetURLPath() string {
	path := "appointment_groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleAppointmentGroup) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetSingleAppointmentGroup) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleAppointmentGroup) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleAppointmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"child_events", "appointments", "all_context_codes"}, v) {
			errs = append(errs, "Include must be one of child_events, appointments, all_context_codes")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleAppointmentGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
