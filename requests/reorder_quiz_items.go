package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ReorderQuizItems Change order of the quiz questions or groups within the quiz
//
// <b>204 No Content</b> response code is returned if the reorder was successful.
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Order (Required) The associated item's unique identifier
// # Order (Optional) . Must be one of question, groupThe type of item is either 'question' or 'group'
//
type ReorderQuizItems struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Order struct {
			ID   []int64  `json:"id"`   //  (Required)
			Type []string `json:"type"` //  (Optional) . Must be one of question, group
		} `json:"order"`
	} `json:"form"`
}

func (t *ReorderQuizItems) GetMethod() string {
	return "POST"
}

func (t *ReorderQuizItems) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{id}/reorder"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ReorderQuizItems) GetQuery() (string, error) {
	return "", nil
}

func (t *ReorderQuizItems) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ReorderQuizItems) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Order.ID == nil {
		errs = append(errs, "'Order' is required")
	}
	for _, v := range t.Form.Order.Type {
		if !string_utils.Include([]string{"question", "group"}, v) {
			errs = append(errs, "Order must be one of question, group")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReorderQuizItems) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
