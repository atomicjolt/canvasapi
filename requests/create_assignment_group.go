package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateAssignmentGroup Create a new assignment group for this course.
// https://canvas.instructure.com/doc/api/assignment_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Name (Optional) The assignment group's name
// # Position (Optional) The position of this assignment group in relation to the other assignment groups
// # GroupWeight (Optional) The percent of the total grade that this assignment group represents
// # SISSourceID (Optional) The sis source id of the Assignment Group
// # IntegrationData (Optional) The integration data of the Assignment Group
// # Rules (Optional) The grading rules that are applied within this assignment group
//    See the Assignment Group object definition for format
//
type CreateAssignmentGroup struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name            string  `json:"name" url:"name,omitempty"`                         //  (Optional)
		Position        int64   `json:"position" url:"position,omitempty"`                 //  (Optional)
		GroupWeight     float64 `json:"group_weight" url:"group_weight,omitempty"`         //  (Optional)
		SISSourceID     string  `json:"sis_source_id" url:"sis_source_id,omitempty"`       //  (Optional)
		IntegrationData string  `json:"integration_data" url:"integration_data,omitempty"` //  (Optional)
		Rules           string  `json:"rules" url:"rules,omitempty"`                       //  (Optional)
	} `json:"form"`
}

func (t *CreateAssignmentGroup) GetMethod() string {
	return "POST"
}

func (t *CreateAssignmentGroup) GetURLPath() string {
	path := "courses/{course_id}/assignment_groups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateAssignmentGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateAssignmentGroup) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateAssignmentGroup) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateAssignmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateAssignmentGroup) Do(c *canvasapi.Canvas) (*models.AssignmentGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AssignmentGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
