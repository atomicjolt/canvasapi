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

// ListPotentialMembersCourses A paginated list of the users who can potentially be added to a
// collaboration in the given context.
//
// For courses, this consists of all enrolled users.  For groups, it is comprised of the
// group members plus the admins of the course containing the group.
// https://canvas.instructure.com/doc/api/collaborations.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type ListPotentialMembersCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListPotentialMembersCourses) GetMethod() string {
	return "GET"
}

func (t *ListPotentialMembersCourses) GetURLPath() string {
	path := "courses/{course_id}/potential_collaborators"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListPotentialMembersCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPotentialMembersCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPotentialMembersCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPotentialMembersCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListPotentialMembersCourses) Do(c *canvasapi.Canvas) ([]*models.User, error) {
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
