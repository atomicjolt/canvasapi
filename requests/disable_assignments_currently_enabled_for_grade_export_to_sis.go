package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DisableAssignmentsCurrentlyEnabledForGradeExportToSIS Disable all assignments flagged as "post_to_sis", with the option of making it
// specific to a grading period, in a course.
//
// On success, the response will be 204 No Content with an empty body.
//
// On failure, the response will be 400 Bad Request with a body of a specific
// message.
//
// For disabling assignments in a specific grading period
// https://canvas.instructure.com/doc/api/sis_integration.html
//
// Path Parameters:
// # Path.CourseID (Required) The ID of the course.
//
// Form Parameters:
// # Form.GradingPeriodID (Optional) The ID of the grading period.
//
type DisableAssignmentsCurrentlyEnabledForGradeExportToSIS struct {
	Path struct {
		CourseID int64 `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		GradingPeriodID int64 `json:"grading_period_id" url:"grading_period_id,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) GetMethod() string {
	return "PUT"
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) GetURLPath() string {
	path := "/sis/courses/{course_id}/disable_post_to_sis"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) GetQuery() (string, error) {
	return "", nil
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) HasErrors() error {
	return nil
}

func (t *DisableAssignmentsCurrentlyEnabledForGradeExportToSIS) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
