package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateAssociatedCourses Send a list of course ids to add or remove new associations for the template.
// Cannot add courses that do not belong to the blueprint course's account. Also cannot add
// other blueprint courses or courses that already have an association with another blueprint course.
//
// After associating new courses, {api:MasterCourses::MasterTemplatesController#queue_migration start a sync} to populate their contents from the blueprint.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TemplateID (Required) ID
//
// Form Parameters:
// # Form.CourseIDsToAdd (Optional) Courses to add as associated courses
// # Form.CourseIDsToRemove (Optional) Courses to remove as associated courses
//
type UpdateAssociatedCourses struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		TemplateID string `json:"template_id" url:"template_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseIDsToAdd    string `json:"course_ids_to_add" url:"course_ids_to_add,omitempty"`       //  (Optional)
		CourseIDsToRemove string `json:"course_ids_to_remove" url:"course_ids_to_remove,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateAssociatedCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateAssociatedCourses) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/update_associations"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	return path
}

func (t *UpdateAssociatedCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateAssociatedCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateAssociatedCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateAssociatedCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'Path.TemplateID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateAssociatedCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
