package requests

import (
	"fmt"
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
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
	} `json:"path"`

	Form struct {
		QuizGroups struct {
			Name                     []string `json:"name"`                        //  (Optional)
			PickCount                []int64  `json:"pick_count"`                  //  (Optional)
			QuestionPoints           []int64  `json:"question_points"`             //  (Optional)
			AssessmentQuestionBankID []int64  `json:"assessment_question_bank_id"` //  (Optional)
		} `json:"quiz_groups"`
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

func (t *CreateQuestionGroup) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
