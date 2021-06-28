package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListCountsForTodoItems Counts of different todo items such as the number of assignments needing grading as well as the number of assignments needing submitting.
//
// There is a limit to the number of todo items this endpoint will count.
// It will only look at the first 100 todo items for the user. If the user has more than 100 todo items this count may not be reliable.
// The largest reliable number for both counts is 100.
// https://canvas.instructure.com/doc/api/users.html
//
// Query Parameters:
// # Include (Optional) . Must be one of ungraded_quizzes"ungraded_quizzes":: Optionally include ungraded quizzes (such as practice quizzes and surveys) in the list.
//                         These will be returned under a +quiz+ key instead of an +assignment+ key in response elements.
//
type ListCountsForTodoItems struct {
	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of ungraded_quizzes
	} `json:"query"`
}

func (t *ListCountsForTodoItems) GetMethod() string {
	return "GET"
}

func (t *ListCountsForTodoItems) GetURLPath() string {
	return ""
}

func (t *ListCountsForTodoItems) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListCountsForTodoItems) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCountsForTodoItems) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCountsForTodoItems) HasErrors() error {
	errs := []string{}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"ungraded_quizzes"}, v) {
			errs = append(errs, "Include must be one of ungraded_quizzes")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCountsForTodoItems) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
