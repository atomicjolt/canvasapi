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

// RetrieveAssignmentOverriddenDatesForNewQuizzes Retrieve the actual due-at, unlock-at, and available-at dates for quizzes
// based on the assignment overrides active for the current API user.
// https://canvas.instructure.com/doc/api/quiz_assignment_overrides.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.QuizAssignmentOverrides (Optional) An array of quiz IDs. If omitted, overrides for all quizzes available to
//    the operating user will be returned.
//
type RetrieveAssignmentOverriddenDatesForNewQuizzes struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		QuizAssignmentOverrides map[string]RetrieveAssignmentOverriddenDatesForNewQuizzesQuizAssignmentOverrides
	} `json:"query"`
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) GetMethod() string {
	return "GET"
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) GetURLPath() string {
	path := "courses/{course_id}/new_quizzes/assignment_overrides"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RetrieveAssignmentOverriddenDatesForNewQuizzes) Do(c *canvasapi.Canvas) (*models.QuizAssignmentOverrideSetContainer, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.QuizAssignmentOverrideSetContainer{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type RetrieveAssignmentOverriddenDatesForNewQuizzesQuizAssignmentOverrides struct {
	QuizIDs []string `json:"quiz_ids" url:"quiz_ids,omitempty"` //  (Optional)
}
