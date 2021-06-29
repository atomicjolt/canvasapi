package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListExternalToolsCourses Returns the paginated list of external tools for the current context.
// See the get request docs for a single tool for a list of properties on an external tool.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.SearchTerm (Optional) The partial name of the tools to match and return.
// # Query.Selectable (Optional) If true, then only tools that are meant to be selectable are returned
// # Query.IncludeParents (Optional) If true, then include tools installed in all accounts above the current context
//
type ListExternalToolsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm     string `json:"search_term" url:"search_term,omitempty"`         //  (Optional)
		Selectable     bool   `json:"selectable" url:"selectable,omitempty"`           //  (Optional)
		IncludeParents bool   `json:"include_parents" url:"include_parents,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListExternalToolsCourses) GetMethod() string {
	return "GET"
}

func (t *ListExternalToolsCourses) GetURLPath() string {
	path := "courses/{course_id}/external_tools"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListExternalToolsCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListExternalToolsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListExternalToolsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListExternalToolsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListExternalToolsCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
