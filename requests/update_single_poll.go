package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateSinglePoll Update an existing poll belonging to the current user
// https://canvas.instructure.com/doc/api/polls.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Polls (Required) The title of the poll.
// # Polls (Optional) A brief description or instructions for the poll.
//
type UpdateSinglePoll struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Polls struct {
			Question    []string `json:"question"`    //  (Required)
			Description []string `json:"description"` //  (Optional)
		} `json:"polls"`
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

func (t *UpdateSinglePoll) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateSinglePoll) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Polls.Question == nil {
		errs = append(errs, "'Polls' is required")
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
