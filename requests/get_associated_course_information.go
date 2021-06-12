package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetAssociatedCourseInformation Returns a list of courses that are configured to receive updates from this blueprint
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TemplateID (Required) ID
//
type GetAssociatedCourseInformation struct {
	Path struct {
		CourseID   string `json:"course_id"`   //  (Required)
		TemplateID string `json:"template_id"` //  (Required)
	} `json:"path"`
}

func (t *GetAssociatedCourseInformation) GetMethod() string {
	return "GET"
}

func (t *GetAssociatedCourseInformation) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/associated_courses"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	return path
}

func (t *GetAssociatedCourseInformation) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAssociatedCourseInformation) GetBody() (string, error) {
	return "", nil
}

func (t *GetAssociatedCourseInformation) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'TemplateID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAssociatedCourseInformation) Do(c *canvasapi.Canvas) ([]*models.Course, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Course{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
