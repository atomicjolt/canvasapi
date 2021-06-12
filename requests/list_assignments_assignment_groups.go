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

// ListAssignmentsAssignmentGroups Returns the paginated list of assignments for the current course or assignment group.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentGroupID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of submission, assignment_visibility, all_dates, overrides, observed_users, can_edit, score_statisticsOptional information to include with each assignment:
//    submission:: The current user's current +Submission+
//    assignment_visibility:: An array of ids of students who can see the assignment
//    all_dates:: An array of +AssignmentDate+ structures, one for each override, and also a +base+ if the assignment has an "Everyone" / "Everyone Else" date
//    overrides:: An array of +AssignmentOverride+ structures
//    observed_users:: An array of submissions for observed users
//    can_edit:: an extra Boolean value will be included with each +Assignment+ (and +AssignmentDate+ if +all_dates+ is supplied) to indicate whether the caller can edit the assignment or date. Moderated grading and closed grading periods may restrict a user's ability to edit an assignment.
//    score_statistics:: An object containing min, max, and mean score on this assignment. This will not be included for students if there are less than 5 graded assignments or if disabled by the instructor. Only valid if 'submission' is also included.
// # SearchTerm (Optional) The partial title of the assignments to match and return.
// # OverrideAssignmentDates (Optional) Apply assignment overrides for each assignment, defaults to true.
// # NeedsGradingCountBySection (Optional) Split up "needs_grading_count" by sections into the "needs_grading_count_by_section" key, defaults to false
// # Bucket (Optional) . Must be one of past, overdue, undated, ungraded, unsubmitted, upcoming, futureIf included, only return certain assignments depending on due date and submission status.
// # AssignmentIDs (Optional) if set, return only assignments specified
// # OrderBy (Optional) . Must be one of position, name, due_atDetermines the order of the assignments. Defaults to "position".
// # PostToSIS (Optional) Return only assignments that have post_to_sis set or not set.
//
type ListAssignmentsAssignmentGroups struct {
	Path struct {
		CourseID          string `json:"course_id"`           //  (Required)
		AssignmentGroupID string `json:"assignment_group_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include                    []string `json:"include"`                        //  (Optional) . Must be one of submission, assignment_visibility, all_dates, overrides, observed_users, can_edit, score_statistics
		SearchTerm                 string   `json:"search_term"`                    //  (Optional)
		OverrideAssignmentDates    bool     `json:"override_assignment_dates"`      //  (Optional)
		NeedsGradingCountBySection bool     `json:"needs_grading_count_by_section"` //  (Optional)
		Bucket                     string   `json:"bucket"`                         //  (Optional) . Must be one of past, overdue, undated, ungraded, unsubmitted, upcoming, future
		AssignmentIDs              []string `json:"assignment_ids"`                 //  (Optional)
		OrderBy                    string   `json:"order_by"`                       //  (Optional) . Must be one of position, name, due_at
		PostToSIS                  bool     `json:"post_to_sis"`                    //  (Optional)
	} `json:"query"`
}

func (t *ListAssignmentsAssignmentGroups) GetMethod() string {
	return "GET"
}

func (t *ListAssignmentsAssignmentGroups) GetURLPath() string {
	path := "courses/{course_id}/assignment_groups/{assignment_group_id}/assignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_group_id}", fmt.Sprintf("%v", t.Path.AssignmentGroupID))
	return path
}

func (t *ListAssignmentsAssignmentGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAssignmentsAssignmentGroups) GetBody() (string, error) {
	return "", nil
}

func (t *ListAssignmentsAssignmentGroups) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentGroupID == "" {
		errs = append(errs, "'AssignmentGroupID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"submission", "assignment_visibility", "all_dates", "overrides", "observed_users", "can_edit", "score_statistics"}, v) {
			errs = append(errs, "Include must be one of submission, assignment_visibility, all_dates, overrides, observed_users, can_edit, score_statistics")
		}
	}
	if !string_utils.Include([]string{"past", "overdue", "undated", "ungraded", "unsubmitted", "upcoming", "future"}, t.Query.Bucket) {
		errs = append(errs, "Bucket must be one of past, overdue, undated, ungraded, unsubmitted, upcoming, future")
	}
	if !string_utils.Include([]string{"position", "name", "due_at"}, t.Query.OrderBy) {
		errs = append(errs, "OrderBy must be one of position, name, due_at")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAssignmentsAssignmentGroups) Do(c *canvasapi.Canvas) ([]*models.Assignment, error) {
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
