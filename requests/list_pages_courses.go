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

// ListPagesCourses A paginated list of the wiki pages associated with a course or group
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.Sort (Optional) . Must be one of title, created_at, updated_atSort results by this field.
// # Query.Order (Optional) . Must be one of asc, descThe sorting order. Defaults to 'asc'.
// # Query.SearchTerm (Optional) The partial title of the pages to match and return.
// # Query.Published (Optional) If true, include only published paqes. If false, exclude published
//    pages. If not present, do not filter on published status.
//
type ListPagesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Sort       string `json:"sort" url:"sort,omitempty"`               //  (Optional) . Must be one of title, created_at, updated_at
		Order      string `json:"order" url:"order,omitempty"`             //  (Optional) . Must be one of asc, desc
		SearchTerm string `json:"search_term" url:"search_term,omitempty"` //  (Optional)
		Published  bool   `json:"published" url:"published,omitempty"`     //  (Optional)
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
	return v.Encode(), nil
}

func (t *ListPagesCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPagesCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPagesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Query.Sort != "" && !string_utils.Include([]string{"title", "created_at", "updated_at"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of title, created_at, updated_at")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
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
