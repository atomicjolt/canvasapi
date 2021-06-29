package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListAvailableTabsForCourseOrGroupUsers Returns a paginated list of navigation tabs available in the current context.
// https://canvas.instructure.com/doc/api/tabs.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of course_subject_tabs- "course_subject_tabs": Optional flag to return the tabs associated with a canvas_for_elementary subject course's
//      home page instead of the typical sidebar navigation. Only takes effect if this request is for a course context
//      in a canvas_for_elementary-enabled account or sub-account.
//
type ListAvailableTabsForCourseOrGroupUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of course_subject_tabs
	} `json:"query"`
}

func (t *ListAvailableTabsForCourseOrGroupUsers) GetMethod() string {
	return "GET"
}

func (t *ListAvailableTabsForCourseOrGroupUsers) GetURLPath() string {
	path := "users/{user_id}/tabs"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListAvailableTabsForCourseOrGroupUsers) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListAvailableTabsForCourseOrGroupUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAvailableTabsForCourseOrGroupUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAvailableTabsForCourseOrGroupUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"course_subject_tabs"}, v) {
			errs = append(errs, "Include must be one of course_subject_tabs")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAvailableTabsForCourseOrGroupUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
