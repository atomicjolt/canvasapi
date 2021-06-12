package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// GetModuleItemSequence Given an asset in a course, find the ModuleItem it belongs to, the previous and next Module Items
// in the course sequence, and also any applicable mastery path rules
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # AssetType (Optional) . Must be one of ModuleItem, File, Page, Discussion, Assignment, Quiz, ExternalToolThe type of asset to find module sequence information for. Use the ModuleItem if it is known
//    (e.g., the user navigated from a module item), since this will avoid ambiguity if the asset
//    appears more than once in the module sequence.
// # AssetID (Optional) The id of the asset (or the url in the case of a Page)
//
type GetModuleItemSequence struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		AssetType string `json:"asset_type"` //  (Optional) . Must be one of ModuleItem, File, Page, Discussion, Assignment, Quiz, ExternalTool
		AssetID   int64  `json:"asset_id"`   //  (Optional)
	} `json:"query"`
}

func (t *GetModuleItemSequence) GetMethod() string {
	return "GET"
}

func (t *GetModuleItemSequence) GetURLPath() string {
	path := "courses/{course_id}/module_item_sequence"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetModuleItemSequence) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetModuleItemSequence) GetBody() (string, error) {
	return "", nil
}

func (t *GetModuleItemSequence) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"ModuleItem", "File", "Page", "Discussion", "Assignment", "Quiz", "ExternalTool"}, t.Query.AssetType) {
		errs = append(errs, "AssetType must be one of ModuleItem, File, Page, Discussion, Assignment, Quiz, ExternalTool")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetModuleItemSequence) Do(c *canvasapi.Canvas) (*models.ModuleItemSequence, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ModuleItemSequence{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
