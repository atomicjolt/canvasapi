package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteSection Delete an existing section.  Returns the former Section.
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # ID (Required) ID
//
type DeleteSection struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteSection) GetMethod() string {
	return "DELETE"
}

func (t *DeleteSection) GetURLPath() string {
	path := "sections/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteSection) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteSection) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteSection) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteSection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteSection) Do(c *canvasapi.Canvas) (*models.Section, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Section{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
