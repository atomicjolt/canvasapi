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

// ShowModule Get information about a single module
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of items, content_details- "items": Return module items inline if possible.
//      This parameter suggests that Canvas return module items directly
//      in the Module object JSON, to avoid having to make separate API
//      requests for each module when enumerating modules and items. Canvas
//      is free to omit 'items' for any particular module if it deems them
//      too numerous to return inline. Callers must be prepared to use the
//      {api:ContextModuleItemsApiController#index List Module Items API}
//      if items are not returned.
//    - "content_details": Requires 'items'. Returns additional
//      details with module items specific to their associated content items.
//      Includes standard lock information for each item.
// # StudentID (Optional) Returns module completion information for the student with this id.
//
type ShowModule struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Include   []string `json:"include" url:"include,omitempty"`       //  (Optional) . Must be one of items, content_details
		StudentID string   `json:"student_id" url:"student_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ShowModule) GetMethod() string {
	return "GET"
}

func (t *ShowModule) GetURLPath() string {
	path := "courses/{course_id}/modules/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowModule) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ShowModule) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowModule) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowModule) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"items", "content_details"}, v) {
			errs = append(errs, "Include must be one of items, content_details")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowModule) Do(c *canvasapi.Canvas) (*models.Module, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Module{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
