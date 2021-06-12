package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteSingleRubric Deletes a Rubric and removes all RubricAssociations.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type DeleteSingleRubric struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *DeleteSingleRubric) GetMethod() string {
	return "DELETE"
}

func (t *DeleteSingleRubric) GetURLPath() string {
	path := "courses/{course_id}/rubrics/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteSingleRubric) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteSingleRubric) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteSingleRubric) HasErrors() error {
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

func (t *DeleteSingleRubric) Do(c *canvasapi.Canvas) (*models.Rubric, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Rubric{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
