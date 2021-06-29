package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type GroupCategory struct {
	ID                 int64     `json:"id" url:"id,omitempty"`                                       // The ID of the group category..Example: 17
	Name               string    `json:"name" url:"name,omitempty"`                                   // The display name of the group category..Example: Math Groups
	Role               string    `json:"role" url:"role,omitempty"`                                   // Certain types of group categories have special role designations. Currently, these include: 'communities', 'student_organized', and 'imported'. Regular course/account group categories have a role of null..Example: communities
	SelfSignup         string    `json:"self_signup" url:"self_signup,omitempty"`                     // If the group category allows users to join a group themselves, thought they may only be a member of one group per group category at a time. Values include 'restricted', 'enabled', and null 'enabled' allows students to assign themselves to a group 'restricted' restricts them to only joining a group in their section null disallows students from joining groups.
	AutoLeader         string    `json:"auto_leader" url:"auto_leader,omitempty"`                     // Gives instructors the ability to automatically have group leaders assigned.  Values include 'random', 'first', and null; 'random' picks a student from the group at random as the leader, 'first' sets the first student to be assigned to the group as the leader.
	ContextType        string    `json:"context_type" url:"context_type,omitempty"`                   // The course or account that the category group belongs to. The pattern here is that whatever the context_type is, there will be an _id field named after that type. So if instead context_type was 'Course', the course_id field would be replaced by an course_id field..Example: Account
	AccountID          int64     `json:"account_id" url:"account_id,omitempty"`                       // Example: 3
	GroupLimit         int64     `json:"group_limit" url:"group_limit,omitempty"`                     // If self-signup is enabled, group_limit can be set to cap the number of users in each group. If null, there is no limit..
	SISGroupCategoryID string    `json:"sis_group_category_id" url:"sis_group_category_id,omitempty"` // The SIS identifier for the group category. This field is only included if the user has permission to manage or view SIS information..
	SISImportID        int64     `json:"sis_import_id" url:"sis_import_id,omitempty"`                 // The unique identifier for the SIS import. This field is only included if the user has permission to manage SIS information..
	Progress           *Progress `json:"progress" url:"progress,omitempty"`                           // If the group category has not yet finished a randomly student assignment request, a progress object will be attached, which will contain information related to the progress of the assignment request. Refer to the Progress API for more information.
}

func (t *GroupCategory) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"restricted", "enabled"}
	if t.SelfSignup != "" && !string_utils.Include(s, t.SelfSignup) {
		errs = append(errs, fmt.Sprintf("expected 'SelfSignup' to be one of %v", s))
	}
	s = []string{"first", "random"}
	if t.AutoLeader != "" && !string_utils.Include(s, t.AutoLeader) {
		errs = append(errs, fmt.Sprintf("expected 'AutoLeader' to be one of %v", s))
	}
	return nil
}
