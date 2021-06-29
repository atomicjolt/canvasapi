package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateSinglePollChoice Update an existing poll choice for this poll
// https://canvas.instructure.com/doc/api/poll_choices.html
//
// Path Parameters:
// # Path.PollID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.PollChoices.Text (Required) The descriptive text of the poll choice.
// # Form.PollChoices.IsCorrect (Optional) Whether this poll choice is considered correct or not.  Defaults to false.
// # Form.PollChoices.Position (Optional) The order this poll choice should be returned in the context it's sibling poll choices.
//
type UpdateSinglePollChoice struct {
	Path struct {
		PollID string `json:"poll_id" url:"poll_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`

	Form struct {
		PollChoices struct {
			Text      []string `json:"text" url:"text,omitempty"`             //  (Required)
			IsCorrect []string `json:"is_correct" url:"is_correct,omitempty"` //  (Optional)
			Position  []string `json:"position" url:"position,omitempty"`     //  (Optional)
		} `json:"poll_choices" url:"poll_choices,omitempty"`
	} `json:"form"`
}

func (t *UpdateSinglePollChoice) GetMethod() string {
	return "PUT"
}

func (t *UpdateSinglePollChoice) GetURLPath() string {
	path := "polls/{poll_id}/poll_choices/{id}"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateSinglePollChoice) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSinglePollChoice) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateSinglePollChoice) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateSinglePollChoice) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'Path.PollID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.PollChoices.Text == nil {
		errs = append(errs, "'Form.PollChoices.Text' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateSinglePollChoice) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
