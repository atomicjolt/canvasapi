package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

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

func (t *ListPlannerOverrides) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPlannerOverrides) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPlannerOverrides) HasErrors() error {
	return nil
}

func (t *ListPlannerOverrides) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.PlannerOverride, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.PlannerOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
