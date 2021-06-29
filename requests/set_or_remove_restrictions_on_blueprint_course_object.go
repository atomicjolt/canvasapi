package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// SetOrRemoveRestrictionsOnBlueprintCourseObject If a blueprint course object is restricted, editing will be limited for copies in associated courses.
// https://canvas.instructure.com/doc/api/blueprint_courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.TemplateID (Required) ID
//
// Form Parameters:
// # Form.ContentType (Optional) . Must be one of assignment, attachment, discussion_topic, external_tool, quiz, wiki_pageThe type of the object.
// # Form.ContentID (Optional) The ID of the object.
// # Form.Restricted (Optional) Whether to apply restrictions.
// # Form.Restrictions (Optional) (Optional) If the object is restricted, this specifies a set of restrictions. If not specified,
//    the course-level restrictions will be used. See {api:CoursesController#update Course API update documentation}
//
type SetOrRemoveRestrictionsOnBlueprintCourseObject struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		TemplateID string `json:"template_id" url:"template_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ContentType  string                       `json:"content_type" url:"content_type,omitempty"` //  (Optional) . Must be one of assignment, attachment, discussion_topic, external_tool, quiz, wiki_page
		ContentID    int64                        `json:"content_id" url:"content_id,omitempty"`     //  (Optional)
		Restricted   bool                         `json:"restricted" url:"restricted,omitempty"`     //  (Optional)
		Restrictions *models.BlueprintRestriction `json:"restrictions" url:"restrictions,omitempty"` //  (Optional)
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

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SetOrRemoveRestrictionsOnBlueprintCourseObject) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.TemplateID == "" {
		errs = append(errs, "'Path.TemplateID' is required")
	}
	if t.Form.ContentType != "" && !string_utils.Include([]string{"assignment", "attachment", "discussion_topic", "external_tool", "quiz", "wiki_page"}, t.Form.ContentType) {
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
