package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type Group struct {
	ID              int64                    `json:"id" url:"id,omitempty"`                               // The ID of the group..Example: 17
	Name            string                   `json:"name" url:"name,omitempty"`                           // The display name of the group..Example: Math Group 1
	Description     string                   `json:"description" url:"description,omitempty"`             // A description of the group. This is plain text..
	IsPublic        bool                     `json:"is_public" url:"is_public,omitempty"`                 // Whether or not the group is public.  Currently only community groups can be made public.  Also, once a group has been set to public, it cannot be changed back to private..
	FollowedByUser  bool                     `json:"followed_by_user" url:"followed_by_user,omitempty"`   // Whether or not the current user is following this group..
	JoinLevel       string                   `json:"join_level" url:"join_level,omitempty"`               // How people are allowed to join the group.  For all groups except for community groups, the user must share the group's parent course or account.  For student organized or community groups, where a user can be a member of as many or few as they want, the applicable levels are 'parent_context_auto_join', 'parent_context_request', and 'invitation_only'.  For class groups, where students are divided up and should only be part of one group of the category, this value will always be 'invitation_only', and is not relevant. * If 'parent_context_auto_join', anyone can join and will be automatically accepted. * If 'parent_context_request', anyone  can request to join, which must be approved by a group moderator. * If 'invitation_only', only those how have received an invitation my join the group, by accepting that invitation..Example: invitation_only
	MembersCount    int64                    `json:"members_count" url:"members_count,omitempty"`         // The number of members currently in the group.Example: 0
	AvatarUrl       string                   `json:"avatar_url" url:"avatar_url,omitempty"`               // The url of the group's avatar.Example: https://<canvas>/files/avatar_image.png
	ContextType     string                   `json:"context_type" url:"context_type,omitempty"`           // The course or account that the group belongs to. The pattern here is that whatever the context_type is, there will be an _id field named after that type. So if instead context_type was 'account', the course_id field would be replaced by an account_id field..Example: Course
	CourseID        int64                    `json:"course_id" url:"course_id,omitempty"`                 // Example: 3
	Role            string                   `json:"role" url:"role,omitempty"`                           // Certain types of groups have special role designations. Currently, these include: 'communities', 'student_organized', and 'imported'. Regular course/account groups have a role of null..
	GroupCategoryID int64                    `json:"group_category_id" url:"group_category_id,omitempty"` // The ID of the group's category..Example: 4
	SISGroupID      string                   `json:"sis_group_id" url:"sis_group_id,omitempty"`           // The SIS ID of the group. Only included if the user has permission to view SIS information..Example: group4a
	SISImportID     int64                    `json:"sis_import_id" url:"sis_import_id,omitempty"`         // The id of the SIS import if created through SIS. Only included if the user has permission to manage SIS information..Example: 14
	StorageQuotaMb  int64                    `json:"storage_quota_mb" url:"storage_quota_mb,omitempty"`   // the storage quota for the group, in megabytes.Example: 50
	Permissions     map[string](interface{}) `json:"permissions" url:"permissions,omitempty"`             // optional: the permissions the user has for the group. returned only for a single group and include[]=permissions.Example: true, true
	Users           []*User                  `json:"users" url:"users,omitempty"`                         // optional: A list of users that are members in the group. Returned only if include[]=users. WARNING: this collection's size is capped (if there are an extremely large number of users in the group (thousands) not all of them will be returned).  If you need to capture all the users in a group with certainty consider using the paginated /api/v1/groups/<group_id>/memberships endpoint..
}

func (t *Group) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"parent_context_auto_join", "parent_context_request", "invitation_only"}
	if t.JoinLevel != "" && !string_utils.Include(s, t.JoinLevel) {
		errs = append(errs, fmt.Sprintf("expected 'JoinLevel' to be one of %v", s))
	}
	s = []string{"communities", "student_organized", "imported"}
	if t.Role != "" && !string_utils.Include(s, t.Role) {
		errs = append(errs, fmt.Sprintf("expected 'Role' to be one of %v", s))
	}
	return nil
}
