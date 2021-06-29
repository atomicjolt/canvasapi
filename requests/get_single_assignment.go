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

// GetSingleAssignment Returns the assignment with the given id.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of submission, assignment_visibility, overrides, observed_users, can_edit, score_statisticsAssociations to include with the assignment. The "assignment_visibility" option
//    requires that the Differentiated Assignments course feature be turned on. If
//    "observed_users" is passed, submissions for observed users will also be included.
//    For "score_statistics" to be included, the "submission" option must also be set.
// # Query.OverrideAssignmentDates (Optional) Apply assignment overrides to the assignment, defaults to true.
// # Query.NeedsGradingCountBySection (Optional) Split up "needs_grading_count" by sections into the "needs_grading_count_by_section" key, defaults to false
// # Query.AllDates (Optional) All dates associated with the assignment, if applicable
//
type GetSingleAssignment struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Include                    []string `json:"include" url:"include,omitempty"`                                               //  (Optional) . Must be one of submission, assignment_visibility, overrides, observed_users, can_edit, score_statistics
		OverrideAssignmentDates    bool     `json:"override_assignment_dates" url:"override_assignment_dates,omitempty"`           //  (Optional)
		NeedsGradingCountBySection bool     `json:"needs_grading_count_by_section" url:"needs_grading_count_by_section,omitempty"` //  (Optional)
		AllDates                   bool     `json:"all_dates" url:"all_dates,omitempty"`                                           //  (Optional)
	} `json:"query"`
}

func (t *GetSingleAssignment) GetMethod() string {
	return "GET"
}

func (t *GetSingleAssignment) GetURLPath() string {
	path := "courses/{course_id}/assignments/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleAssignment) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetSingleAssignment) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleAssignment) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleAssignment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission", "assignment_visibility", "overrides", "observed_users", "can_edit", "score_statistics"}, v) {
			errs = append(errs, "Include must be one of submission, assignment_visibility, overrides, observed_users, can_edit, score_statistics")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleAssignment) Do(c *canvasapi.Canvas) (*models.Assignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Assignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
