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

// ListGroupCategoriesForContextCourses Returns a paginated list of group categories in a context
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListGroupCategoriesForContextCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListGroupCategoriesForContextCourses) GetMethod() string {
	return "GET"
}

func (t *ListGroupCategoriesForContextCourses) GetURLPath() string {
	path := "courses/{course_id}/group_categories"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListGroupCategoriesForContextCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListGroupCategoriesForContextCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListGroupCategoriesForContextCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListGroupCategoriesForContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupCategoriesForContextCourses) Do(c *canvasapi.Canvas) ([]*models.GroupCategory, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.GroupCategory{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
