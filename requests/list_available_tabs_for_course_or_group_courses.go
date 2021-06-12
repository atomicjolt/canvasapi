package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListAvailableTabsForCourseOrGroupCourses Returns a paginated list of navigation tabs available in the current context.
// https://canvas.instructure.com/doc/api/tabs.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of course_subject_tabs- "course_subject_tabs": Optional flag to return the tabs associated with a canvas_for_elementary subject course's
//      home page instead of the typical sidebar navigation. Only takes effect if this request is for a course context
//      in a canvas_for_elementary-enabled account or sub-account.
//
type ListAvailableTabsForCourseOrGroupCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of course_subject_tabs
	} `json:"query"`
}

func (t *ListAvailableTabsForCourseOrGroupCourses) GetMethod() string {
	return "GET"
}

func (t *ListAvailableTabsForCourseOrGroupCourses) GetURLPath() string {
	path := "courses/{course_id}/tabs"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListAvailableTabsForCourseOrGroupCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAvailableTabsForCourseOrGroupCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListAvailableTabsForCourseOrGroupCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"course_subject_tabs"}, v) {
			errs = append(errs, "Include must be one of course_subject_tabs")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAvailableTabsForCourseOrGroupCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
