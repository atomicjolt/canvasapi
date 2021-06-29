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

// ListSubgroupsGlobal A paginated list of the immediate OutcomeGroup children of the outcome group.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type ListSubgroupsGlobal struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListSubgroupsGlobal) GetMethod() string {
	return "GET"
}

func (t *ListSubgroupsGlobal) GetURLPath() string {
	path := "global/outcome_groups/{id}/subgroups"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListSubgroupsGlobal) GetQuery() (string, error) {
	return "", nil
}

func (t *ListSubgroupsGlobal) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListSubgroupsGlobal) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListSubgroupsGlobal) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListSubgroupsGlobal) Do(c *canvasapi.Canvas) ([]*models.OutcomeGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.OutcomeGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
