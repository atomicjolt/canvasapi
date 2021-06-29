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

// UpdateAppointmentGroup Update and return an appointment group. If new_appointments are specified,
// the response will return a new_appointments array (same format as
// appointments array, see "List appointment groups" action).
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.AppointmentGroup.ContextCodes (Required) Array of context codes (courses, e.g. course_1) this group should be
//    linked to (1 or more). Users in the course(s) with appropriate permissions
//    will be able to sign up for this appointment group.
// # Form.AppointmentGroup.SubContextCodes (Optional) Array of sub context codes (course sections or a single group category)
//    this group should be linked to. Used to limit the appointment group to
//    particular sections. If a group category is specified, students will sign
//    up in groups and the participant_type will be "Group" instead of "User".
// # Form.AppointmentGroup.Title (Optional) Short title for the appointment group.
// # Form.AppointmentGroup.Description (Optional) Longer text description of the appointment group.
// # Form.AppointmentGroup.LocationName (Optional) Location name of the appointment group.
// # Form.AppointmentGroup.LocationAddress (Optional) Location address.
// # Form.AppointmentGroup.Publish (Optional) Indicates whether this appointment group should be published (i.e. made
//    available for signup). Once published, an appointment group cannot be
//    unpublished. Defaults to false.
// # Form.AppointmentGroup.ParticipantsPerAppointment (Optional) Maximum number of participants that may register for each time slot.
//    Defaults to null (no limit).
// # Form.AppointmentGroup.MinAppointmentsPerParticipant (Optional) Minimum number of time slots a user must register for. If not set, users
//    do not need to sign up for any time slots.
// # Form.AppointmentGroup.MaxAppointmentsPerParticipant (Optional) Maximum number of time slots a user may register for.
// # Form.AppointmentGroup (Optional) Nested array of start time/end time pairs indicating time slots for this
//    appointment group. Refer to the example request.
// # Form.AppointmentGroup.ParticipantVisibility (Optional) . Must be one of private, protected"private":: participants cannot see who has signed up for a particular
//                time slot
//    "protected":: participants can see who has signed up. Defaults to "private".
//
type UpdateAppointmentGroup struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		AppointmentGroup struct {
			ContextCodes                  []string `json:"context_codes" url:"context_codes,omitempty"`                                       //  (Required)
			SubContextCodes               []string `json:"sub_context_codes" url:"sub_context_codes,omitempty"`                               //  (Optional)
			Title                         string   `json:"title" url:"title,omitempty"`                                                       //  (Optional)
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

func (t *UpdateAppointmentGroup) GetMethod() string {
	return "PUT"
}

func (t *UpdateAppointmentGroup) GetURLPath() string {
	path := "appointment_groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateAppointmentGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateAppointmentGroup) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateAppointmentGroup) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateAppointmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.AppointmentGroup.ContextCodes == nil {
		errs = append(errs, "'Form.AppointmentGroup.ContextCodes' is required")
	}
	if t.Form.AppointmentGroup.ParticipantVisibility != "" && !string_utils.Include([]string{"private", "protected"}, t.Form.AppointmentGroup.ParticipantVisibility) {
		errs = append(errs, "AppointmentGroup must be one of private, protected")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateAppointmentGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
