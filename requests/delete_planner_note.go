package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeletePlannerNote Delete a planner note for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # ID (Required) ID
//
type DeletePlannerNote struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *DeletePlannerNote) GetMethod() string {
	return "DELETE"
}

func (t *DeletePlannerNote) GetURLPath() string {
	path := "planner_notes/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeletePlannerNote) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePlannerNote) GetBody() (string, error) {
	return "", nil
}

func (t *DeletePlannerNote) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePlannerNote) Do(c *canvasapi.Canvas) (*models.PlannerNote, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PlannerNote{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
