package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetBlueprintInformation Using 'default' as the template_id should suffice for the current implmentation (as there should be only one template per course).
// However, using specific template ids may become necessary in the future
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TemplateID (Required) ID
//
type GetBlueprintInformation struct {
	Path struct {
		CourseID   string `json:"course_id"`   //  (Required)
		TemplateID string `json:"template_id"` //  (Required)
	} `json:"path"`
}

func (t *GetBlueprintInformation) GetMethod() string {
	return "GET"
}

func (t *GetBlueprintInformation) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	return path
}

func (t *GetBlueprintInformation) GetQuery() (string, error) {
	return "", nil
}

func (t *GetBlueprintInformation) GetBody() (string, error) {
	return "", nil
}

func (t *GetBlueprintInformation) HasErrors() error {
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

func (t *GetBlueprintInformation) Do(c *canvasapi.Canvas) (*models.BlueprintTemplate, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.BlueprintTemplate{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
