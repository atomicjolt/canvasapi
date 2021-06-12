package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ShowProvisionalGradeStatusForStudent Tell whether the student's submission needs one or more provisional grades.
// https://canvas.instructure.com/doc/api/moderated_grading.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
// Query Parameters:
// # StudentID (Optional) The id of the student to show the status for
//
type ShowProvisionalGradeStatusForStudent struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`

	Query struct {
		StudentID int64 `json:"student_id"` //  (Optional)
	} `json:"query"`
}

func (t *ShowProvisionalGradeStatusForStudent) GetMethod() string {
	return "GET"
}

func (t *ShowProvisionalGradeStatusForStudent) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/provisional_grades/status"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ShowProvisionalGradeStatusForStudent) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ShowProvisionalGradeStatusForStudent) GetBody() (string, error) {
	return "", nil
}

func (t *ShowProvisionalGradeStatusForStudent) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowProvisionalGradeStatusForStudent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
