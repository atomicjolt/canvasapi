package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type GroupMembership struct {
	ID            int64  `json:"id" url:"id,omitempty"`                         // The id of the membership object.Example: 92
	GroupID       int64  `json:"group_id" url:"group_id,omitempty"`             // The id of the group object to which the membership belongs.Example: 17
	UserID        int64  `json:"user_id" url:"user_id,omitempty"`               // The id of the user object to which the membership belongs.Example: 3
	WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` // The current state of the membership. Current possible values are 'accepted', 'invited', and 'requested'.Example: accepted
	Moderator     bool   `json:"moderator" url:"moderator,omitempty"`           // Whether or not the user is a moderator of the group (the must also be an active member of the group to moderate).Example: true
	JustCreated   bool   `json:"just_created" url:"just_created,omitempty"`     // optional: whether or not the record was just created on a create call (POST), i.e. was the user just added to the group, or was the user already a member.Example: true
	SISImportID   int64  `json:"sis_import_id" url:"sis_import_id,omitempty"`   // The id of the SIS import if created through SIS. Only included if the user has permission to manage SIS information..Example: 4
}

func (t *GroupMembership) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"accepted", "invited", "requested"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		errs = append(errs, fmt.Sprintf("expected 'WorkflowState' to be one of %v", s))
	}
	return nil
}
