package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SetCourseTimetable Creates and updates "timetable" events for a course.
// Can automaticaly generate a series of calendar events based on simple schedules
// (e.g. "Monday and Wednesday at 2:00pm" )
//
// Existing timetable events for the course and course sections
// will be updated if they still are part of the timetable.
// Otherwise, they will be deleted.
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Timetables (Optional) An array of timetable objects for the course section specified by course_section_id.
//    If course_section_id is set to "all", events will be created for the entire course.
// # Timetables (Optional) A comma-separated list of abbreviated weekdays
//    (Mon-Monday, Tue-Tuesday, Wed-Wednesday, Thu-Thursday, Fri-Friday, Sat-Saturday, Sun-Sunday)
// # Timetables (Optional) Time to start each event at (e.g. "9:00 am")
// # Timetables (Optional) Time to end each event at (e.g. "9:00 am")
// # Timetables (Optional) A location name to set for each event
//
type SetCourseTimetable struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Timetables struct {
			CourseSectionID []string `json:"course_section_id" url:"course_section_id,omitempty"` //  (Optional)
		} `json:"timetables" url:"timetables,omitempty"`
	} `json:"form"`
}

func (t *SetCourseTimetable) GetMethod() string {
	return "POST"
}

func (t *SetCourseTimetable) GetURLPath() string {
	path := "courses/{course_id}/calendar_events/timetable"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *SetCourseTimetable) GetQuery() (string, error) {
	return "", nil
}

func (t *SetCourseTimetable) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SetCourseTimetable) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SetCourseTimetable) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetCourseTimetable) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
