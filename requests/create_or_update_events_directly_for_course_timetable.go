package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateOrUpdateEventsDirectlyForCourseTimetable Creates and updates "timetable" events for a course or course section.
// Similar to {api:CalendarEventsApiController#set_course_timetable setting a course timetable},
// but instead of generating a list of events based on a timetable schedule,
// this endpoint expects a complete list of events.
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.CourseSectionID (Optional) Events will be created for the course section specified by course_section_id.
//    If not present, events will be created for the entire course.
// # Form.Events (Optional) An array of event objects to use.
// # Form.Events.StartAt (Optional) Start time for the event
// # Form.Events.EndAt (Optional) End time for the event
// # Form.Events.LocationName (Optional) Location name for the event
// # Form.Events.Code (Optional) A unique identifier that can be used to update the event at a later time
//    If one is not specified, an identifier will be generated based on the start and end times
// # Form.Events.Title (Optional) Title for the meeting. If not present, will default to the associated course's name
//
type CreateOrUpdateEventsDirectlyForCourseTimetable struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseSectionID string                                               `json:"course_section_id" url:"course_section_id,omitempty"` //  (Optional)
		Events          CreateOrUpdateEventsDirectlyForCourseTimetableEvents `json:"events" url:"events,omitempty"`                       //  (Optional)
	} `json:"form"`
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) GetMethod() string {
	return "POST"
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) GetURLPath() string {
	path := "courses/{course_id}/calendar_events/timetable_events"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}

type CreateOrUpdateEventsDirectlyForCourseTimetableEvents struct {
	StartAt      []time.Time `json:"start_at" url:"start_at,omitempty"`           //  (Optional)
	EndAt        []time.Time `json:"end_at" url:"end_at,omitempty"`               //  (Optional)
	LocationName []string    `json:"location_name" url:"location_name,omitempty"` //  (Optional)
	Code         []string    `json:"code" url:"code,omitempty"`                   //  (Optional)
	Title        []string    `json:"title" url:"title,omitempty"`                 //  (Optional)
}
