package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetAllOutcomeGroupsForContextCourses
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
//
type GetAllOutcomeGroupsForContextCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`
}

func (t *GetAllOutcomeGroupsForContextCourses) GetMethod() string {
	return "GET"
}

func (t *GetAllOutcomeGroupsForContextCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_groups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetAllOutcomeGroupsForContextCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAllOutcomeGroupsForContextCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetAllOutcomeGroupsForContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllOutcomeGroupsForContextCourses) Do(c *canvasapi.Canvas) ([]*models.OutcomeGroup, error) {
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
