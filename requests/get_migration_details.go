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

// GetMigrationDetails Show the changes that were propagated in a blueprint migration. This endpoint can be called on a
// blueprint course. See also {api:MasterCourses::MasterTemplatesController#import_details the associated course side}.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TemplateID (Required) ID
// # Path.ID (Required) ID
//
type GetMigrationDetails struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		TemplateID string `json:"template_id" url:"template_id,omitempty"` //  (Required)
		ID         string `json:"id" url:"id,omitempty"`                   //  (Required)
	} `json:"path"`
}

func (t *GetMigrationDetails) GetMethod() string {
	return "GET"
}

func (t *GetMigrationDetails) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/migrations/{id}/details"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetMigrationDetails) GetQuery() (string, error) {
	return "", nil
}

func (t *GetMigrationDetails) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetMigrationDetails) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetMigrationDetails) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'Path.TemplateID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetMigrationDetails) Do(c *canvasapi.Canvas) ([]*models.ChangeRecord, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.ChangeRecord{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
