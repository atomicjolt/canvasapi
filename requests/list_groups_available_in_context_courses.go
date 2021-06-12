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

// ListGroupsAvailableInContextCourses Returns the paginated list of active groups in the given context that are visible to user.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # OnlyOwnGroups (Optional) Will only include groups that the user belongs to if this is set
// # Include (Optional) . Must be one of tabs- "tabs": Include the list of tabs configured for each group.  See the
//      {api:TabsController#index List available tabs API} for more information.
//
type ListGroupsAvailableInContextCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		OnlyOwnGroups bool     `json:"only_own_groups"` //  (Optional)
		Include       []string `json:"include"`         //  (Optional) . Must be one of tabs
	} `json:"query"`
}

func (t *ListGroupsAvailableInContextCourses) GetMethod() string {
	return "GET"
}

func (t *ListGroupsAvailableInContextCourses) GetURLPath() string {
	path := "courses/{course_id}/groups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListGroupsAvailableInContextCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListGroupsAvailableInContextCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListGroupsAvailableInContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"tabs"}, v) {
			errs = append(errs, "Include must be one of tabs")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListGroupsAvailableInContextCourses) Do(c *canvasapi.Canvas) ([]*models.Group, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Group{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
