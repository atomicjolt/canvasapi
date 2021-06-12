package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// PublishProvisionalGradesForAssignment Publish the selected provisional grade for all submissions to an assignment.
// Use the "Select provisional grade" endpoint to choose which provisional grade to publish
// for a particular submission.
//
// Students not in the moderation set will have their one and only provisional grade published.
//
// WARNING: This is irreversible. This will overwrite existing grades in the gradebook.
// https://canvas.instructure.com/doc/api/moderated_grading.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
type PublishProvisionalGradesForAssignment struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`
}

func (t *PublishProvisionalGradesForAssignment) GetMethod() string {
	return "POST"
}

func (t *PublishProvisionalGradesForAssignment) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/provisional_grades/publish"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *PublishProvisionalGradesForAssignment) GetQuery() (string, error) {
	return "", nil
}

func (t *PublishProvisionalGradesForAssignment) GetBody() (string, error) {
	return "", nil
}

func (t *PublishProvisionalGradesForAssignment) HasErrors() error {
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

func (t *PublishProvisionalGradesForAssignment) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
