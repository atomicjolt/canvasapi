package requests

import (
	"encoding/json"
	"io/ioutil"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListPlannerOverrides Retrieve a planner override for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
type ListPlannerOverrides struct {
}

func (t *ListPlannerOverrides) GetMethod() string {
	return "GET"
}

func (t *ListPlannerOverrides) GetURLPath() string {
	return ""
}

func (t *ListPlannerOverrides) GetQuery() (string, error) {
	return "", nil
}

func (t *ListPlannerOverrides) GetBody() (string, error) {
	return "", nil
}

func (t *ListPlannerOverrides) HasErrors() error {
	return nil
}

func (t *ListPlannerOverrides) Do(c *canvasapi.Canvas) ([]*models.PlannerOverride, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.PlannerOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}