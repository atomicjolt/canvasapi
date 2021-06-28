package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetEffectiveDueDates For each assignment in the course, returns each assigned student's ID
// and their corresponding due date along with some grading period data.
// Returns a collection with keys representing assignment IDs and values as a
// collection containing keys representing student IDs and values representing
// the student's effective due_at, the grading_period_id of which the due_at falls
// in, and whether or not the grading period is closed (in_closed_grading_period)
//
// The list of assignment IDs for which effective student due dates are
// requested. If not provided, all assignments in the course will be used.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # AssignmentIDs (Optional) no description
//
type GetEffectiveDueDates struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		AssignmentIDs []string `json:"assignment_ids" url:"assignment_ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetEffectiveDueDates) GetMethod() string {
	return "GET"
}

func (t *GetEffectiveDueDates) GetURLPath() string {
	path := "courses/{course_id}/effective_due_dates"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetEffectiveDueDates) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetEffectiveDueDates) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetEffectiveDueDates) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetEffectiveDueDates) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetEffectiveDueDates) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
