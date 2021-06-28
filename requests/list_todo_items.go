package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListTodoItems A paginated list of the current user's list of todo items.
//
// There is a limit to the number of items returned.
//
// The `ignore` and `ignore_permanently` URLs can be used to update the user's
// preferences on what items will be displayed.
// Performing a DELETE request against the `ignore` URL will hide that item
// from future todo item requests, until the item changes.
// Performing a DELETE request against the `ignore_permanently` URL will hide
// that item forever.
// https://canvas.instructure.com/doc/api/users.html
//
// Query Parameters:
// # Include (Optional) . Must be one of ungraded_quizzes"ungraded_quizzes":: Optionally include ungraded quizzes (such as practice quizzes and surveys) in the list.
//                         These will be returned under a +quiz+ key instead of an +assignment+ key in response elements.
//
type ListTodoItems struct {
	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of ungraded_quizzes
	} `json:"query"`
}

func (t *ListTodoItems) GetMethod() string {
	return "GET"
}

func (t *ListTodoItems) GetURLPath() string {
	return ""
}

func (t *ListTodoItems) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListTodoItems) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListTodoItems) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListTodoItems) HasErrors() error {
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

func (t *ListTodoItems) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
