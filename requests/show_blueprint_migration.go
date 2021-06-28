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

// ShowBlueprintMigration Shows the status of a migration. This endpoint can be called on a blueprint course. See also
// {api:MasterCourses::MasterTemplatesController#imports_show the associated course side}.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TemplateID (Required) ID
// # ID (Required) ID
//
type ShowBlueprintMigration struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		TemplateID string `json:"template_id" url:"template_id,omitempty"` //  (Required)
		ID         string `json:"id" url:"id,omitempty"`                   //  (Required)
	} `json:"path"`
}

func (t *ShowBlueprintMigration) GetMethod() string {
	return "GET"
}

func (t *ShowBlueprintMigration) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/migrations/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowBlueprintMigration) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowBlueprintMigration) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowBlueprintMigration) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowBlueprintMigration) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'TemplateID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowBlueprintMigration) Do(c *canvasapi.Canvas) (*models.BlueprintMigration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.BlueprintMigration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
