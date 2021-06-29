package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListRecentlyLoggedInStudents Returns the paginated list of users in this course, ordered by how recently they have
// logged in. The records include the 'last_login' field which contains
// a timestamp of the last time that user logged into canvas.  The querying
// user must have the 'View usage reports' permission.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListRecentlyLoggedInStudents struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListRecentlyLoggedInStudents) GetMethod() string {
	return "GET"
}

func (t *ListRecentlyLoggedInStudents) GetURLPath() string {
	path := "courses/{course_id}/recent_students"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListRecentlyLoggedInStudents) GetQuery() (string, error) {
	return "", nil
}

func (t *ListRecentlyLoggedInStudents) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListRecentlyLoggedInStudents) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListRecentlyLoggedInStudents) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRecentlyLoggedInStudents) Do(c *canvasapi.Canvas) ([]*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
