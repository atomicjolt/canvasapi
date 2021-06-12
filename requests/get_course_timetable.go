package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCourseTimetable Returns the last timetable set by the
// {api:CalendarEventsApiController#set_course_timetable Set a course timetable} endpoint
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetCourseTimetable struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *GetCourseTimetable) GetMethod() string {
	return "GET"
}

func (t *GetCourseTimetable) GetURLPath() string {
	path := "courses/{course_id}/calendar_events/timetable"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCourseTimetable) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCourseTimetable) GetBody() (string, error) {
	return "", nil
}

func (t *GetCourseTimetable) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseTimetable) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
