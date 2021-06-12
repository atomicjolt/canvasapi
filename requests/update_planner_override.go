package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdatePlannerOverride Update a planner override's visibilty for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # MarkedComplete (Optional) determines whether the planner item is marked as completed
// # Dismissed (Optional) determines whether the planner item shows in the opportunities list
//
type UpdatePlannerOverride struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		MarkedComplete string `json:"marked_complete"` //  (Optional)
		Dismissed      string `json:"dismissed"`       //  (Optional)
	} `json:"form"`
}

func (t *UpdatePlannerOverride) GetMethod() string {
	return "PUT"
}

func (t *UpdatePlannerOverride) GetURLPath() string {
	path := "planner/overrides/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdatePlannerOverride) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdatePlannerOverride) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdatePlannerOverride) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdatePlannerOverride) Do(c *canvasapi.Canvas) (*models.PlannerOverride, error) {
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
