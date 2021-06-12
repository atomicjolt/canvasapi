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

// CreateGroupCategoryCourses Create a new group category
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Name (Required) Name of the group category
// # SelfSignup (Optional) . Must be one of enabled, restrictedAllow students to sign up for a group themselves (Course Only).
//    valid values are:
//    "enabled":: allows students to self sign up for any group in course
//    "restricted":: allows students to self sign up only for groups in the
//                   same section null disallows self sign up
// # AutoLeader (Optional) . Must be one of first, randomAssigns group leaders automatically when generating and allocating students to groups
//    Valid values are:
//    "first":: the first student to be allocated to a group is the leader
//    "random":: a random student from all members is chosen as the leader
// # GroupLimit (Optional) Limit the maximum number of users in each group (Course Only). Requires
//    self signup.
// # SISGroupCategoryID (Optional) The unique SIS identifier.
// # CreateGroupCount (Optional) Create this number of groups (Course Only).
// # SplitGroupCount (Optional) (Deprecated)
//    Create this number of groups, and evenly distribute students
//    among them. not allowed with "enable_self_signup". because
//    the group assignment happens synchronously, it's recommended
//    that you instead use the assign_unassigned_members endpoint.
//    (Course Only)
//
type CreateGroupCategoryCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Name               string `json:"name"`                  //  (Required)
		SelfSignup         string `json:"self_signup"`           //  (Optional) . Must be one of enabled, restricted
		AutoLeader         string `json:"auto_leader"`           //  (Optional) . Must be one of first, random
		GroupLimit         int64  `json:"group_limit"`           //  (Optional)
		SISGroupCategoryID string `json:"sis_group_category_id"` //  (Optional)
		CreateGroupCount   int64  `json:"create_group_count"`    //  (Optional)
		SplitGroupCount    string `json:"split_group_count"`     //  (Optional)
	} `json:"form"`
}

func (t *CreateGroupCategoryCourses) GetMethod() string {
	return "POST"
}

func (t *CreateGroupCategoryCourses) GetURLPath() string {
	path := "courses/{course_id}/group_categories"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateGroupCategoryCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateGroupCategoryCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateGroupCategoryCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Name == "" {
		errs = append(errs, "'Name' is required")
	}
	if !string_utils.Include([]string{"enabled", "restricted"}, t.Form.SelfSignup) {
		errs = append(errs, "SelfSignup must be one of enabled, restricted")
	}
	if !string_utils.Include([]string{"first", "random"}, t.Form.AutoLeader) {
		errs = append(errs, "AutoLeader must be one of first, random")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateGroupCategoryCourses) Do(c *canvasapi.Canvas) (*models.GroupCategory, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.GroupCategory{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
