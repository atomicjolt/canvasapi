package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateQuestionGroup Create a new question group for this quiz
//
// <b>201 Created</b> response code is returned if the creation was successful.
// https://canvas.instructure.com/doc/api/quiz_question_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Form Parameters:
// # QuizGroups (Optional) The name of the question group.
// # QuizGroups (Optional) The number of questions to randomly select for this group.
// # QuizGroups (Optional) The number of points to assign to each question in the group.
// # QuizGroups (Optional) The id of the assessment question bank to pull questions from.
//
type CreateQuestionGroup struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Form struct {
		QuizGroups struct {
			Name                     []string `json:"name" url:"name,omitempty"`                                               //  (Optional)
			PickCount                []int64  `json:"pick_count" url:"pick_count,omitempty"`                                   //  (Optional)
			QuestionPoints           []int64  `json:"question_points" url:"question_points,omitempty"`                         //  (Optional)
			AssessmentQuestionBankID []int64  `json:"assessment_question_bank_id" url:"assessment_question_bank_id,omitempty"` //  (Optional)
		} `json:"quiz_groups" url:"quiz_groups,omitempty"`
	} `json:"form"`
}

func (t *CreateQuestionGroup) GetMethod() string {
	return "POST"
}

func (t *CreateQuestionGroup) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/groups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *CreateQuestionGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateQuestionGroup) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateQuestionGroup) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateQuestionGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateQuestionGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
