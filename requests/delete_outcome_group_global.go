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

// DeleteOutcomeGroupGlobal Deleting an outcome group deletes descendant outcome groups and outcome
// links. The linked outcomes themselves are only deleted if all links to the
// outcome were deleted.
//
// Aligned outcomes cannot be deleted; as such, if all remaining links to an
// aligned outcome are included in this group's descendants, the group
// deletion will fail.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type DeleteOutcomeGroupGlobal struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteOutcomeGroupGlobal) GetMethod() string {
	return "DELETE"
}

func (t *DeleteOutcomeGroupGlobal) GetURLPath() string {
	path := "global/outcome_groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteOutcomeGroupGlobal) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteOutcomeGroupGlobal) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteOutcomeGroupGlobal) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteOutcomeGroupGlobal) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteOutcomeGroupGlobal) Do(c *canvasapi.Canvas) (*models.OutcomeGroup, error) {
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
