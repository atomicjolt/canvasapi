package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DeleteFolder Remove the specified folder. You can only delete empty folders unless you
// set the 'force' flag
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # Force (Optional) Set to 'true' to allow deleting a non-empty folder
//
type DeleteFolder struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Query struct {
		Force bool `json:"force"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *DeleteFolder) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteFolder) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
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