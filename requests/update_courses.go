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

// UpdateCourses Update multiple courses in an account.  Operates asynchronously; use the {api:ProgressController#show progress endpoint}
// to query the status of an operation.
//
// The action to take on each course.  Must be one of 'offer', 'conclude', 'delete', or 'undelete'.
//   * 'offer' makes a course visible to students. This action is also called "publish" on the web site.
//   * 'conclude' prevents future enrollments and makes a course read-only for all participants. The course still appears
//     in prior-enrollment lists.
//   * 'delete' completely removes the course from the web site (including course menus and prior-enrollment lists).
//     All enrollments are deleted. Course content may be physically deleted at a future date.
//   * 'undelete' attempts to recover a course that has been deleted. (Recovery is not guaranteed; please conclude
//     rather than delete a course if there is any possibility the course will be used again.) The recovered course
//     will be unpublished. Deleted enrollments will not be recovered.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.CourseIDs (Required) List of ids of courses to update. At most 500 courses may be updated in one call.
// # Form.Event (Required) . Must be one of offer, conclude, delete, undeleteno description
//
type UpdateCourses struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		CourseIDs []string `json:"course_ids" url:"course_ids,omitempty"` //  (Required)
		Event     string   `json:"event" url:"event,omitempty"`           //  (Required) . Must be one of offer, conclude, delete, undelete
	} `json:"form"`
}

func (t *UpdateCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateCourses) GetURLPath() string {
	path := "accounts/{account_id}/courses"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *UpdateCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateCourses) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Form.CourseIDs == nil {
		errs = append(errs, "'Form.CourseIDs' is required")
	}
	if t.Form.Event == "" {
		errs = append(errs, "'Form.Event' is required")
	}
	if t.Form.Event != "" && !string_utils.Include([]string{"offer", "conclude", "delete", "undelete"}, t.Form.Event) {
		errs = append(errs, "Event must be one of offer, conclude, delete, undelete")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCourses) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
