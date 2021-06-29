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

// ListAssignmentsAssignments Returns the paginated list of assignments for the current course or assignment group.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of submission, assignment_visibility, all_dates, overrides, observed_users, can_edit, score_statisticsOptional information to include with each assignment:
//    submission:: The current user's current +Submission+
//    assignment_visibility:: An array of ids of students who can see the assignment
//    all_dates:: An array of +AssignmentDate+ structures, one for each override, and also a +base+ if the assignment has an "Everyone" / "Everyone Else" date
//    overrides:: An array of +AssignmentOverride+ structures
//    observed_users:: An array of submissions for observed users
//    can_edit:: an extra Boolean value will be included with each +Assignment+ (and +AssignmentDate+ if +all_dates+ is supplied) to indicate whether the caller can edit the assignment or date. Moderated grading and closed grading periods may restrict a user's ability to edit an assignment.
//    score_statistics:: An object containing min, max, and mean score on this assignment. This will not be included for students if there are less than 5 graded assignments or if disabled by the instructor. Only valid if 'submission' is also included.
// # Query.SearchTerm (Optional) The partial title of the assignments to match and return.
// # Query.OverrideAssignmentDates (Optional) Apply assignment overrides for each assignment, defaults to true.
// # Query.NeedsGradingCountBySection (Optional) Split up "needs_grading_count" by sections into the "needs_grading_count_by_section" key, defaults to false
// # Query.Bucket (Optional) . Must be one of past, overdue, undated, ungraded, unsubmitted, upcoming, futureIf included, only return certain assignments depending on due date and submission status.
// # Query.AssignmentIDs (Optional) if set, return only assignments specified
// # Query.OrderBy (Optional) . Must be one of position, name, due_atDetermines the order of the assignments. Defaults to "position".
// # Query.PostToSIS (Optional) Return only assignments that have post_to_sis set or not set.
//
type ListAssignmentsAssignments struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include                    []string `json:"include" url:"include,omitempty"`                                               //  (Optional) . Must be one of submission, assignment_visibility, all_dates, overrides, observed_users, can_edit, score_statistics
		SearchTerm                 string   `json:"search_term" url:"search_term,omitempty"`                                       //  (Optional)
		OverrideAssignmentDates    bool     `json:"override_assignment_dates" url:"override_assignment_dates,omitempty"`           //  (Optional)
		NeedsGradingCountBySection bool     `json:"needs_grading_count_by_section" url:"needs_grading_count_by_section,omitempty"` //  (Optional)
		Bucket                     string   `json:"bucket" url:"bucket,omitempty"`                                                 //  (Optional) . Must be one of past, overdue, undated, ungraded, unsubmitted, upcoming, future
		AssignmentIDs              []string `json:"assignment_ids" url:"assignment_ids,omitempty"`                                 //  (Optional)
		OrderBy                    string   `json:"order_by" url:"order_by,omitempty"`                                             //  (Optional) . Must be one of position, name, due_at
		PostToSIS                  bool     `json:"post_to_sis" url:"post_to_sis,omitempty"`                                       //  (Optional)
	} `json:"query"`
}

func (t *ListAssignmentsAssignments) GetMethod() string {
	return "GET"
}

func (t *ListAssignmentsAssignments) GetURLPath() string {
	path := "courses/{course_id}/assignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListAssignmentsAssignments) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListAssignmentsAssignments) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAssignmentsAssignments) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAssignmentsAssignments) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission", "assignment_visibility", "all_dates", "overrides", "observed_users", "can_edit", "score_statistics"}, v) {
			errs = append(errs, "Include must be one of submission, assignment_visibility, all_dates, overrides, observed_users, can_edit, score_statistics")
		}
	}
	if t.Query.Bucket != "" && !string_utils.Include([]string{"past", "overdue", "undated", "ungraded", "unsubmitted", "upcoming", "future"}, t.Query.Bucket) {
		errs = append(errs, "Bucket must be one of past, overdue, undated, ungraded, unsubmitted, upcoming, future")
	}
	if t.Query.OrderBy != "" && !string_utils.Include([]string{"position", "name", "due_at"}, t.Query.OrderBy) {
		errs = append(errs, "OrderBy must be one of position, name, due_at")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAssignmentsAssignments) Do(c *canvasapi.Canvas) ([]*models.Assignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Assignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
