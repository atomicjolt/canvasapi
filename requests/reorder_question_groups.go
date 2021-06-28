package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
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
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		Order struct {
			ID   []int64  `json:"id" url:"id,omitempty"`     //  (Required)
			Type []string `json:"type" url:"type,omitempty"` //  (Optional) . Must be one of question
		} `json:"order" url:"order,omitempty"`
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

func (t *ReorderQuestionGroups) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ReorderQuestionGroups) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
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
		if v != "" && !string_utils.Include([]string{"question"}, v) {
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
