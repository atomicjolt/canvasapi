package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListAssignmentSubmissionsCourses A paginated list of all existing submissions for an assignment.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_statusAssociations to include with the group.  "group" will add group_id and group_name.
// # Grouped (Optional) If this argument is true, the response will be grouped by student groups.
//
type ListAssignmentSubmissionsCourses struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_status
		Grouped bool     `json:"grouped"` //  (Optional)
	} `json:"query"`
}

func (t *ListAssignmentSubmissionsCourses) GetMethod() string {
	return "GET"
}

func (t *ListAssignmentSubmissionsCourses) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/submissions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ListAssignmentSubmissionsCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAssignmentSubmissionsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListAssignmentSubmissionsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"submission_history", "submission_comments", "rubric_assessment", "assignment", "visibility", "course", "user", "group", "read_status"}, v) {
			errs = append(errs, "Include must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_status")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAssignmentSubmissionsCourses) Do(c *canvasapi.Canvas) ([]*models.Submission, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Submission{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}