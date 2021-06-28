package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ConfirmImageSelection After you have used the search API, you should hit this API to indicate photo usage to the server.
// https://canvas.instructure.com/doc/api/image_search.html
//
// Path Parameters:
// # ID (Required) The ID from the image_search result.
//
type ConfirmImageSelection struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ConfirmImageSelection) GetMethod() string {
	return "POST"
}

func (t *ConfirmImageSelection) GetURLPath() string {
	path := "image_selection/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ConfirmImageSelection) GetQuery() (string, error) {
	return "", nil
}

func (t *ConfirmImageSelection) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ConfirmImageSelection) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ConfirmImageSelection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ConfirmImageSelection) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
