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

// ShowPlannerNote Retrieve a planner note for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # ID (Required) ID
//
type ShowPlannerNote struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ShowPlannerNote) GetMethod() string {
	return "GET"
}

func (t *ShowPlannerNote) GetURLPath() string {
	path := "planner_notes/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowPlannerNote) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowPlannerNote) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowPlannerNote) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowPlannerNote) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowPlannerNote) Do(c *canvasapi.Canvas) (*models.PlannerNote, error) {
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
