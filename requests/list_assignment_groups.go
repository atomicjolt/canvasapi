package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListAssignmentGroups Returns the paginated list of assignment groups for the current context.
// The returned groups are sorted by their position field.
// https://canvas.instructure.com/doc/api/assignment_groups.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of assignments, discussion_topic, all_dates, assignment_visibility, overrides, submission, observed_users, can_edit, score_statisticsAssociations to include with the group. "discussion_topic", "all_dates", "can_edit",
//    "assignment_visibility" & "submission" are only valid if "assignments" is also included.
//    "score_statistics" requires that the "assignments" and "submission" options are included.
//    The "assignment_visibility" option additionally requires that the Differentiated Assignments course feature be turned on.
//    If "observed_users" is passed along with "assignments" and "submission", submissions for observed users will also be included as an array.
// # Query.AssignmentIDs (Optional) If "assignments" are included, optionally return only assignments having their ID in this array. This argument may also be passed as
//    a comma separated string.
// # Query.ExcludeAssignmentSubmissionTypes (Optional) . Must be one of online_quiz, discussion_topic, wiki_page, external_toolIf "assignments" are included, those with the specified submission types
//    will be excluded from the assignment groups.
// # Query.OverrideAssignmentDates (Optional) Apply assignment overrides for each assignment, defaults to true.
// # Query.GradingPeriodID (Optional) The id of the grading period in which assignment groups are being requested
//    (Requires grading periods to exist.)
// # Query.ScopeAssignmentsToStudent (Optional) If true, all assignments returned will apply to the current user in the
//    specified grading period. If assignments apply to other students in the
//    specified grading period, but not the current user, they will not be
//    returned. (Requires the grading_period_id argument and grading periods to
//    exist. In addition, the current user must be a student.)
//
type ListAssignmentGroups struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include                          []string `json:"include" url:"include,omitempty"`                                                         //  (Optional) . Must be one of assignments, discussion_topic, all_dates, assignment_visibility, overrides, submission, observed_users, can_edit, score_statistics
		AssignmentIDs                    []string `json:"assignment_ids" url:"assignment_ids,omitempty"`                                           //  (Optional)
		ExcludeAssignmentSubmissionTypes []string `json:"exclude_assignment_submission_types" url:"exclude_assignment_submission_types,omitempty"` //  (Optional) . Must be one of online_quiz, discussion_topic, wiki_page, external_tool
		OverrideAssignmentDates          bool     `json:"override_assignment_dates" url:"override_assignment_dates,omitempty"`                     //  (Optional)
		GradingPeriodID                  int64    `json:"grading_period_id" url:"grading_period_id,omitempty"`                                     //  (Optional)
		ScopeAssignmentsToStudent        bool     `json:"scope_assignments_to_student" url:"scope_assignments_to_student,omitempty"`               //  (Optional)
	} `json:"query"`
}

func (t *ListAssignmentGroups) GetMethod() string {
	return "GET"
}

func (t *ListAssignmentGroups) GetURLPath() string {
	path := "courses/{course_id}/assignment_groups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListAssignmentGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListAssignmentGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAssignmentGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAssignmentGroups) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"assignments", "discussion_topic", "all_dates", "assignment_visibility", "overrides", "submission", "observed_users", "can_edit", "score_statistics"}, v) {
			errs = append(errs, "Include must be one of assignments, discussion_topic, all_dates, assignment_visibility, overrides, submission, observed_users, can_edit, score_statistics")
		}
	}
	for _, v := range t.Query.ExcludeAssignmentSubmissionTypes {
		if v != "" && !string_utils.Include([]string{"online_quiz", "discussion_topic", "wiki_page", "external_tool"}, v) {
			errs = append(errs, "ExcludeAssignmentSubmissionTypes must be one of online_quiz, discussion_topic, wiki_page, external_tool")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAssignmentGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.AssignmentGroup, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.AssignmentGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
