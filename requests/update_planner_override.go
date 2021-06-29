package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdatePlannerOverride Update a planner override's visibilty for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.MarkedComplete (Optional) determines whether the planner item is marked as completed
// # Form.Dismissed (Optional) determines whether the planner item shows in the opportunities list
//
type UpdatePlannerOverride struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		MarkedComplete string `json:"marked_complete" url:"marked_complete,omitempty"` //  (Optional)
		Dismissed      string `json:"dismissed" url:"dismissed,omitempty"`             //  (Optional)
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

func (t *UpdatePlannerOverride) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdatePlannerOverride) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdatePlannerOverride) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
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
