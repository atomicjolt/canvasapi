package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ReLockModuleProgressions Resets module progressions to their default locked state and
// recalculates them based on the current requirements.
//
// Adding progression requirements to an active course will not lock students
// out of modules they have already unlocked unless this action is called.
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type ReLockModuleProgressions struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *ReLockModuleProgressions) GetMethod() string {
	return "PUT"
}

func (t *ReLockModuleProgressions) GetURLPath() string {
	path := "courses/{course_id}/modules/{id}/relock"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ReLockModuleProgressions) GetQuery() (string, error) {
	return "", nil
}

func (t *ReLockModuleProgressions) GetBody() (string, error) {
	return "", nil
}

func (t *ReLockModuleProgressions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReLockModuleProgressions) Do(c *canvasapi.Canvas) (*models.Module, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Module{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
