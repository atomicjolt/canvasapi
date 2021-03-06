package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListModuleItems A paginated list of the items in a module
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ModuleID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of content_detailsIf included, will return additional details specific to the content
//    associated with each item. Refer to the {api:Modules:Module%20Item Module
//    Item specification} for more details.
//    Includes standard lock information for each item.
// # Query.SearchTerm (Optional) The partial title of the items to match and return.
// # Query.StudentID (Optional) Returns module completion information for the student with this id.
//
type ListModuleItems struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ModuleID string `json:"module_id" url:"module_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include    []string `json:"include" url:"include,omitempty"`         //  (Optional) . Must be one of content_details
		SearchTerm string   `json:"search_term" url:"search_term,omitempty"` //  (Optional)
		StudentID  string   `json:"student_id" url:"student_id,omitempty"`   //  (Optional)
	} `json:"query"`
}

func (t *ListModuleItems) GetMethod() string {
	return "GET"
}

func (t *ListModuleItems) GetURLPath() string {
	path := "courses/{course_id}/modules/{module_id}/items"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{module_id}", fmt.Sprintf("%v", t.Path.ModuleID))
	return path
}

func (t *ListModuleItems) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListModuleItems) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListModuleItems) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListModuleItems) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ModuleID == "" {
		errs = append(errs, "'Path.ModuleID' is required")
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

func (t *ListModuleItems) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.ModuleItem, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.ModuleItem{}
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
