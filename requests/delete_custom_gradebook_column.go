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

// DeleteCustomGradebookColumn Permanently deletes a custom column and its associated data
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
type DeleteCustomGradebookColumn struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *DeleteCustomGradebookColumn) GetMethod() string {
	return "DELETE"
}

func (t *DeleteCustomGradebookColumn) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteCustomGradebookColumn) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteCustomGradebookColumn) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteCustomGradebookColumn) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteCustomGradebookColumn) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteCustomGradebookColumn) Do(c *canvasapi.Canvas) (*models.CustomColumn, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CustomColumn{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
