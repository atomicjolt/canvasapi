package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateSinglePoll Create a new poll for the current user
// https://canvas.instructure.com/doc/api/polls.html
//
// Form Parameters:
// # Polls (Required) The title of the poll.
// # Polls (Optional) A brief description or instructions for the poll.
//
type CreateSinglePoll struct {
	Form struct {
		Polls struct {
			Question    []string `json:"question" url:"question,omitempty"`       //  (Required)
			Description []string `json:"description" url:"description,omitempty"` //  (Optional)
		} `json:"polls" url:"polls,omitempty"`
	} `json:"form"`
}

func (t *CreateSinglePoll) GetMethod() string {
	return "POST"
}

func (t *CreateSinglePoll) GetURLPath() string {
	return ""
}

func (t *CreateSinglePoll) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSinglePoll) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateSinglePoll) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateSinglePoll) HasErrors() error {
	errs := []string{}
	if t.Form.Polls.Question == nil {
		errs = append(errs, "'Polls' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSinglePoll) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
