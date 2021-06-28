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
	"github.com/atomicjolt/string_utils"
)

// UpdateModuleItem Update and return an existing module item
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ModuleID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # ModuleItem (Optional) The name of the module item
// # ModuleItem (Optional) The position of this item in the module (1-based)
// # ModuleItem (Optional) 0-based indent level; module items may be indented to show a hierarchy
// # ModuleItem (Optional) External url that the item points to. Only applies to 'ExternalUrl' type.
// # ModuleItem (Optional) Whether the external tool opens in a new tab. Only applies to
//    'ExternalTool' type.
// # ModuleItem (Optional) . Must be one of must_view, must_contribute, must_submit, must_mark_doneCompletion requirement for this module item.
//    "must_view": Applies to all item types
//    "must_contribute": Only applies to "Assignment", "Discussion", and "Page" types
//    "must_submit", "min_score": Only apply to "Assignment" and "Quiz" types
//    "must_mark_done": Only applies to "Assignment" and "Page" types
//    Inapplicable types will be ignored
// # ModuleItem (Optional) Minimum score required to complete, Required for completion_requirement
//    type 'min_score'.
// # ModuleItem (Optional) Whether the module item is published and visible to students.
// # ModuleItem (Optional) Move this item to another module by specifying the target module id here.
//    The target module must be in the same course.
//
type UpdateModuleItem struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		ModuleItem struct {
			Title                 string `json:"title" url:"title,omitempty"`               //  (Optional)
			Position              int64  `json:"position" url:"position,omitempty"`         //  (Optional)
			Indent                int64  `json:"indent" url:"indent,omitempty"`             //  (Optional)
			ExternalUrl           string `json:"external_url" url:"external_url,omitempty"` //  (Optional)
			NewTab                bool   `json:"new_tab" url:"new_tab,omitempty"`           //  (Optional)
			CompletionRequirement struct {
				Type     string `json:"type" url:"type,omitempty"`           //  (Optional) . Must be one of must_view, must_contribute, must_submit, must_mark_done
				MinScore int64  `json:"min_score" url:"min_score,omitempty"` //  (Optional)
			} `json:"completion_requirement" url:"completion_requirement,omitempty"`

			Published bool   `json:"published" url:"published,omitempty"` //  (Optional)
			ModuleID  string `json:"module_id" url:"module_id,omitempty"` //  (Optional)
		} `json:"module_item" url:"module_item,omitempty"`
	} `json:"form"`
}

func (t *UpdateModuleItem) GetMethod() string {
	return "PUT"
}

func (t *UpdateModuleItem) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateModuleItem) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateModuleItem) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateModuleItem) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateModuleItem) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ModuleID == "" {
		errs = append(errs, "'ModuleID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.ModuleItem.CompletionRequirement.Type != "" && !string_utils.Include([]string{"must_view", "must_contribute", "must_submit", "must_mark_done"}, t.Form.ModuleItem.CompletionRequirement.Type) {
		errs = append(errs, "ModuleItem must be one of must_view, must_contribute, must_submit, must_mark_done")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateModuleItem) Do(c *canvasapi.Canvas) (*models.ModuleItem, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ModuleItem{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
