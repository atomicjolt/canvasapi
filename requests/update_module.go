package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Module struct {
			Name                      string    `json:"name"`                        //  (Optional)
			UnlockAt                  time.Time `json:"unlock_at"`                   //  (Optional)
			Position                  int64     `json:"position"`                    //  (Optional)
			RequireSequentialProgress bool      `json:"require_sequential_progress"` //  (Optional)
			PrerequisiteModuleIDs     []string  `json:"prerequisite_module_ids"`     //  (Optional)
			PublishFinalGrade         bool      `json:"publish_final_grade"`         //  (Optional)
			Published                 bool      `json:"published"`                   //  (Optional)
		} `json:"module"`
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

func (t *UpdateModule) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
