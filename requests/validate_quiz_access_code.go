package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ValidateQuizAccessCode Accepts an access code and returns a boolean indicating whether that access code is correct
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.AccessCode (Required) The access code being validated
//
type ValidateQuizAccessCode struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		AccessCode string `json:"access_code" url:"access_code,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *ValidateQuizAccessCode) GetMethod() string {
	return "POST"
}

func (t *ValidateQuizAccessCode) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{id}/validate_access_code"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ValidateQuizAccessCode) GetQuery() (string, error) {
	return "", nil
}

func (t *ValidateQuizAccessCode) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ValidateQuizAccessCode) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ValidateQuizAccessCode) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.AccessCode == "" {
		errs = append(errs, "'Form.AccessCode' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ValidateQuizAccessCode) Do(c *canvasapi.Canvas) (bool, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return false, err
	}
	// TODO. I doubt these conversions to string and int below really work. Figure what Canvas returns and test against that return value
	ret := string(body) == "true"

	return ret, nil
}
