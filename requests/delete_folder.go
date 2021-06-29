package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DeleteFolder Remove the specified folder. You can only delete empty folders unless you
// set the 'force' flag
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Force (Optional) Set to 'true' to allow deleting a non-empty folder
//
type DeleteFolder struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Force bool `json:"force" url:"force,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *DeleteFolder) GetMethod() string {
	return "DELETE"
}

func (t *DeleteFolder) GetURLPath() string {
	path := "folders/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteFolder) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *DeleteFolder) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteFolder) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteFolder) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteFolder) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
