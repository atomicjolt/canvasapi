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

// CreateSubgroupCourses Creates a new empty subgroup under the outcome group with the given title
// and description.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Title (Required) The title of the new outcome group.
// # Description (Optional) The description of the new outcome group.
// # VendorGuid (Optional) A custom GUID for the learning standard
//
type CreateSubgroupCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Title       string `json:"title"`       //  (Required)
		Description string `json:"description"` //  (Optional)
		VendorGuid  string `json:"vendor_guid"` //  (Optional)
	} `json:"form"`
}

func (t *CreateSubgroupCourses) GetMethod() string {
	return "POST"
}

func (t *CreateSubgroupCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_groups/{id}/subgroups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *CreateSubgroupCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSubgroupCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateSubgroupCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Title == "" {
		errs = append(errs, "'Title' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSubgroupCourses) Do(c *canvasapi.Canvas) (*models.OutcomeGroup, error) {
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
