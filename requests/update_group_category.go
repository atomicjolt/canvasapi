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

// UpdateGroupCategory Modifies an existing group category.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
// Form Parameters:
// # Name (Optional) Name of the group category
// # SelfSignup (Optional) . Must be one of enabled, restrictedAllow students to sign up for a group themselves (Course Only).
//    Valid values are:
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
type UpdateGroupCategory struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name               string `json:"name" url:"name,omitempty"`                                   //  (Optional)
		SelfSignup         string `json:"self_signup" url:"self_signup,omitempty"`                     //  (Optional) . Must be one of enabled, restricted
		AutoLeader         string `json:"auto_leader" url:"auto_leader,omitempty"`                     //  (Optional) . Must be one of first, random
		GroupLimit         int64  `json:"group_limit" url:"group_limit,omitempty"`                     //  (Optional)
		SISGroupCategoryID string `json:"sis_group_category_id" url:"sis_group_category_id,omitempty"` //  (Optional)
		CreateGroupCount   int64  `json:"create_group_count" url:"create_group_count,omitempty"`       //  (Optional)
		SplitGroupCount    string `json:"split_group_count" url:"split_group_count,omitempty"`         //  (Optional)
	} `json:"form"`
}

func (t *UpdateGroupCategory) GetMethod() string {
	return "PUT"
}

func (t *UpdateGroupCategory) GetURLPath() string {
	path := "group_categories/{group_category_id}"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *UpdateGroupCategory) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateGroupCategory) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateGroupCategory) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateGroupCategory) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if t.Form.SelfSignup != "" && !string_utils.Include([]string{"enabled", "restricted"}, t.Form.SelfSignup) {
		errs = append(errs, "SelfSignup must be one of enabled, restricted")
	}
	if t.Form.AutoLeader != "" && !string_utils.Include([]string{"first", "random"}, t.Form.AutoLeader) {
		errs = append(errs, "AutoLeader must be one of first, random")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateGroupCategory) Do(c *canvasapi.Canvas) (*models.GroupCategory, error) {
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
