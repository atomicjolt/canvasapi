package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowOutcomeGroupGlobal
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # ID (Required) ID
//
type ShowOutcomeGroupGlobal struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *ShowOutcomeGroupGlobal) GetMethod() string {
	return "GET"
}

func (t *ShowOutcomeGroupGlobal) GetURLPath() string {
	path := "global/outcome_groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowOutcomeGroupGlobal) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowOutcomeGroupGlobal) GetBody() (string, error) {
	return "", nil
}

func (t *ShowOutcomeGroupGlobal) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowOutcomeGroupGlobal) Do(c *canvasapi.Canvas) (*models.OutcomeGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
