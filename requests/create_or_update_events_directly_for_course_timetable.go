package requests

import (
	"fmt"
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
// # CourseID (Required) ID
//
// Form Parameters:
// # CourseSectionID (Optional) Events will be created for the course section specified by course_section_id.
//    If not present, events will be created for the entire course.
// # Events (Optional) An array of event objects to use.
// # Events (Optional) Start time for the event
// # Events (Optional) End time for the event
// # Events (Optional) Location name for the event
// # Events (Optional) A unique identifier that can be used to update the event at a later time
//    If one is not specified, an identifier will be generated based on the start and end times
// # Events (Optional) Title for the meeting. If not present, will default to the associated course's name
//
type CreateOrUpdateEventsDirectlyForCourseTimetable struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseSectionID string                                               `json:"course_section_id"` //  (Optional)
		Events          CreateOrUpdateEventsDirectlyForCourseTimetableEvents `json:"events"`            //  (Optional)
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

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateOrUpdateEventsDirectlyForCourseTimetable) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
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
	StartAt      []time.Time `json:"start_at"`      //  (Optional)
	EndAt        []time.Time `json:"end_at"`        //  (Optional)
	LocationName []string    `json:"location_name"` //  (Optional)
	Code         []string    `json:"code"`          //  (Optional)
	Title        []string    `json:"title"`         //  (Optional)
}
