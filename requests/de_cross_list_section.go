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

// DeCrossListSection Undo cross-listing of a Section, returning it to its original course.
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # ID (Required) ID
//
type DeCrossListSection struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeCrossListSection) GetMethod() string {
	return "DELETE"
}

func (t *DeCrossListSection) GetURLPath() string {
	path := "sections/{id}/crosslist"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeCrossListSection) GetQuery() (string, error) {
	return "", nil
}

func (t *DeCrossListSection) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeCrossListSection) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeCrossListSection) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeCrossListSection) Do(c *canvasapi.Canvas) (*models.Section, error) {
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
