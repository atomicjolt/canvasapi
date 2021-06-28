package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetCourseLevelAssignmentData Returns a list of assignments for the course sorted by due date. For
// each assignment returns basic assignment information, the grade breakdown,
// and a breakdown of on-time/late status of homework submissions.
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Async (Optional) If async is true, then the course_assignments call can happen asynch-
//    ronously and MAY return a response containing a progress_url key instead
//    of an assignments array. If it does, then it is the caller's
//    responsibility to poll the API again to see if the progress is complete.
//    If the data is ready (possibly even on the first async call) then it
//    will be passed back normally, as documented in the example response.
//
type GetCourseLevelAssignmentData struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Async bool `json:"async" url:"async,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetCourseLevelAssignmentData) GetMethod() string {
	return "GET"
}

func (t *GetCourseLevelAssignmentData) GetURLPath() string {
	path := "courses/{course_id}/analytics/assignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCourseLevelAssignmentData) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetCourseLevelAssignmentData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCourseLevelAssignmentData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCourseLevelAssignmentData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseLevelAssignmentData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
