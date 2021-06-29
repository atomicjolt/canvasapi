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
)

// GetAlignedAssignmentsForOutcomeInCourseForParticularStudent
// https://canvas.instructure.com/doc/api/outcomes.html
//
// Path Parameters:
// # CourseID (Required) The id of the course
//
// Query Parameters:
// # StudentID (Optional) The id of the student
//
type GetAlignedAssignmentsForOutcomeInCourseForParticularStudent struct {
	Path struct {
		CourseID int64 `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StudentID int64 `json:"student_id" url:"student_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) GetMethod() string {
	return "GET"
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) GetURLPath() string {
	path := "courses/{course_id}/outcome_alignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) HasErrors() error {
	return nil
}

func (t *GetAlignedAssignmentsForOutcomeInCourseForParticularStudent) Do(c *canvasapi.Canvas) ([]*models.OutcomeAlignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.OutcomeAlignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
