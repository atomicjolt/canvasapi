package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateSinglePollChoice Create a new poll choice for this poll
// https://canvas.instructure.com/doc/api/poll_choices.html
//
// Path Parameters:
// # PollID (Required) ID
//
// Form Parameters:
// # PollChoices (Required) The descriptive text of the poll choice.
// # PollChoices (Optional) Whether this poll choice is considered correct or not. Defaults to false.
// # PollChoices (Optional) The order this poll choice should be returned in the context it's sibling poll choices.
//
type CreateSinglePollChoice struct {
	Path struct {
		PollID string `json:"poll_id"` //  (Required)
	} `json:"path"`

	Form struct {
		PollChoices struct {
			Text      []string `json:"text"`       //  (Required)
			IsCorrect []bool   `json:"is_correct"` //  (Optional)
			Position  []int64  `json:"position"`   //  (Optional)
		} `json:"poll_choices"`
	} `json:"form"`
}

func (t *CreateSinglePollChoice) GetMethod() string {
	return "POST"
}

func (t *CreateSinglePollChoice) GetURLPath() string {
	path := "polls/{poll_id}/poll_choices"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	return path
}

func (t *CreateSinglePollChoice) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSinglePollChoice) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateSinglePollChoice) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Form.PollChoices.Text == nil {
		errs = append(errs, "'PollChoices' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSinglePollChoice) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
