package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// SetCourseNickname Set a nickname for the given course. This will replace the course's name
// in output of API calls you make subsequently, as well as in selected
// places in the Canvas web user interface.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Nickname (Required) The nickname to set.  It must be non-empty and shorter than 60 characters.
//
type SetCourseNickname struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Nickname string `json:"nickname" url:"nickname,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *SetCourseNickname) GetMethod() string {
	return "PUT"
}

func (t *SetCourseNickname) GetURLPath() string {
	path := "users/self/course_nicknames/{course_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *SetCourseNickname) GetQuery() (string, error) {
	return "", nil
}

func (t *SetCourseNickname) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SetCourseNickname) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SetCourseNickname) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Nickname == "" {
		errs = append(errs, "'Nickname' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetCourseNickname) Do(c *canvasapi.Canvas) (*models.CourseNickname, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CourseNickname{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
