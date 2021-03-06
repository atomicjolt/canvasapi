package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteBookmark Deletes a bookmark
// https://canvas.instructure.com/doc/api/bookmarks.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type DeleteBookmark struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteBookmark) GetMethod() string {
	return "DELETE"
}

func (t *DeleteBookmark) GetURLPath() string {
	path := "users/self/bookmarks/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteBookmark) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteBookmark) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteBookmark) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteBookmark) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteBookmark) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
