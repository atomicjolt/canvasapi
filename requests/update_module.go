package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateModule Update and return an existing module
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Module (Optional) The name of the module
// # Module (Optional) The date the module will unlock
// # Module (Optional) The position of the module in the course (1-based)
// # Module (Optional) Whether module items must be unlocked in order
// # Module (Optional) IDs of Modules that must be completed before this one is unlocked
//    Prerequisite modules must precede this module (i.e. have a lower position
//    value), otherwise they will be ignored
// # Module (Optional) Whether to publish the student's final grade for the course upon
//    completion of this module.
// # Module (Optional) Whether the module is published and visible to students
//
type UpdateModule struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		Module struct {
			Name                      string    `json:"name" url:"name,omitempty"`                                               //  (Optional)
			UnlockAt                  time.Time `json:"unlock_at" url:"unlock_at,omitempty"`                                     //  (Optional)
			Position                  int64     `json:"position" url:"position,omitempty"`                                       //  (Optional)
			RequireSequentialProgress bool      `json:"require_sequential_progress" url:"require_sequential_progress,omitempty"` //  (Optional)
			PrerequisiteModuleIDs     []string  `json:"prerequisite_module_ids" url:"prerequisite_module_ids,omitempty"`         //  (Optional)
			PublishFinalGrade         bool      `json:"publish_final_grade" url:"publish_final_grade,omitempty"`                 //  (Optional)
			Published                 bool      `json:"published" url:"published,omitempty"`                                     //  (Optional)
		} `json:"module" url:"module,omitempty"`
	} `json:"form"`
}

func (t *UpdateModule) GetMethod() string {
	return "PUT"
}

func (t *UpdateModule) GetURLPath() string {
	path := "courses/{course_id}/modules/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateModule) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateModule) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateModule) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateModule) HasErrors() error {
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

func (t *UpdateModule) Do(c *canvasapi.Canvas) (*models.Module, error) {
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
