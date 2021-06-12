package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// SetOrRemoveRestrictionsOnBlueprintCourseObject If a blueprint course object is restricted, editing will be limited for copies in associated courses.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TemplateID (Required) ID
//
// Form Parameters:
// # ContentType (Optional) . Must be one of assignment, attachment, discussion_topic, external_tool, quiz, wiki_pageThe type of the object.
// # ContentID (Optional) The ID of the object.
// # Restricted (Optional) Whether to apply restrictions.
// # Restrictions (Optional) (Optional) If the object is restricted, this specifies a set of restrictions. If not specified,
//    the course-level restrictions will be used. See {api:CoursesController#update Course API update documentation}
//
type SetOrRemoveRestrictionsOnBlueprintCourseObject struct {
	Path struct {
		CourseID   string `json:"course_id"`   //  (Required)
		TemplateID string `json:"template_id"` //  (Required)
	} `json:"path"`

	Form struct {
		ContentType  string `json:"content_type"` //  (Optional) . Must be one of assignment, attachment, discussion_topic, external_tool, quiz, wiki_page
		ContentID    int64  `json:"content_id"`   //  (Optional)
		Restricted   bool   `json:"restricted"`   //  (Optional)
		Restrictions string `json:"restrictions"` //  (Optional)
	} `json:"form"`
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) GetMethod() string {
	return "PUT"
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) GetURLPath() string {
	path := "courses/{course_id}/blueprint_templates/{template_id}/restrict_item"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{template_id}", fmt.Sprintf("%v", t.Path.TemplateID))
	return path
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) GetQuery() (string, error) {
	return "", nil
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'TemplateID' is required")
	}
	if !string_utils.Include([]string{"assignment", "attachment", "discussion_topic", "external_tool", "quiz", "wiki_page"}, t.Form.ContentType) {
		errs = append(errs, "ContentType must be one of assignment, attachment, discussion_topic, external_tool, quiz, wiki_page")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
