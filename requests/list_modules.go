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

// ListModules A paginated list of the modules in a course
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of items, content_details- "items": Return module items inline if possible.
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
// # Query.SearchTerm (Optional) The partial name of the modules (and module items, if 'items' is
//    specified with include[]) to match and return.
// # Query.StudentID (Optional) Returns module completion information for the student with this id.
//
type ListModules struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include    []string `json:"include" url:"include,omitempty"`         //  (Optional) . Must be one of items, content_details
		SearchTerm string   `json:"search_term" url:"search_term,omitempty"` //  (Optional)
		StudentID  string   `json:"student_id" url:"student_id,omitempty"`   //  (Optional)
	} `json:"query"`
}

func (t *ListModules) GetMethod() string {
	return "GET"
}

func (t *ListModules) GetURLPath() string {
	path := "courses/{course_id}/modules"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListModules) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListModules) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListModules) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListModules) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
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

func (t *ListModules) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Module, *canvasapi.PagedResource, error) {
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
	ret := []*models.Module{}
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
