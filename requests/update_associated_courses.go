package requests

import (
	"fmt"
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
// # CourseID (Required) ID
// # TemplateID (Required) ID
//
// Form Parameters:
// # CourseIDsToAdd (Optional) Courses to add as associated courses
// # CourseIDsToRemove (Optional) Courses to remove as associated courses
//
type UpdateAssociatedCourses struct {
	Path struct {
		CourseID   string `json:"course_id"`   //  (Required)
		TemplateID string `json:"template_id"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseIDsToAdd    string `json:"course_ids_to_add"`    //  (Optional)
		CourseIDsToRemove string `json:"course_ids_to_remove"` //  (Optional)
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

func (t *UpdateAssociatedCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateAssociatedCourses) HasErrors() error {
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

func (t *UpdateAssociatedCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
