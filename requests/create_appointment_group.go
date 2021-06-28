package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateAppointmentGroup Create and return a new appointment group. If new_appointments are
// specified, the response will return a new_appointments array (same format
// as appointments array, see "List appointment groups" action)
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Form Parameters:
// # AppointmentGroup (Required) Array of context codes (courses, e.g. course_1) this group should be
//    linked to (1 or more). Users in the course(s) with appropriate permissions
//    will be able to sign up for this appointment group.
// # AppointmentGroup (Optional) Array of sub context codes (course sections or a single group category)
//    this group should be linked to. Used to limit the appointment group to
//    particular sections. If a group category is specified, students will sign
//    up in groups and the participant_type will be "Group" instead of "User".
// # AppointmentGroup (Required) Short title for the appointment group.
// # AppointmentGroup (Optional) Longer text description of the appointment group.
// # AppointmentGroup (Optional) Location name of the appointment group.
// # AppointmentGroup (Optional) Location address.
// # AppointmentGroup (Optional) Indicates whether this appointment group should be published (i.e. made
//    available for signup). Once published, an appointment group cannot be
//    unpublished. Defaults to false.
// # AppointmentGroup (Optional) Maximum number of participants that may register for each time slot.
//    Defaults to null (no limit).
// # AppointmentGroup (Optional) Minimum number of time slots a user must register for. If not set, users
//    do not need to sign up for any time slots.
// # AppointmentGroup (Optional) Maximum number of time slots a user may register for.
// # AppointmentGroup (Optional) Nested array of start time/end time pairs indicating time slots for this
//    appointment group. Refer to the example request.
// # AppointmentGroup (Optional) . Must be one of private, protected"private":: participants cannot see who has signed up for a particular
//                time slot
//    "protected":: participants can see who has signed up.  Defaults to
//                  "private".
//
type CreateAppointmentGroup struct {
	Form struct {
		AppointmentGroup struct {
			ContextCodes                  []string `json:"context_codes" url:"context_codes,omitempty"`                                       //  (Required)
			SubContextCodes               []string `json:"sub_context_codes" url:"sub_context_codes,omitempty"`                               //  (Optional)
			Title                         string   `json:"title" url:"title,omitempty"`                                                       //  (Required)
			Description                   string   `json:"description" url:"description,omitempty"`                                           //  (Optional)
			LocationName                  string   `json:"location_name" url:"location_name,omitempty"`                                       //  (Optional)
			LocationAddress               string   `json:"location_address" url:"location_address,omitempty"`                                 //  (Optional)
			Publish                       bool     `json:"publish" url:"publish,omitempty"`                                                   //  (Optional)
			ParticipantsPerAppointment    int64    `json:"participants_per_appointment" url:"participants_per_appointment,omitempty"`         //  (Optional)
			MinAppointmentsPerParticipant int64    `json:"min_appointments_per_participant" url:"min_appointments_per_participant,omitempty"` //  (Optional)
			MaxAppointmentsPerParticipant int64    `json:"max_appointments_per_participant" url:"max_appointments_per_participant,omitempty"` //  (Optional)
			NewAppointments               struct {
				X []string `json:"x" url:"x,omitempty"` //  (Optional)
			} `json:"new_appointments" url:"new_appointments,omitempty"`

			ParticipantVisibility string `json:"participant_visibility" url:"participant_visibility,omitempty"` //  (Optional) . Must be one of private, protected
		} `json:"appointment_group" url:"appointment_group,omitempty"`
	} `json:"form"`
}

func (t *CreateAppointmentGroup) GetMethod() string {
	return "POST"
}

func (t *CreateAppointmentGroup) GetURLPath() string {
	return ""
}

func (t *CreateAppointmentGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateAppointmentGroup) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateAppointmentGroup) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateAppointmentGroup) HasErrors() error {
	errs := []string{}
	if t.Form.AppointmentGroup.ContextCodes == nil {
		errs = append(errs, "'AppointmentGroup' is required")
	}
	if t.Form.AppointmentGroup.Title == "" {
		errs = append(errs, "'AppointmentGroup' is required")
	}
	if t.Form.AppointmentGroup.ParticipantVisibility != "" && !string_utils.Include([]string{"private", "protected"}, t.Form.AppointmentGroup.ParticipantVisibility) {
		errs = append(errs, "AppointmentGroup must be one of private, protected")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateAppointmentGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
