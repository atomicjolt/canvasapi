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

// RetrieveAssignmentsEnabledForGradeExportToSISCourses Retrieve a list of published assignments flagged as "post_to_sis".
// See the Assignments API for more details on assignments.
// Assignment group and section information are included for convenience.
//
// Each section includes course information for the origin course and the
// cross-listed course, if applicable. The `origin_course` is the course to
// which the section belongs or the course from which the section was
// cross-listed. Generally, the `origin_course` should be preferred when
// performing integration work. The `xlist_course` is provided for consistency
// and is only present when the section has been cross-listed.
// See Sections API and Courses Api for me details.
//
// The `override` is only provided if the Differentiated Assignments course
// feature is turned on and the assignment has an override for that section.
// When there is an override for the assignment the override object's
// keys/values can be merged with the top level assignment object to create a
// view of the assignment object specific to that section.
// See Assignments api for more information on assignment overrides.
//
// restricts to courses that start before this date (if they have a start date)
// restricts to courses that end after this date (if they have an end date)
// information to include.
//
//   "student_overrides":: returns individual student override information
// https://canvas.instructure.com/doc/api/sis_integration.html
//
// Path Parameters:
// # Path.CourseID (Required) The ID of the course to query.
//
// Query Parameters:
// # Query.AccountID (Optional) The ID of the account to query.
// # Query.StartsBefore (Optional) When searching on an account,
// # Query.EndsAfter (Optional) When searching on an account,
// # Query.Include (Optional) . Must be one of student_overridesArray of additional
//
type RetrieveAssignmentsEnabledForGradeExportToSISCourses struct {
	Path struct {
		CourseID int64 `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		AccountID    int64     `json:"account_id" url:"account_id,omitempty"`       //  (Optional)
		StartsBefore time.Time `json:"starts_before" url:"starts_before,omitempty"` //  (Optional)
		EndsAfter    time.Time `json:"ends_after" url:"ends_after,omitempty"`       //  (Optional)
		Include      string    `json:"include" url:"include,omitempty"`             //  (Optional) . Must be one of student_overrides
	} `json:"query"`
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) GetMethod() string {
	return "GET"
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) GetURLPath() string {
	path := "/sis/courses/{course_id}/assignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) HasErrors() error {
	errs := []string{}
	if t.Query.Include != "" && !string_utils.Include([]string{"student_overrides"}, t.Query.Include) {
		errs = append(errs, "Include must be one of student_overrides")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RetrieveAssignmentsEnabledForGradeExportToSISCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
