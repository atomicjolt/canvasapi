package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type AppointmentGroup struct {
	ID                            int64            `json:"id" url:"id,omitempty"`                                                             // The ID of the appointment group.Example: 543
	Title                         string           `json:"title" url:"title,omitempty"`                                                       // The title of the appointment group.Example: Final Presentation
	StartAt                       time.Time        `json:"start_at" url:"start_at,omitempty"`                                                 // The start of the first time slot in the appointment group.Example: 2012-07-20T15:00:00-06:00
	EndAt                         time.Time        `json:"end_at" url:"end_at,omitempty"`                                                     // The end of the last time slot in the appointment group.Example: 2012-07-20T17:00:00-06:00
	Description                   string           `json:"description" url:"description,omitempty"`                                           // The text description of the appointment group.Example: Es muy importante
	LocationName                  string           `json:"location_name" url:"location_name,omitempty"`                                       // The location name of the appointment group.Example: El Tigre Chino's office
	LocationAddress               string           `json:"location_address" url:"location_address,omitempty"`                                 // The address of the appointment group's location.Example: Room 234
	ParticipantCount              int64            `json:"participant_count" url:"participant_count,omitempty"`                               // The number of participant who have reserved slots (see include[] argument).Example: 2
	ReservedTimes                 []*Appointment   `json:"reserved_times" url:"reserved_times,omitempty"`                                     // The start and end times of slots reserved by the current user as well as the id of the calendar event for the reservation (see include[] argument).Example: {'id'=>987, 'start_at'=>'2012-07-20T15:00:00-06:00', 'end_at'=>'2012-07-20T15:00:00-06:00'}
	ContextCodes                  []string         `json:"context_codes" url:"context_codes,omitempty"`                                       // The context codes (i.e. courses) this appointment group belongs to. Only people in these courses will be eligible to sign up..Example: course_123
	SubContextCodes               []string         `json:"sub_context_codes" url:"sub_context_codes,omitempty"`                               // The sub-context codes (i.e. course sections and group categories) this appointment group is restricted to.Example: course_section_234
	WorkflowState                 string           `json:"workflow_state" url:"workflow_state,omitempty"`                                     // Current state of the appointment group ('pending', 'active' or 'deleted'). 'pending' indicates that it has not been published yet and is invisible to participants..Example: active
	RequiringAction               bool             `json:"requiring_action" url:"requiring_action,omitempty"`                                 // Boolean indicating whether the current user needs to sign up for this appointment group (i.e. it's reservable and the min_appointments_per_participant limit has not been met by this user)..Example: true
	AppointmentsCount             int64            `json:"appointments_count" url:"appointments_count,omitempty"`                             // Number of time slots in this appointment group.Example: 2
	Appointments                  []*CalendarEvent `json:"appointments" url:"appointments,omitempty"`                                         // Calendar Events representing the time slots (see include[] argument) Refer to the Calendar Events API for more information.
	NewAppointments               []*CalendarEvent `json:"new_appointments" url:"new_appointments,omitempty"`                                 // Newly created time slots (same format as appointments above). Only returned in Create/Update responses where new time slots have been added.
	MaxAppointmentsPerParticipant int64            `json:"max_appointments_per_participant" url:"max_appointments_per_participant,omitempty"` // Maximum number of time slots a user may register for, or null if no limit.Example: 1
	MinAppointmentsPerParticipant int64            `json:"min_appointments_per_participant" url:"min_appointments_per_participant,omitempty"` // Minimum number of time slots a user must register for. If not set, users do not need to sign up for any time slots.Example: 1
	ParticipantsPerAppointment    int64            `json:"participants_per_appointment" url:"participants_per_appointment,omitempty"`         // Maximum number of participants that may register for each time slot, or null if no limit.Example: 1
	ParticipantVisibility         string           `json:"participant_visibility" url:"participant_visibility,omitempty"`                     // 'private' means participants cannot see who has signed up for a particular time slot, 'protected' means that they can.Example: private
	ParticipantType               string           `json:"participant_type" url:"participant_type,omitempty"`                                 // Indicates how participants sign up for the appointment group, either as individuals ('User') or in student groups ('Group'). Related to sub_context_codes (i.e. 'Group' signups always have a single group category).Example: User
	Url                           string           `json:"url" url:"url,omitempty"`                                                           // URL for this appointment group (to update, delete, etc.).Example: https://example.com/api/v1/appointment_groups/543
	HtmlUrl                       string           `json:"html_url" url:"html_url,omitempty"`                                                 // URL for a user to view this appointment group.Example: http://example.com/appointment_groups/1
	CreatedAt                     time.Time        `json:"created_at" url:"created_at,omitempty"`                                             // When the appointment group was created.Example: 2012-07-13T10:55:20-06:00
	UpdatedAt                     time.Time        `json:"updated_at" url:"updated_at,omitempty"`                                             // When the appointment group was last updated.Example: 2012-07-13T10:55:20-06:00
}

func (t *AppointmentGroup) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"pending", "active", "deleted"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		errs = append(errs, fmt.Sprintf("expected 'WorkflowState' to be one of %v", s))
	}
	s = []string{"private", "protected"}
	if t.ParticipantVisibility != "" && !string_utils.Include(s, t.ParticipantVisibility) {
		errs = append(errs, fmt.Sprintf("expected 'ParticipantVisibility' to be one of %v", s))
	}
	s = []string{"User", "Group"}
	if t.ParticipantType != "" && !string_utils.Include(s, t.ParticipantType) {
		errs = append(errs, fmt.Sprintf("expected 'ParticipantType' to be one of %v", s))
	}
	return nil
}
