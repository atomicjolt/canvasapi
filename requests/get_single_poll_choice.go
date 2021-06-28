package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSinglePollChoice Returns the poll choice with the given id
// https://canvas.instructure.com/doc/api/poll_choices.html
//
// Path Parameters:
// # PollID (Required) ID
// # ID (Required) ID
//
type GetSinglePollChoice struct {
	Path struct {
		PollID string `json:"poll_id" url:"poll_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *GetSinglePollChoice) GetMethod() string {
	return "GET"
}

func (t *GetSinglePollChoice) GetURLPath() string {
	path := "polls/{poll_id}/poll_choices/{id}"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSinglePollChoice) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSinglePollChoice) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSinglePollChoice) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSinglePollChoice) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'PollID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSinglePollChoice) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
