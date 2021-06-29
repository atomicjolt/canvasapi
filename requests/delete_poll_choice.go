package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeletePollChoice <b>204 No Content</b> response code is returned if the deletion was successful.
// https://canvas.instructure.com/doc/api/poll_choices.html
//
// Path Parameters:
// # Path.PollID (Required) ID
// # Path.ID (Required) ID
//
type DeletePollChoice struct {
	Path struct {
		PollID string `json:"poll_id" url:"poll_id,omitempty"` //  (Required)
		ID     string `json:"id" url:"id,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *DeletePollChoice) GetMethod() string {
	return "DELETE"
}

func (t *DeletePollChoice) GetURLPath() string {
	path := "polls/{poll_id}/poll_choices/{id}"
	path = strings.ReplaceAll(path, "{poll_id}", fmt.Sprintf("%v", t.Path.PollID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeletePollChoice) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePollChoice) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeletePollChoice) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeletePollChoice) HasErrors() error {
	errs := []string{}
	if t.Path.PollID == "" {
		errs = append(errs, "'Path.PollID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePollChoice) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
