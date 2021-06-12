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

// CreateModule Create and return a new module
// https://canvas.instructure.com/doc/api/modules.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Module (Required) The name of the module
// # Module (Optional) The date the module will unlock
// # Module (Optional) The position of this module in the course (1-based)
// # Module (Optional) Whether module items must be unlocked in order
// # Module (Optional) IDs of Modules that must be completed before this one is unlocked.
//    Prerequisite modules must precede this module (i.e. have a lower position
//    value), otherwise they will be ignored
// # Module (Optional) Whether to publish the student's final grade for the course upon
//    completion of this module.
//
type CreateModule struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Module struct {
			Name                      string    `json:"name"`                        //  (Required)
			UnlockAt                  time.Time `json:"unlock_at"`                   //  (Optional)
			Position                  int64     `json:"position"`                    //  (Optional)
			RequireSequentialProgress bool      `json:"require_sequential_progress"` //  (Optional)
			PrerequisiteModuleIDs     []string  `json:"prerequisite_module_ids"`     //  (Optional)
			PublishFinalGrade         bool      `json:"publish_final_grade"`         //  (Optional)
		} `json:"module"`
	} `json:"form"`
}

func (t *CreateModule) GetMethod() string {
	return "POST"
}

func (t *CreateModule) GetURLPath() string {
	path := "courses/{course_id}/modules"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateModule) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateModule) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateModule) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Module.Name == "" {
		errs = append(errs, "'Module' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateModule) Do(c *canvasapi.Canvas) (*models.Module, error) {
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
