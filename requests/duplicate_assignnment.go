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
	"github.com/atomicjolt/string_utils"
)

// DuplicateAssignnment Duplicate an assignment and return a json based on result_type argument.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
// Form Parameters:
// # ResultType (Optional) . Must be one of QuizOptional information:
//    When the root account has the feature `newquizzes_on_quiz_page` enabled
//    and this argument is set to "Quiz" the response will be serialized into a
//    quiz format({file:doc/api/quizzes.html#Quiz});
//    When this argument isn't specified the response will be serialized into an
//    assignment format;
//
type DuplicateAssignnment struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ResultType string `json:"result_type" url:"result_type,omitempty"` //  (Optional) . Must be one of Quiz
	} `json:"form"`
}

func (t *DuplicateAssignnment) GetMethod() string {
	return "POST"
}

func (t *DuplicateAssignnment) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/duplicate"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *DuplicateAssignnment) GetQuery() (string, error) {
	return "", nil
}

func (t *DuplicateAssignnment) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *DuplicateAssignnment) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *DuplicateAssignnment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Form.ResultType != "" && !string_utils.Include([]string{"Quiz"}, t.Form.ResultType) {
		errs = append(errs, "ResultType must be one of Quiz")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DuplicateAssignnment) Do(c *canvasapi.Canvas) (*models.Assignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Assignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
