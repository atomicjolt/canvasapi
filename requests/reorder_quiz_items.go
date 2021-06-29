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

// ReorderQuizItems Change order of the quiz questions or groups within the quiz
//
// <b>204 No Content</b> response code is returned if the reorder was successful.
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Order.ID (Required) The associated item's unique identifier
// # Form.Order.Type (Optional) . Must be one of question, groupThe type of item is either 'question' or 'group'
//
type ReorderQuizItems struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		Order struct {
			ID   []string `json:"id" url:"id,omitempty"`     //  (Required)
			Type []string `json:"type" url:"type,omitempty"` //  (Optional) . Must be one of question, group
		} `json:"order" url:"order,omitempty"`
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

func (t *ReorderQuizItems) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ReorderQuizItems) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ReorderQuizItems) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.Order.ID == nil {
		errs = append(errs, "'Form.Order.ID' is required")
	}
	for _, v := range t.Form.Order.Type {
		if v != "" && !string_utils.Include([]string{"question", "group"}, v) {
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
