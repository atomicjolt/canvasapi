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

// UnlinkOutcomeCourses Unlinking an outcome only deletes the outcome itself if this was the last
// link to the outcome in any group in any context. Aligned outcomes cannot be
// deleted; as such, if this is the last link to an aligned outcome, the
// unlinking will fail.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
// # OutcomeID (Required) ID
//
type UnlinkOutcomeCourses struct {
	Path struct {
		CourseID  string `json:"course_id" url:"course_id,omitempty"`   //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
		OutcomeID string `json:"outcome_id" url:"outcome_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *UnlinkOutcomeCourses) GetMethod() string {
	return "DELETE"
}

func (t *UnlinkOutcomeCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_groups/{id}/outcomes/{outcome_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{outcome_id}", fmt.Sprintf("%v", t.Path.OutcomeID))
	return path
}

func (t *UnlinkOutcomeCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UnlinkOutcomeCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UnlinkOutcomeCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UnlinkOutcomeCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.OutcomeID == "" {
		errs = append(errs, "'OutcomeID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UnlinkOutcomeCourses) Do(c *canvasapi.Canvas) (*models.OutcomeLink, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeLink{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
