package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateTabForCourse Home and Settings tabs are not manageable, and can't be hidden or moved
//
// Returns a tab object
// https://canvas.instructure.com/doc/api/tabs.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TabID (Required) ID
//
// Form Parameters:
// # Position (Optional) The new position of the tab, 1-based
// # Hidden (Optional) no description
//
type UpdateTabForCourse struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		TabID    string `json:"tab_id"`    //  (Required)
	} `json:"path"`

	Form struct {
		Position int64 `json:"position"` //  (Optional)
		Hidden   bool  `json:"hidden"`   //  (Optional)
	} `json:"form"`
}

func (t *UpdateTabForCourse) GetMethod() string {
	return "PUT"
}

func (t *UpdateTabForCourse) GetURLPath() string {
	path := "courses/{course_id}/tabs/{tab_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{tab_id}", fmt.Sprintf("%v", t.Path.TabID))
	return path
}

func (t *UpdateTabForCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateTabForCourse) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateTabForCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TabID == "" {
		errs = append(errs, "'TabID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateTabForCourse) Do(c *canvasapi.Canvas) (*models.Tab, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Tab{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
