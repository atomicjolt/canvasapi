package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetSinglePoll Returns the poll with the given id
// https://canvas.instructure.com/doc/api/polls.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type GetSinglePoll struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetSinglePoll) GetMethod() string {
	return "GET"
}

func (t *GetSinglePoll) GetURLPath() string {
	path := "polls/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSinglePoll) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSinglePoll) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSinglePoll) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSinglePoll) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSinglePoll) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
