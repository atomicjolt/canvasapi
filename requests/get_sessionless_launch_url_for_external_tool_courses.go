package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetSessionlessLaunchUrlForExternalToolCourses Returns a sessionless launch url for an external tool.
//
// NOTE: Either the id or url must be provided unless launch_type is assessment or module_item.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # ID (Optional) The external id of the tool to launch.
// # Url (Optional) The LTI launch url for the external tool.
// # AssignmentID (Optional) The assignment id for an assignment launch. Required if launch_type is set to "assessment".
// # ModuleItemID (Optional) The assignment id for a module item launch. Required if launch_type is set to "module_item".
// # LaunchType (Optional) . Must be one of assessment, module_itemThe type of launch to perform on the external tool. Placement names (eg. "course_navigation")
//    can also be specified to use the custom launch url for that placement; if done, the tool id
//    must be provided.
//
type GetSessionlessLaunchUrlForExternalToolCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		ID           string `json:"id"`             //  (Optional)
		Url          string `json:"url"`            //  (Optional)
		AssignmentID string `json:"assignment_id"`  //  (Optional)
		ModuleItemID string `json:"module_item_id"` //  (Optional)
		LaunchType   string `json:"launch_type"`    //  (Optional) . Must be one of assessment, module_item
	} `json:"query"`
}

func (t *GetSessionlessLaunchUrlForExternalToolCourses) GetMethod() string {
	return "GET"
}

func (t *GetSessionlessLaunchUrlForExternalToolCourses) GetURLPath() string {
	path := "courses/{course_id}/external_tools/sessionless_launch"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetSessionlessLaunchUrlForExternalToolCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSessionlessLaunchUrlForExternalToolCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetSessionlessLaunchUrlForExternalToolCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"assessment", "module_item"}, t.Query.LaunchType) {
		errs = append(errs, "LaunchType must be one of assessment, module_item")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSessionlessLaunchUrlForExternalToolCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
