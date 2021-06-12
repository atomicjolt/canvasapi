package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type GroupCategory struct {
	ID                 int64     `json:"id"`                    // The ID of the group category..Example: 17
	Name               string    `json:"name"`                  // The display name of the group category..Example: Math Groups
	Role               string    `json:"role"`                  // Certain types of group categories have special role designations. Currently, these include: 'communities', 'student_organized', and 'imported'. Regular course/account group categories have a role of null..Example: communities
	SelfSignup         string    `json:"self_signup"`           // If the group category allows users to join a group themselves, thought they may only be a member of one group per group category at a time. Values include 'restricted', 'enabled', and null 'enabled' allows students to assign themselves to a group 'restricted' restricts them to only joining a group in their section null disallows students from joining groups.
	AutoLeader         string    `json:"auto_leader"`           // Gives instructors the ability to automatically have group leaders assigned.  Values include 'random', 'first', and null; 'random' picks a student from the group at random as the leader, 'first' sets the first student to be assigned to the group as the leader.
	ContextType        string    `json:"context_type"`          // The course or account that the category group belongs to. The pattern here is that whatever the context_type is, there will be an _id field named after that type. So if instead context_type was 'Course', the course_id field would be replaced by an course_id field..Example: Account
	AccountID          int64     `json:"account_id"`            // Example: 3
	GroupLimit         int64     `json:"group_limit"`           // If self-signup is enabled, group_limit can be set to cap the number of users in each group. If null, there is no limit..
	SISGroupCategoryID string    `json:"sis_group_category_id"` // The SIS identifier for the group category. This field is only included if the user has permission to manage or view SIS information..
	SISImportID        int64     `json:"sis_import_id"`         // The unique identifier for the SIS import. This field is only included if the user has permission to manage SIS information..
	Progress           *Progress `json:"progress"`              // If the group category has not yet finished a randomly student assignment request, a progress object will be attached, which will contain information related to the progress of the assignment request. Refer to the Progress API for more information.
}

func (t *GroupCategory) HasError() error {
	var s []string
	s = []string{"restricted", "enabled"}
	if !string_utils.Include(s, t.SelfSignup) {
		return fmt.Errorf("expected 'self_signup' to be one of %v", s)
	}
	s = []string{"first", "random"}
	if !string_utils.Include(s, t.AutoLeader) {
		return fmt.Errorf("expected 'auto_leader' to be one of %v", s)
	}
	return nil
}
