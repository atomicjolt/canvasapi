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

// GetAssignmentGroup Returns the assignment group with the given id.
// https://canvas.instructure.com/doc/api/assignment_groups.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentGroupID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of assignments, discussion_topic, assignment_visibility, submission, score_statisticsAssociations to include with the group. "discussion_topic" and "assignment_visibility" and "submission"
//    are only valid if "assignments" is also included. "score_statistics" is only valid if "submission" and
//    "assignments" are also included. The "assignment_visibility" option additionally requires that the Differentiated Assignments
//    course feature be turned on.
// # Query.OverrideAssignmentDates (Optional) Apply assignment overrides for each assignment, defaults to true.
// # Query.GradingPeriodID (Optional) The id of the grading period in which assignment groups are being requested
//    (Requires grading periods to exist on the account)
//
type GetAssignmentGroup struct {
	Path struct {
		CourseID          string `json:"course_id" url:"course_id,omitempty"`                     //  (Required)
		AssignmentGroupID string `json:"assignment_group_id" url:"assignment_group_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include                 []string `json:"include" url:"include,omitempty"`                                     //  (Optional) . Must be one of assignments, discussion_topic, assignment_visibility, submission, score_statistics
		OverrideAssignmentDates bool     `json:"override_assignment_dates" url:"override_assignment_dates,omitempty"` //  (Optional)
		GradingPeriodID         int64    `json:"grading_period_id" url:"grading_period_id,omitempty"`                 //  (Optional)
	} `json:"query"`
}

func (t *GetAssignmentGroup) GetMethod() string {
	return "GET"
}

func (t *GetAssignmentGroup) GetURLPath() string {
	path := "courses/{course_id}/assignment_groups/{assignment_group_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_group_id}", fmt.Sprintf("%v", t.Path.AssignmentGroupID))
	return path
}

func (t *GetAssignmentGroup) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAssignmentGroup) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAssignmentGroup) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAssignmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentGroupID == "" {
		errs = append(errs, "'Path.AssignmentGroupID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"assignments", "discussion_topic", "assignment_visibility", "submission", "score_statistics"}, v) {
			errs = append(errs, "Include must be one of assignments, discussion_topic, assignment_visibility, submission, score_statistics")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAssignmentGroup) Do(c *canvasapi.Canvas) (*models.AssignmentGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AssignmentGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
