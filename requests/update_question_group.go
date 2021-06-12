package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateQuestionGroup Update a question group
// https://canvas.instructure.com/doc/api/quiz_question_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # QuizGroups (Optional) The name of the question group.
// # QuizGroups (Optional) The number of questions to randomly select for this group.
// # QuizGroups (Optional) The number of points to assign to each question in the group.
//
type UpdateQuestionGroup struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		QuizGroups struct {
			Name           []string `json:"name"`            //  (Optional)
			PickCount      []int64  `json:"pick_count"`      //  (Optional)
			QuestionPoints []int64  `json:"question_points"` //  (Optional)
		} `json:"quiz_groups"`
	} `json:"form"`
}

func (t *UpdateQuestionGroup) GetMethod() string {
	return "PUT"
}

func (t *UpdateQuestionGroup) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/groups/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateQuestionGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateQuestionGroup) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateQuestionGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateQuestionGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
