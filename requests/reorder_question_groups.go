package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ReorderQuestionGroups Change the order of the quiz questions within the group
//
// <b>204 No Content<b> response code is returned if the reorder was successful.
// https://canvas.instructure.com/doc/api/quiz_question_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Order (Required) The associated item's unique identifier
// # Order (Optional) . Must be one of questionThe type of item is always 'question' for a group
//
type ReorderQuestionGroups struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Order struct {
			ID   []int64  `json:"id"`   //  (Required)
			Type []string `json:"type"` //  (Optional) . Must be one of question
		} `json:"order"`
	} `json:"form"`
}

func (t *ReorderQuestionGroups) GetMethod() string {
	return "POST"
}

func (t *ReorderQuestionGroups) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/groups/{id}/reorder"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ReorderQuestionGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ReorderQuestionGroups) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ReorderQuestionGroups) HasErrors() error {
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
	if t.Form.Order.ID == nil {
		errs = append(errs, "'Order' is required")
	}
	for _, v := range t.Form.Order.Type {
		if !string_utils.Include([]string{"question"}, v) {
			errs = append(errs, "Order must be one of question")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReorderQuestionGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
