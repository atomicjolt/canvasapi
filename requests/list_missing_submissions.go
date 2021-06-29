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

// ListMissingSubmissions A paginated list of past-due assignments for which the student does not have a submission.
// The user sending the request must either be the student, an admin or a parent observer using the parent app
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # Path.UserID (Required) the student's ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of planner_overrides, course"planner_overrides":: Optionally include the assignment's associated planner override, if it exists, for the current user.
//                          These will be returned under a +planner_override+ key
//    "course":: Optionally include the assignments' courses
// # Query.Filter (Optional) . Must be one of submittable"submittable":: Only return assignments that the current user can submit (i.e. filter out locked assignments)
// # Query.CourseIDs (Optional) Optionally restricts the list of past-due assignments to only those associated with the specified
//    course IDs.
//
type ListMissingSubmissions struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include   []string `json:"include" url:"include,omitempty"`       //  (Optional) . Must be one of planner_overrides, course
		Filter    []string `json:"filter" url:"filter,omitempty"`         //  (Optional) . Must be one of submittable
		CourseIDs []string `json:"course_ids" url:"course_ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListMissingSubmissions) GetMethod() string {
	return "GET"
}

func (t *ListMissingSubmissions) GetURLPath() string {
	path := "users/{user_id}/missing_submissions"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListMissingSubmissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListMissingSubmissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMissingSubmissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMissingSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"planner_overrides", "course"}, v) {
			errs = append(errs, "Include must be one of planner_overrides, course")
		}
	}
	for _, v := range t.Query.Filter {
		if v != "" && !string_utils.Include([]string{"submittable"}, v) {
			errs = append(errs, "Filter must be one of submittable")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMissingSubmissions) Do(c *canvasapi.Canvas) ([]*models.Assignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Assignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
