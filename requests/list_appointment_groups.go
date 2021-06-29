package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListAppointmentGroups Retrieve the paginated list of appointment groups that can be reserved or
// managed by the current user.
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Query Parameters:
// # Query.Scope (Optional) . Must be one of reservable, manageableDefaults to "reservable"
// # Query.ContextCodes (Optional) Array of context codes used to limit returned results.
// # Query.IncludePastAppointments (Optional) Defaults to false. If true, includes past appointment groups
// # Query.Include (Optional) . Must be one of appointments, child_events, participant_count, reserved_times, all_context_codesArray of additional information to include.
//
//    "appointments":: calendar event time slots for this appointment group
//    "child_events":: reservations of those time slots
//    "participant_count":: number of reservations
//    "reserved_times":: the event id, start time and end time of reservations
//                       the current user has made)
//    "all_context_codes":: all context codes associated with this appointment group
//
type ListAppointmentGroups struct {
	Query struct {
		Scope                   string   `json:"scope" url:"scope,omitempty"`                                         //  (Optional) . Must be one of reservable, manageable
		ContextCodes            []string `json:"context_codes" url:"context_codes,omitempty"`                         //  (Optional)
		IncludePastAppointments bool     `json:"include_past_appointments" url:"include_past_appointments,omitempty"` //  (Optional)
		Include                 []string `json:"include" url:"include,omitempty"`                                     //  (Optional) . Must be one of appointments, child_events, participant_count, reserved_times, all_context_codes
	} `json:"query"`
}

func (t *ListAppointmentGroups) GetMethod() string {
	return "GET"
}

func (t *ListAppointmentGroups) GetURLPath() string {
	return ""
}

func (t *ListAppointmentGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListAppointmentGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAppointmentGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAppointmentGroups) HasErrors() error {
	errs := []string{}
	if t.Query.Scope != "" && !string_utils.Include([]string{"reservable", "manageable"}, t.Query.Scope) {
		errs = append(errs, "Scope must be one of reservable, manageable")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"appointments", "child_events", "participant_count", "reserved_times", "all_context_codes"}, v) {
			errs = append(errs, "Include must be one of appointments, child_events, participant_count, reserved_times, all_context_codes")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAppointmentGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
