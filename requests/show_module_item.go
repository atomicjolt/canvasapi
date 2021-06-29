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

// ShowModuleItem Get information about a single module item
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ModuleID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of content_detailsIf included, will return additional details specific to the content
//    associated with this item. Refer to the {api:Modules:Module%20Item Module
//    Item specification} for more details.
//    Includes standard lock information for each item.
// # StudentID (Optional) Returns module completion information for the student with this id.
//
type ShowModuleItem struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Include   []string `json:"include" url:"include,omitempty"`       //  (Optional) . Must be one of content_details
		StudentID string   `json:"student_id" url:"student_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ShowModuleItem) GetMethod() string {
	return "GET"
}

func (t *ShowModuleItem) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowModuleItem) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ShowModuleItem) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowModuleItem) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowModuleItem) HasErrors() error {
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
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"content_details"}, v) {
			errs = append(errs, "Include must be one of content_details")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowModuleItem) Do(c *canvasapi.Canvas) (*models.ModuleItem, error) {
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
