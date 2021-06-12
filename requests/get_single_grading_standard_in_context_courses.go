package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleGradingStandardInContextCourses Returns a grading standard for the given context that is visible to the user.
// https://canvas.instructure.com/doc/api/grading_standards.html
//
// Path Parameters:
// # CourseID (Required) ID
// # GradingStandardID (Required) ID
//
type GetSingleGradingStandardInContextCourses struct {
	Path struct {
		CourseID          string `json:"course_id"`           //  (Required)
		GradingStandardID string `json:"grading_standard_id"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleGradingStandardInContextCourses) GetMethod() string {
	return "GET"
}

func (t *GetSingleGradingStandardInContextCourses) GetURLPath() string {
	path := "courses/{course_id}/grading_standards/{grading_standard_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{grading_standard_id}", fmt.Sprintf("%v", t.Path.GradingStandardID))
	return path
}

func (t *GetSingleGradingStandardInContextCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleGradingStandardInContextCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleGradingStandardInContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.GradingStandardID == "" {
		errs = append(errs, "'GradingStandardID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleGradingStandardInContextCourses) Do(c *canvasapi.Canvas) (*models.GradingStandard, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.GradingStandard{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
