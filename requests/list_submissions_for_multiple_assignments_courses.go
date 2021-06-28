package requests

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListSubmissionsForMultipleAssignmentsCourses A paginated list of all existing submissions for a given set of students and assignments.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # StudentIDs (Optional) List of student ids to return submissions for. If this argument is
//    omitted, return submissions for the calling user. Students may only list
//    their own submissions. Observers may only list those of associated
//    students. The special id "all" will return submissions for all students
//    in the course/section as appropriate.
// # AssignmentIDs (Optional) List of assignments to return submissions for. If none are given,
//    submissions for all assignments are returned.
// # Grouped (Optional) If this argument is present, the response will be grouped by student,
//    rather than a flat array of submissions.
// # PostToSIS (Optional) If this argument is set to true, the response will only include
//    submissions for assignments that have the post_to_sis flag set to true and
//    user enrollments that were added through sis.
// # SubmittedSince (Optional) If this argument is set, the response will only include submissions that
//    were submitted after the specified date_time. This will exclude
//    submissions that do not have a submitted_at which will exclude unsubmitted
//    submissions.
//    The value must be formatted as ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # GradedSince (Optional) If this argument is set, the response will only include submissions that
//    were graded after the specified date_time. This will exclude
//    submissions that have not been graded.
//    The value must be formatted as ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # GradingPeriodID (Optional) The id of the grading period in which submissions are being requested
//    (Requires grading periods to exist on the account)
// # WorkflowState (Optional) . Must be one of submitted, unsubmitted, graded, pending_reviewThe current status of the submission
// # EnrollmentState (Optional) . Must be one of active, concludedThe current state of the enrollments. If omitted will include all
//    enrollments that are not deleted.
// # StateBasedOnDate (Optional) If omitted it is set to true. When set to false it will ignore the effective
//    state of the student enrollments and use the workflow_state for the
//    enrollments. The argument is ignored unless enrollment_state argument is
//    also passed.
// # Order (Optional) . Must be one of id, graded_atThe order submissions will be returned in.  Defaults to "id".  Doesn't
//    affect results for "grouped" mode.
// # OrderDirection (Optional) . Must be one of ascending, descendingDetermines whether ordered results are returned in ascending or descending
//    order.  Defaults to "ascending".  Doesn't affect results for "grouped" mode.
// # Include (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, total_scores, visibility, course, userAssociations to include with the group. `total_scores` requires the
//    `grouped` argument.
//
type ListSubmissionsForMultipleAssignmentsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StudentIDs       []string  `json:"student_ids" url:"student_ids,omitempty"`                 //  (Optional)
		AssignmentIDs    []string  `json:"assignment_ids" url:"assignment_ids,omitempty"`           //  (Optional)
		Grouped          bool      `json:"grouped" url:"grouped,omitempty"`                         //  (Optional)
		PostToSIS        bool      `json:"post_to_sis" url:"post_to_sis,omitempty"`                 //  (Optional)
		SubmittedSince   time.Time `json:"submitted_since" url:"submitted_since,omitempty"`         //  (Optional)
		GradedSince      time.Time `json:"graded_since" url:"graded_since,omitempty"`               //  (Optional)
		GradingPeriodID  int64     `json:"grading_period_id" url:"grading_period_id,omitempty"`     //  (Optional)
		WorkflowState    string    `json:"workflow_state" url:"workflow_state,omitempty"`           //  (Optional) . Must be one of submitted, unsubmitted, graded, pending_review
		EnrollmentState  string    `json:"enrollment_state" url:"enrollment_state,omitempty"`       //  (Optional) . Must be one of active, concluded
		StateBasedOnDate bool      `json:"state_based_on_date" url:"state_based_on_date,omitempty"` //  (Optional)
		Order            string    `json:"order" url:"order,omitempty"`                             //  (Optional) . Must be one of id, graded_at
		OrderDirection   string    `json:"order_direction" url:"order_direction,omitempty"`         //  (Optional) . Must be one of ascending, descending
		Include          []string  `json:"include" url:"include,omitempty"`                         //  (Optional) . Must be one of submission_history, submission_comments, rubric_assessment, assignment, total_scores, visibility, course, user
	} `json:"query"`
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) GetMethod() string {
	return "GET"
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) GetURLPath() string {
	path := "courses/{course_id}/students/submissions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Query.WorkflowState != "" && !string_utils.Include([]string{"submitted", "unsubmitted", "graded", "pending_review"}, t.Query.WorkflowState) {
		errs = append(errs, "WorkflowState must be one of submitted, unsubmitted, graded, pending_review")
	}
	if t.Query.EnrollmentState != "" && !string_utils.Include([]string{"active", "concluded"}, t.Query.EnrollmentState) {
		errs = append(errs, "EnrollmentState must be one of active, concluded")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"id", "graded_at"}, t.Query.Order) {
		errs = append(errs, "Order must be one of id, graded_at")
	}
	if t.Query.OrderDirection != "" && !string_utils.Include([]string{"ascending", "descending"}, t.Query.OrderDirection) {
		errs = append(errs, "OrderDirection must be one of ascending, descending")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"submission_history", "submission_comments", "rubric_assessment", "assignment", "total_scores", "visibility", "course", "user"}, v) {
			errs = append(errs, "Include must be one of submission_history, submission_comments, rubric_assessment, assignment, total_scores, visibility, course, user")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListSubmissionsForMultipleAssignmentsCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
