package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateCustomGradebookColumn Accepts the same parameters as custom gradebook column creation
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type UpdateCustomGradebookColumn struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *UpdateCustomGradebookColumn) GetMethod() string {
	return "PUT"
}

func (t *UpdateCustomGradebookColumn) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateCustomGradebookColumn) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCustomGradebookColumn) GetBody() (string, error) {
	return "", nil
}

func (t *UpdateCustomGradebookColumn) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCustomGradebookColumn) Do(c *canvasapi.Canvas) (*models.CustomColumn, error) {
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
