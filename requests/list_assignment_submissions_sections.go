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

// ListAssignmentSubmissionsSections A paginated list of all existing submissions for an assignment.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # Path.SectionID (Required) ID
// # Path.AssignmentID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_statusAssociations to include with the group.  "group" will add group_id and group_name.
// # Query.Grouped (Optional) If this argument is true, the response will be grouped by student groups.
//
type ListAssignmentSubmissionsSections struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_status
		Grouped bool     `json:"grouped" url:"grouped,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListAssignmentSubmissionsSections) GetMethod() string {
	return "GET"
}

func (t *ListAssignmentSubmissionsSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *ListAssignmentSubmissionsSections) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListAssignmentSubmissionsSections) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAssignmentSubmissionsSections) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAssignmentSubmissionsSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'Path.SectionID' is required")
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

func (t *ListAssignmentSubmissionsSections) Do(c *canvasapi.Canvas) ([]*models.Submission, error) {
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
