package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListCourseSections A paginated list of the list of sections for this course.
// https://canvas.instructure.com/doc/api/sections.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of students, avatar_url, enrollments, total_students, passback_status- "students": Associations to include with the group. Note: this is only
//      available if you have permission to view users or grades in the course
//    - "avatar_url": Include the avatar URLs for students returned.
//    - "enrollments": If 'students' is also included, return the section
//      enrollment for each student
//    - "total_students": Returns the total amount of active and invited students
//      for the course section
//    - "passback_status": Include the grade passback status.
//
type ListCourseSections struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of students, avatar_url, enrollments, total_students, passback_status
	} `json:"query"`
}

func (t *ListCourseSections) GetMethod() string {
	return "GET"
}

func (t *ListCourseSections) GetURLPath() string {
	path := "courses/{course_id}/sections"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListCourseSections) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListCourseSections) GetBody() (string, error) {
	return "", nil
}

func (t *ListCourseSections) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"students", "avatar_url", "enrollments", "total_students", "passback_status"}, v) {
			errs = append(errs, "Include must be one of students, avatar_url, enrollments, total_students, passback_status")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCourseSections) Do(c *canvasapi.Canvas) ([]*models.Section, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Section{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
