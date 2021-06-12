package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetUserInACourseLevelMessagingData Returns messaging "hits" grouped by day through the entire history of the
// course. Returns a hash containing the number of instructor-to-student messages,
// and student-to-instructor messages, where the hash keys are dates
// in the format "YYYY-MM-DD". Message hits include Conversation messages and
// comments on homework submissions.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # StudentID (Required) ID
//
type GetUserInACourseLevelMessagingData struct {
	Path struct {
		CourseID  string `json:"course_id"`  //  (Required)
		StudentID string `json:"student_id"` //  (Required)
	} `json:"path"`
}

func (t *GetUserInACourseLevelMessagingData) GetMethod() string {
	return "GET"
}

func (t *GetUserInACourseLevelMessagingData) GetURLPath() string {
	path := "courses/{course_id}/analytics/users/{student_id}/communication"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{student_id}", fmt.Sprintf("%v", t.Path.StudentID))
	return path
}

func (t *GetUserInACourseLevelMessagingData) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUserInACourseLevelMessagingData) GetBody() (string, error) {
	return "", nil
}

func (t *GetUserInACourseLevelMessagingData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.StudentID == "" {
		errs = append(errs, "'StudentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUserInACourseLevelMessagingData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
