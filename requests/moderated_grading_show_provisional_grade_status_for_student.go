package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ModeratedGradingShowProvisionalGradeStatusForStudent Determine whether or not the student's submission needs one or more provisional grades.
// https://canvas.instructure.com/doc/api/moderated_grading.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
// Query Parameters:
// # AnonymousID (Optional) The id of the student to show the status for
//
type ModeratedGradingShowProvisionalGradeStatusForStudent struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`

	Query struct {
		AnonymousID string `json:"anonymous_id"` //  (Optional)
	} `json:"query"`
}

func (t *ModeratedGradingShowProvisionalGradeStatusForStudent) GetMethod() string {
	return "GET"
}

func (t *ModeratedGradingShowProvisionalGradeStatusForStudent) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/anonymous_provisional_grades/status"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ModeratedGradingShowProvisionalGradeStatusForStudent) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ModeratedGradingShowProvisionalGradeStatusForStudent) GetBody() (string, error) {
	return "", nil
}

func (t *ModeratedGradingShowProvisionalGradeStatusForStudent) HasErrors() error {
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

func (t *ModeratedGradingShowProvisionalGradeStatusForStudent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
