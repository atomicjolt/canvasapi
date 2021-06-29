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

// DeletePlannerOverride Delete a planner override for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type DeletePlannerOverride struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeletePlannerOverride) GetMethod() string {
	return "DELETE"
}

func (t *DeletePlannerOverride) GetURLPath() string {
	path := "planner/overrides/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeletePlannerOverride) GetQuery() (string, error) {
	return "", nil
}

func (t *DeletePlannerOverride) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeletePlannerOverride) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeletePlannerOverride) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeletePlannerOverride) Do(c *canvasapi.Canvas) (*models.PlannerOverride, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PlannerOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
