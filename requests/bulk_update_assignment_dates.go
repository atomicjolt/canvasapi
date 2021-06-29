package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// BulkUpdateAssignmentDates Update due dates and availability dates for multiple assignments in a course.
//
// Accepts a JSON array of objects containing two keys each: +id+, the assignment id,
// and +all_dates+, an array of +AssignmentDate+ structures containing the base and/or override
// dates for the assignment, as returned from the {api:AssignmentsApiController#index List assignments}
// endpoint with +include[]=all_dates+.
//
// This endpoint cannot create or destroy assignment overrides; any existing assignment overrides
// that are not referenced in the arguments will be left alone. If an override is given, any dates
// that are not supplied with it will be defaulted. To clear a date, specify null explicitly.
//
// All referenced assignments will be validated before any are saved. A list of errors will
// be returned if any provided dates are invalid, and no changes will be saved.
//
// The bulk update is performed in a background job, use the {api:ProgressController#show Progress API}
// to check its status.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type BulkUpdateAssignmentDates struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *BulkUpdateAssignmentDates) GetMethod() string {
	return "PUT"
}

func (t *BulkUpdateAssignmentDates) GetURLPath() string {
	path := "courses/{course_id}/assignments/bulk_update"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *BulkUpdateAssignmentDates) GetQuery() (string, error) {
	return "", nil
}

func (t *BulkUpdateAssignmentDates) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *BulkUpdateAssignmentDates) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *BulkUpdateAssignmentDates) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BulkUpdateAssignmentDates) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
