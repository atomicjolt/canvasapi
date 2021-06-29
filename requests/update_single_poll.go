package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateSinglePoll Update an existing poll belonging to the current user
// https://canvas.instructure.com/doc/api/polls.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Polls.Question (Required) The title of the poll.
// # Form.Polls.Description (Optional) A brief description or instructions for the poll.
//
type UpdateSinglePoll struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Polls struct {
			Question    []string `json:"question" url:"question,omitempty"`       //  (Required)
			Description []string `json:"description" url:"description,omitempty"` //  (Optional)
		} `json:"polls" url:"polls,omitempty"`
	} `json:"form"`
}

func (t *UpdateSinglePoll) GetMethod() string {
	return "PUT"
}

func (t *UpdateSinglePoll) GetURLPath() string {
	path := "polls/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateSinglePoll) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSinglePoll) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateSinglePoll) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateSinglePoll) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.Polls.Question == nil {
		errs = append(errs, "'Form.Polls.Question' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateSinglePoll) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
