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

// ListPagesCourses A paginated list of the wiki pages associated with a course or group
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Sort (Optional) . Must be one of title, created_at, updated_atSort results by this field.
// # Order (Optional) . Must be one of asc, descThe sorting order. Defaults to 'asc'.
// # SearchTerm (Optional) The partial title of the pages to match and return.
// # Published (Optional) If true, include only published paqes. If false, exclude published
//    pages. If not present, do not filter on published status.
//
type ListPagesCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Sort       string `json:"sort"`        //  (Optional) . Must be one of title, created_at, updated_at
		Order      string `json:"order"`       //  (Optional) . Must be one of asc, desc
		SearchTerm string `json:"search_term"` //  (Optional)
		Published  bool   `json:"published"`   //  (Optional)
	} `json:"query"`
}

func (t *ListPagesCourses) GetMethod() string {
	return "GET"
}

func (t *ListPagesCourses) GetURLPath() string {
	path := "courses/{course_id}/pages"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListPagesCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListPagesCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListPagesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"title", "created_at", "updated_at"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of title, created_at, updated_at")
	}
	if !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPagesCourses) Do(c *canvasapi.Canvas) ([]*models.Page, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Page{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
