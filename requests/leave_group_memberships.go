package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// LeaveGroupMemberships Leave a group if you are allowed to leave (some groups, such as sets of
// course groups created by teachers, cannot be left). You may also use 'self'
// in place of a membership_id.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupID (Required) ID
// # MembershipID (Required) ID
//
type LeaveGroupMemberships struct {
	Path struct {
		GroupID      string `json:"group_id"`      //  (Required)
		MembershipID string `json:"membership_id"` //  (Required)
	} `json:"path"`
}

func (t *LeaveGroupMemberships) GetMethod() string {
	return "DELETE"
}

func (t *LeaveGroupMemberships) GetURLPath() string {
	path := "groups/{group_id}/memberships/{membership_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{membership_id}", fmt.Sprintf("%v", t.Path.MembershipID))
	return path
}

func (t *LeaveGroupMemberships) GetQuery() (string, error) {
	return "", nil
}

func (t *LeaveGroupMemberships) GetBody() (string, error) {
	return "", nil
}

func (t *LeaveGroupMemberships) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.MembershipID == "" {
		errs = append(errs, "'MembershipID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *LeaveGroupMemberships) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
