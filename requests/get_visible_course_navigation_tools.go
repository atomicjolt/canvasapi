package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetVisibleCourseNavigationTools Get a list of external tools with the course_navigation placement that have not been hidden in
// course settings and whose visibility settings apply to the requesting user. These tools are the
// same that appear in the course navigation.
//
// The response format is the same as for List external tools, but with additional context_id and
// context_name fields on each element in the array.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Query Parameters:
// # ContextCodes (Required) List of context_codes to retrieve visible course nav tools for (for example, +course_123+). Only
//    courses are presently supported.
//
type GetVisibleCourseNavigationTools struct {
	Query struct {
		ContextCodes []string `json:"context_codes"` //  (Required)
	} `json:"query"`
}

func (t *GetVisibleCourseNavigationTools) GetMethod() string {
	return "GET"
}

func (t *GetVisibleCourseNavigationTools) GetURLPath() string {
	return ""
}

func (t *GetVisibleCourseNavigationTools) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetVisibleCourseNavigationTools) GetBody() (string, error) {
	return "", nil
}

func (t *GetVisibleCourseNavigationTools) HasErrors() error {
	errs := []string{}
	if t.Query.ContextCodes == nil {
		errs = append(errs, "'ContextCodes' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetVisibleCourseNavigationTools) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
