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

// UpdateOutcomeGroupCourses Modify an existing outcome group. Fields not provided are left as is;
// unrecognized fields are ignored.
//
// When changing the parent outcome group, the new parent group must belong to
// the same context as this outcome group, and must not be a descendant of
// this outcome group (i.e. no cycles allowed).
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Title (Optional) The new outcome group title.
// # Description (Optional) The new outcome group description.
// # VendorGuid (Optional) A custom GUID for the learning standard.
// # ParentOutcomeGroupID (Optional) The id of the new parent outcome group.
//
type UpdateOutcomeGroupCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		Title                string `json:"title"`                   //  (Optional)
		Description          string `json:"description"`             //  (Optional)
		VendorGuid           string `json:"vendor_guid"`             //  (Optional)
		ParentOutcomeGroupID int64  `json:"parent_outcome_group_id"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateOutcomeGroupCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateOutcomeGroupCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_groups/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateOutcomeGroupCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateOutcomeGroupCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateOutcomeGroupCourses) HasErrors() error {
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

func (t *UpdateOutcomeGroupCourses) Do(c *canvasapi.Canvas) (*models.OutcomeGroup, error) {
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
