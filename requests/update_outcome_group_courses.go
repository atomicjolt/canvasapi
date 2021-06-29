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

// UpdateOutcomeGroupCourses Modify an existing outcome group. Fields not provided are left as is;
// unrecognized fields are ignored.
//
// When changing the parent outcome group, the new parent group must belong to
// the same context as this outcome group, and must not be a descendant of
// this outcome group (i.e. no cycles allowed).
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Title (Optional) The new outcome group title.
// # Form.Description (Optional) The new outcome group description.
// # Form.VendorGuid (Optional) A custom GUID for the learning standard.
// # Form.ParentOutcomeGroupID (Optional) The id of the new parent outcome group.
//
type UpdateOutcomeGroupCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		Title                string `json:"title" url:"title,omitempty"`                                     //  (Optional)
		Description          string `json:"description" url:"description,omitempty"`                         //  (Optional)
		VendorGuid           string `json:"vendor_guid" url:"vendor_guid,omitempty"`                         //  (Optional)
		ParentOutcomeGroupID int64  `json:"parent_outcome_group_id" url:"parent_outcome_group_id,omitempty"` //  (Optional)
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

func (t *UpdateOutcomeGroupCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateOutcomeGroupCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateOutcomeGroupCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
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
