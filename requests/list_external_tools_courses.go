package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListExternalToolsCourses Returns the paginated list of external tools for the current context.
// See the get request docs for a single tool for a list of properties on an external tool.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # SearchTerm (Optional) The partial name of the tools to match and return.
// # Selectable (Optional) If true, then only tools that are meant to be selectable are returned
// # IncludeParents (Optional) If true, then include tools installed in all accounts above the current context
//
type ListExternalToolsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm     string `json:"search_term"`     //  (Optional)
		Selectable     bool   `json:"selectable"`      //  (Optional)
		IncludeParents bool   `json:"include_parents"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListExternalToolsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListExternalToolsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
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
