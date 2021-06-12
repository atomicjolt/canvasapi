package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type AppointmentGroup struct {
	ID                            int64            `json:"id"`                               // The ID of the appointment group.Example: 543
	Title                         string           `json:"title"`                            // The title of the appointment group.Example: Final Presentation
	StartAt                       time.Time        `json:"start_at"`                         // The start of the first time slot in the appointment group.Example: 2012-07-20T15:00:00-06:00
	EndAt                         time.Time        `json:"end_at"`                           // The end of the last time slot in the appointment group.Example: 2012-07-20T17:00:00-06:00
	Description                   string           `json:"description"`                      // The text description of the appointment group.Example: Es muy importante
	LocationName                  string           `json:"location_name"`                    // The location name of the appointment group.Example: El Tigre Chino's office
	LocationAddress               string           `json:"location_address"`                 // The address of the appointment group's location.Example: Room 234
	ParticipantCount              int64            `json:"participant_count"`                // The number of participant who have reserved slots (see include[] argument).Example: 2
	ReservedTimes                 []*Appointment   `json:"reserved_times"`                   // The start and end times of slots reserved by the current user as well as the id of the calendar event for the reservation (see include[] argument).Example: {'id'=>987, 'start_at'=>'2012-07-20T15:00:00-06:00', 'end_at'=>'2012-07-20T15:00:00-06:00'}
	ContextCodes                  []string         `json:"context_codes"`                    // The context codes (i.e. courses) this appointment group belongs to. Only people in these courses will be eligible to sign up..Example: course_123
	SubContextCodes               []int64          `json:"sub_context_codes"`                // The sub-context codes (i.e. course sections and group categories) this appointment group is restricted to.Example: course_section_234
	WorkflowState                 string           `json:"workflow_state"`                   // Current state of the appointment group ('pending', 'active' or 'deleted'). 'pending' indicates that it has not been published yet and is invisible to participants..Example: active
	RequiringAction               bool             `json:"requiring_action"`                 // Boolean indicating whether the current user needs to sign up for this appointment group (i.e. it's reservable and the min_appointments_per_participant limit has not been met by this user)..Example: true
	AppointmentsCount             int64            `json:"appointments_count"`               // Number of time slots in this appointment group.Example: 2
	Appointments                  []*CalendarEvent `json:"appointments"`                     // Calendar Events representing the time slots (see include[] argument) Refer to the Calendar Events API for more information.
	NewAppointments               []*CalendarEvent `json:"new_appointments"`                 // Newly created time slots (same format as appointments above). Only returned in Create/Update responses where new time slots have been added.
	MaxAppointmentsPerParticipant int64            `json:"max_appointments_per_participant"` // Maximum number of time slots a user may register for, or null if no limit.Example: 1
	MinAppointmentsPerParticipant int64            `json:"min_appointments_per_participant"` // Minimum number of time slots a user must register for. If not set, users do not need to sign up for any time slots.Example: 1
	ParticipantsPerAppointment    int64            `json:"participants_per_appointment"`     // Maximum number of participants that may register for each time slot, or null if no limit.Example: 1
	ParticipantVisibility         string           `json:"participant_visibility"`           // 'private' means participants cannot see who has signed up for a particular time slot, 'protected' means that they can.Example: private
	ParticipantType               string           `json:"participant_type"`                 // Indicates how participants sign up for the appointment group, either as individuals ('User') or in student groups ('Group'). Related to sub_context_codes (i.e. 'Group' signups always have a single group category).Example: User
	Url                           string           `json:"url"`                              // URL for this appointment group (to update, delete, etc.).Example: https://example.com/api/v1/appointment_groups/543
	HtmlUrl                       string           `json:"html_url"`                         // URL for a user to view this appointment group.Example: http://example.com/appointment_groups/1
	CreatedAt                     time.Time        `json:"created_at"`                       // When the appointment group was created.Example: 2012-07-13T10:55:20-06:00
	UpdatedAt                     time.Time        `json:"updated_at"`                       // When the appointment group was last updated.Example: 2012-07-13T10:55:20-06:00
}

func (t *AppointmentGroup) HasError() error {
	var s []string
	s = []string{"pending", "active", "deleted"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	s = []string{"private", "protected"}
	if !string_utils.Include(s, t.ParticipantVisibility) {
		return fmt.Errorf("expected 'participant_visibility' to be one of %v", s)
	}
	s = []string{"User", "Group"}
	if !string_utils.Include(s, t.ParticipantType) {
		return fmt.Errorf("expected 'participant_type' to be one of %v", s)
	}
	return nil
}
