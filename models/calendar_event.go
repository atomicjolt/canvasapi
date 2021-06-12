package models

import (
	"time"
)

type CalendarEvent struct {
	ID                         int64     `json:"id"`                           // The ID of the calendar event.Example: 234
	Title                      string    `json:"title"`                        // The title of the calendar event.Example: Paintball Fight!
	StartAt                    time.Time `json:"start_at"`                     // The start timestamp of the event.Example: 2012-07-19T15:00:00-06:00
	EndAt                      time.Time `json:"end_at"`                       // The end timestamp of the event.Example: 2012-07-19T16:00:00-06:00
	Description                string    `json:"description"`                  // The HTML description of the event.Example: <b>It's that time again!</b>
	LocationName               string    `json:"location_name"`                // The location name of the event.Example: Greendale Community College
	LocationAddress            string    `json:"location_address"`             // The address where the event is taking place.Example: Greendale, Colorado
	ContextCode                string    `json:"context_code"`                 // the context code of the calendar this event belongs to (course, user or group).Example: course_123
	EffectiveContextCode       string    `json:"effective_context_code"`       // if specified, it indicates which calendar this event should be displayed on. for example, a section-level event would have the course's context code here, while the section's context code would be returned above).
	ContextName                string    `json:"context_name"`                 // the context name of the calendar this event belongs to (course, user or group).Example: Chemistry 101
	AllContextCodes            string    `json:"all_context_codes"`            // a comma-separated list of all calendar contexts this event is part of.Example: course_123,course_456
	WorkflowState              string    `json:"workflow_state"`               // Current state of the event ('active', 'locked' or 'deleted') 'locked' indicates that start_at/end_at cannot be changed (though the event could be deleted). Normally only reservations or time slots with reservations are locked (see the Appointment Groups API).Example: active
	Hidden                     bool      `json:"hidden"`                       // Whether this event should be displayed on the calendar. Only true for course-level events with section-level child events..
	ParentEventID              int64     `json:"parent_event_id"`              // Normally null. If this is a reservation (see the Appointment Groups API), the id will indicate the time slot it is for. If this is a section-level event, this will be the course-level parent event..
	ChildEventsCount           int64     `json:"child_events_count"`           // The number of child_events. See child_events (and parent_event_id).Example: 0
	ChildEvents                []int64   `json:"child_events"`                 // Included by default, but may be excluded (see include[] option). If this is a time slot (see the Appointment Groups API) this will be a list of any reservations. If this is a course-level event, this will be a list of section-level events (if any).
	Url                        string    `json:"url"`                          // URL for this calendar event (to update, delete, etc.).Example: https://example.com/api/v1/calendar_events/234
	HtmlUrl                    string    `json:"html_url"`                     // URL for a user to view this event.Example: https://example.com/calendar?event_id=234&include_contexts=course_123
	AllDayDate                 time.Time `json:"all_day_date"`                 // The date of this event.Example: 2012-07-19
	AllDay                     bool      `json:"all_day"`                      // Boolean indicating whether this is an all-day event (midnight to midnight).
	CreatedAt                  time.Time `json:"created_at"`                   // When the calendar event was created.Example: 2012-07-12T10:55:20-06:00
	UpdatedAt                  time.Time `json:"updated_at"`                   // When the calendar event was last updated.Example: 2012-07-12T10:55:20-06:00
	AppointmentGroupID         int64     `json:"appointment_group_id"`         // Various Appointment-Group-related fields.These fields are only pertinent to time slots (appointments) and reservations of those time slots. See the Appointment Groups API. The id of the appointment group.
	AppointmentGroupUrl        string    `json:"appointment_group_url"`        // The API URL of the appointment group.
	OwnReservation             bool      `json:"own_reservation"`              // If the event is a reservation, this a boolean indicating whether it is the current user's reservation, or someone else's.
	ReserveUrl                 string    `json:"reserve_url"`                  // If the event is a time slot, the API URL for reserving it.
	Reserved                   bool      `json:"reserved"`                     // If the event is a time slot, a boolean indicating whether the user has already made a reservation for it.
	ParticipantType            string    `json:"participant_type"`             // The type of participant to sign up for a slot: 'User' or 'Group'.Example: User
	ParticipantsPerAppointment int64     `json:"participants_per_appointment"` // If the event is a time slot, this is the participant limit.
	AvailableSlots             int64     `json:"available_slots"`              // If the event is a time slot and it has a participant limit, an integer indicating how many slots are available.
	User                       string    `json:"user"`                         // If the event is a user-level reservation, this will contain the user participant JSON (refer to the Users API)..
	Group                      string    `json:"group"`                        // If the event is a group-level reservation, this will contain the group participant JSON (refer to the Groups API)..
}

func (t *CalendarEvent) HasError() error {
	return nil
}
