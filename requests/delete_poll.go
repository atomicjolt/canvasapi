package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeletePoll <b>204 No Content</b> response code is returned if the deletion was successful.
// https://canvas.instructure.com/doc/api/polls.html
//
// Path Parameters:
// # ID (Required) ID
//
type DeletePoll struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *DeletePoll) GetMethod() string {
	return "DELETE"
}

func (t *DeletePoll) GetURLPath() string {
	path := "polls/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeletePoll) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePoll) GetBody() (string, error) {
	return "", nil
}

func (t *DeletePoll) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePoll) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
