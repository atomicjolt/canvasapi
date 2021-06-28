package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListAvailableTabsForCourseOrGroupAccounts Returns a paginated list of navigation tabs available in the current context.
// https://canvas.instructure.com/doc/api/tabs.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of course_subject_tabs- "course_subject_tabs": Optional flag to return the tabs associated with a canvas_for_elementary subject course's
//      home page instead of the typical sidebar navigation. Only takes effect if this request is for a course context
//      in a canvas_for_elementary-enabled account or sub-account.
//
type ListAvailableTabsForCourseOrGroupAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of course_subject_tabs
	} `json:"query"`
}

func (t *ListAvailableTabsForCourseOrGroupAccounts) GetMethod() string {
	return "GET"
}

func (t *ListAvailableTabsForCourseOrGroupAccounts) GetURLPath() string {
	path := "accounts/{account_id}/tabs"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListAvailableTabsForCourseOrGroupAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAvailableTabsForCourseOrGroupAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAvailableTabsForCourseOrGroupAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAvailableTabsForCourseOrGroupAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
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

func (t *ListAvailableTabsForCourseOrGroupAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
