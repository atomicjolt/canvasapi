package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
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
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_statusAssociations to include with the group.  "group" will add group_id and group_name.
// # Query.Grouped (Optional) If this argument is true, the response will be grouped by student groups.
//
type ListAssignmentSubmissionsCourses struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_status
		Grouped bool     `json:"grouped" url:"grouped,omitempty"` //  (Optional)
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
	return v.Encode(), nil
}

func (t *ListAssignmentSubmissionsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAssignmentSubmissionsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAssignmentSubmissionsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission_history", "submission_comments", "rubric_assessment", "assignment", "visibility", "course", "user", "group", "read_status"}, v) {
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
