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

// ConcludeDeactivateOrDeleteEnrollment Conclude, deactivate, or delete an enrollment. If the +task+ argument isn't given, the enrollment
// will be concluded.
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Task (Optional) . Must be one of conclude, delete, inactivate, deactivateThe action to take on the enrollment.
//    When inactive, a user will still appear in the course roster to admins, but be unable to participate.
//    ("inactivate" and "deactivate" are equivalent tasks)
//
type ConcludeDeactivateOrDeleteEnrollment struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Task string `json:"task" url:"task,omitempty"` //  (Optional) . Must be one of conclude, delete, inactivate, deactivate
	} `json:"query"`
}

func (t *ConcludeDeactivateOrDeleteEnrollment) GetMethod() string {
	return "DELETE"
}

func (t *ConcludeDeactivateOrDeleteEnrollment) GetURLPath() string {
	path := "courses/{course_id}/enrollments/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ConcludeDeactivateOrDeleteEnrollment) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ConcludeDeactivateOrDeleteEnrollment) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ConcludeDeactivateOrDeleteEnrollment) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ConcludeDeactivateOrDeleteEnrollment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Query.Task != "" && !string_utils.Include([]string{"conclude", "delete", "inactivate", "deactivate"}, t.Query.Task) {
		errs = append(errs, "Task must be one of conclude, delete, inactivate, deactivate")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ConcludeDeactivateOrDeleteEnrollment) Do(c *canvasapi.Canvas) (*models.Enrollment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Enrollment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
