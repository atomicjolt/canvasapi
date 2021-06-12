package requests

import (
	"fmt"
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
			Question    []string `json:"question"`    //  (Required)
			Description []string `json:"description"` //  (Optional)
		} `json:"polls"`
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

func (t *CreateSinglePoll) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
