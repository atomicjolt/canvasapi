package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetUserInACourseLevelParticipationData Returns page view hits grouped by hour, and participation details through the
// entire history of the course.
//
// `page_views` are returned as a hash, where the keys are iso8601 dates, bucketed by the hour.
// `participations` are returned as an array of hashes, sorted oldest to newest.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.StudentID (Required) ID
//
type GetUserInACourseLevelParticipationData struct {
	Path struct {
		CourseID  string `json:"course_id" url:"course_id,omitempty"`   //  (Required)
		StudentID string `json:"student_id" url:"student_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetUserInACourseLevelParticipationData) GetMethod() string {
	return "GET"
}

func (t *GetUserInACourseLevelParticipationData) GetURLPath() string {
	path := "courses/{course_id}/analytics/users/{student_id}/activity"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{student_id}", fmt.Sprintf("%v", t.Path.StudentID))
	return path
}

func (t *GetUserInACourseLevelParticipationData) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUserInACourseLevelParticipationData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetUserInACourseLevelParticipationData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetUserInACourseLevelParticipationData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.StudentID == "" {
		errs = append(errs, "'Path.StudentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUserInACourseLevelParticipationData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
