package requests

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListAllCourses A paginated list of all courses visible in the public index
// https://canvas.instructure.com/doc/api/search.html
//
// Query Parameters:
// # Search (Optional) Search terms used for matching users/courses/groups (e.g. "bob smith"). If
//    multiple terms are given (separated via whitespace), only results matching
//    all terms will be returned.
// # PublicOnly (Optional) Only return courses with public content. Defaults to false.
// # OpenEnrollmentOnly (Optional) Only return courses that allow self enrollment. Defaults to false.
//
type ListAllCourses struct {
	Query struct {
		Search             string `json:"search" url:"search,omitempty"`                             //  (Optional)
		PublicOnly         bool   `json:"public_only" url:"public_only,omitempty"`                   //  (Optional)
		OpenEnrollmentOnly bool   `json:"open_enrollment_only" url:"open_enrollment_only,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListAllCourses) GetMethod() string {
	return "GET"
}

func (t *ListAllCourses) GetURLPath() string {
	return ""
}

func (t *ListAllCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAllCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListAllCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListAllCourses) HasErrors() error {
	return nil
}

func (t *ListAllCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
