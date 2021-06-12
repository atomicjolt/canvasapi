package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CoursesPermissions Returns permission information for the calling user in the given course.
// See also the {api:AccountsController#permissions Account} and
// {api:GroupsController#permissions Group} counterparts.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Permissions (Optional) List of permissions to check against the authenticated user.
//    Permission names are documented in the {api:RoleOverridesController#add_role Create a role} endpoint.
//
type CoursesPermissions struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Permissions []string `json:"permissions"` //  (Optional)
	} `json:"query"`
}

func (t *CoursesPermissions) GetMethod() string {
	return "GET"
}

func (t *CoursesPermissions) GetURLPath() string {
	path := "courses/{course_id}/permissions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CoursesPermissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *CoursesPermissions) GetBody() (string, error) {
	return "", nil
}

func (t *CoursesPermissions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CoursesPermissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
