package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowOutcome Returns the details of the outcome with the given id.
// https://canvas.instructure.com/doc/api/outcomes.html
//
// Path Parameters:
// # ID (Required) ID
//
type ShowOutcome struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *ShowOutcome) GetMethod() string {
	return "GET"
}

func (t *ShowOutcome) GetURLPath() string {
	path := "outcomes/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowOutcome) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowOutcome) GetBody() (string, error) {
	return "", nil
}

func (t *ShowOutcome) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowOutcome) Do(c *canvasapi.Canvas) (*models.Outcome, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Outcome{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
