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
	"github.com/atomicjolt/string_utils"
)

// GetSectionInformationCourses Gets details about a specific section
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of students, avatar_url, enrollments, total_students, passback_status- "students": Associations to include with the group. Note: this is only
//      available if you have permission to view users or grades in the course
//    - "avatar_url": Include the avatar URLs for students returned.
//    - "enrollments": If 'students' is also included, return the section
//      enrollment for each student
//    - "total_students": Returns the total amount of active and invited students
//      for the course section
//    - "passback_status": Include the grade passback status.
//
type GetSectionInformationCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of students, avatar_url, enrollments, total_students, passback_status
	} `json:"query"`
}

func (t *GetSectionInformationCourses) GetMethod() string {
	return "GET"
}

func (t *GetSectionInformationCourses) GetURLPath() string {
	path := "courses/{course_id}/sections/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSectionInformationCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetSectionInformationCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSectionInformationCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSectionInformationCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"students", "avatar_url", "enrollments", "total_students", "passback_status"}, v) {
			errs = append(errs, "Include must be one of students, avatar_url, enrollments, total_students, passback_status")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSectionInformationCourses) Do(c *canvasapi.Canvas) (*models.Section, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Section{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
