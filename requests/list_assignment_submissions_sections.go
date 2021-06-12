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

// ListAssignmentSubmissionsSections A paginated list of all existing submissions for an assignment.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # SectionID (Required) ID
// # AssignmentID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_statusAssociations to include with the group.  "group" will add group_id and group_name.
// # Grouped (Optional) If this argument is true, the response will be grouped by student groups.
//
type ListAssignmentSubmissionsSections struct {
	Path struct {
		SectionID    string `json:"section_id"`    //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, visibility, course, user, group, read_status
		Grouped bool     `json:"grouped"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAssignmentSubmissionsSections) GetBody() (string, error) {
	return "", nil
}

func (t *ListAssignmentSubmissionsSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'SectionID' is required")
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
