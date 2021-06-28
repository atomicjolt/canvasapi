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

// CreateModuleItem Create and return a new module item
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ModuleID (Required) ID
//
// Form Parameters:
// # ModuleItem (Optional) The name of the module item and associated content
// # ModuleItem (Required) . Must be one of File, Page, Discussion, Assignment, Quiz, SubHeader, ExternalUrl, ExternalToolThe type of content linked to the item
// # ModuleItem (Required) The id of the content to link to the module item. Required, except for
//    'ExternalUrl', 'Page', and 'SubHeader' types.
// # ModuleItem (Optional) The position of this item in the module (1-based).
// # ModuleItem (Optional) 0-based indent level; module items may be indented to show a hierarchy
// # ModuleItem (Optional) Suffix for the linked wiki page (e.g. 'front-page'). Required for 'Page'
//    type.
// # ModuleItem (Optional) External url that the item points to. [Required for 'ExternalUrl' and
//    'ExternalTool' types.
// # ModuleItem (Optional) Whether the external tool opens in a new tab. Only applies to
//    'ExternalTool' type.
// # ModuleItem (Optional) . Must be one of must_view, must_contribute, must_submit, must_mark_doneCompletion requirement for this module item.
//    "must_view": Applies to all item types
//    "must_contribute": Only applies to "Assignment", "Discussion", and "Page" types
//    "must_submit", "min_score": Only apply to "Assignment" and "Quiz" types
//    "must_mark_done": Only applies to "Assignment" and "Page" types
//    Inapplicable types will be ignored
// # ModuleItem (Optional) Minimum score required to complete. Required for completion_requirement
//    type 'min_score'.
//
type CreateModuleItem struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ModuleItem struct {
			Title                 string `json:"title" url:"title,omitempty"`               //  (Optional)
			Type                  string `json:"type" url:"type,omitempty"`                 //  (Required) . Must be one of File, Page, Discussion, Assignment, Quiz, SubHeader, ExternalUrl, ExternalTool
			ContentID             string `json:"content_id" url:"content_id,omitempty"`     //  (Required)
			Position              int64  `json:"position" url:"position,omitempty"`         //  (Optional)
			Indent                int64  `json:"indent" url:"indent,omitempty"`             //  (Optional)
			PageUrl               string `json:"page_url" url:"page_url,omitempty"`         //  (Optional)
			ExternalUrl           string `json:"external_url" url:"external_url,omitempty"` //  (Optional)
			NewTab                bool   `json:"new_tab" url:"new_tab,omitempty"`           //  (Optional)
			CompletionRequirement struct {
				Type     string `json:"type" url:"type,omitempty"`           //  (Optional) . Must be one of must_view, must_contribute, must_submit, must_mark_done
				MinScore int64  `json:"min_score" url:"min_score,omitempty"` //  (Optional)
			} `json:"completion_requirement" url:"completion_requirement,omitempty"`
		} `json:"module_item" url:"module_item,omitempty"`
	} `json:"form"`
}

func (t *CreateModuleItem) GetMethod() string {
	return "POST"
}

func (t *CreateModuleItem) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	return path
}

func (t *CreateModuleItem) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateModuleItem) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateModuleItem) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateModuleItem) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ModuleID == "" {
		errs = append(errs, "'ModuleID' is required")
	}
	if t.Form.ModuleItem.Type == "" {
		errs = append(errs, "'ModuleItem' is required")
	}
	if t.Form.ModuleItem.Type != "" && !string_utils.Include([]string{"File", "Page", "Discussion", "Assignment", "Quiz", "SubHeader", "ExternalUrl", "ExternalTool"}, t.Form.ModuleItem.Type) {
		errs = append(errs, "ModuleItem must be one of File, Page, Discussion, Assignment, Quiz, SubHeader, ExternalUrl, ExternalTool")
	}
	if t.Form.ModuleItem.ContentID == "" {
		errs = append(errs, "'ModuleItem' is required")
	}
	if t.Form.ModuleItem.CompletionRequirement.Type != "" && !string_utils.Include([]string{"must_view", "must_contribute", "must_submit", "must_mark_done"}, t.Form.ModuleItem.CompletionRequirement.Type) {
		errs = append(errs, "ModuleItem must be one of must_view, must_contribute, must_submit, must_mark_done")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateModuleItem) Do(c *canvasapi.Canvas) (*models.ModuleItem, error) {
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
