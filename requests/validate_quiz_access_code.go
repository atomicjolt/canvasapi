package requests

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ValidateQuizAccessCode Accepts an access code and returns a boolean indicating whether that access code is correct
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # AccessCode (Required) The access code being validated
//
type ValidateQuizAccessCode struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		AccessCode string `json:"access_code"` //  (Required)
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

func (t *ValidateQuizAccessCode) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ValidateQuizAccessCode) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.AccessCode == "" {
		errs = append(errs, "'AccessCode' is required")
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
