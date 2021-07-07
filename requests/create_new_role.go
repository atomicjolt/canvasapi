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

// CreateNewRole Create a new course-level or account-level role.
// https://canvas.instructure.com/doc/api/roles.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.Label (Required) Label for the role.
// # Form.Role (Optional) Deprecated alias for label.
// # Form.BaseRoleType (Optional) . Must be one of AccountMembership, StudentEnrollment, TeacherEnrollment, TaEnrollment, ObserverEnrollment, DesignerEnrollmentSpecifies the role type that will be used as a base
//    for the permissions granted to this role.
//
//    Defaults to 'AccountMembership' if absent
// # Form.Permissions (Optional) no description
// # Form.Permissions (Optional) If explicit is 1 and enabled is 1, permission <X> will be explicitly
//    granted to this role. If explicit is 1 and enabled has any other value
//    (typically 0), permission <X> will be explicitly denied to this role. If
//    explicit is any other value (typically 0) or absent, or if enabled is
//    absent, the value for permission <X> will be inherited from upstream.
//    Ignored if permission <X> is locked upstream (in an ancestor account).
//
//    May occur multiple times with unique values for <X>. Recognized
//    permission names for <X> are:
//
//      [For Account-Level Roles Only]
//      become_user                      -- Users - act as
//      import_sis                       -- SIS Data - import
//      manage_account_memberships       -- Admins - add / remove
//      manage_account_settings          -- Account-level settings - manage
//      manage_alerts                    -- Global announcements - add / edit / delete
//      manage_catalog                   -- Catalog - manage
//      Manage Course Templates granular permissions
//          add_course_template          -- Course Templates - add
//          delete_course_template       -- Course Templates - delete
//          edit_course_template         -- Course Templates - edit
//      Manage Courses granular permissions
//          manage_courses_admin         -- Manage Courses - manage / update
//      manage_developer_keys            -- Developer keys - manage
//      manage_feature_flags             -- Feature Options - enable / disable
//      manage_master_courses            -- Blueprint Courses - add / edit / associate / delete
//      manage_role_overrides            -- Permissions - manage
//      manage_storage_quotas            -- Storage Quotas - manage
//      manage_sis                       -- SIS data - manage
//      manage_user_logins               -- Users - manage login details
//      manage_user_observers            -- Users - manage observers
//      moderate_user_content            -- Users - moderate content
//      read_course_content              -- Course Content - view
//      read_course_list                 -- Courses - view list
//      view_course_changes              -- Courses - view change logs
//      view_feature_flags               -- Feature Options - view
//      view_grade_changes               -- Grades - view change logs
//      view_notifications               -- Notifications - view
//      view_quiz_answer_audits          -- Quizzes - view submission log
//      view_statistics                  -- Statistics - view
//      undelete_courses                 -- Courses - undelete
//
//      [For both Account-Level and Course-Level roles]
//       Note: Applicable enrollment types for course-level roles are given in brackets:
//             S = student, T = teacher (instructor), A = TA, D = designer, O = observer.
//             Lower-case letters indicate permissions that are off by default.
//             A missing letter indicates the permission cannot be enabled for the role
//             or any derived custom roles.
//      allow_course_admin_actions       -- [ Tad ] Users - allow administrative actions in courses
//      create_collaborations            -- [STADo] Student Collaborations - create
//      create_conferences               -- [STADo] Web conferences - create
//      create_forum                     -- [STADo] Discussions - create
//      generate_observer_pairing_code   -- [ tado] Users - Generate observer pairing codes for students
//      import_outcomes                  -- [ TaDo] Learning Outcomes - import
//      lti_add_edit                     -- [ TAD ] LTI - add / edit / delete
//      manage_assignments               -- [ TADo] Assignments and Quizzes - add / edit / delete
//      manage_calendar                  -- [sTADo] Course Calendar - add / edit / delete
//      manage_content                   -- [ TADo] Course Content - add / edit / delete
//      manage_course_visibility         -- [ TAD ] Course - change visibility
//      Manage Courses granular permissions
//          manage_courses_add           -- [sTADo] Courses - add
//          manage_courses_conclude      -- [ TaD ] Courses - conclude
//          manage_courses_delete        -- [ TaD ] Courses - delete
//          manage_courses_publish       -- [ TaD ] Courses - publish
//          manage_courses_reset         -- [ TaD ] Courses - reset
//      Manage Files granular permissions
//          manage_files_add             -- [ TADo] Course Files - add
//          manage_files_edit            -- [ TADo] Course Files - edit
//          manage_files_delete          -- [ TADo] Course Files - delete
//      manage_grades                    -- [ TA  ] Grades - edit
//      manage_groups                    -- [ TAD ] Groups - add / edit / delete
//      manage_interaction_alerts        -- [ Ta  ] Alerts - add / edit / delete
//      manage_outcomes                  -- [sTaDo] Learning Outcomes - add / edit / delete
//      manage_proficiency_calculations  -- [ t d ] Outcome Proficiency Calculations - add / edit / delete
//      manage_proficiency_scales        -- [ t d ] Outcome Proficiency/Mastery Scales - add / edit / delete
//      Manage Sections granular permissions
//          manage_sections_add          -- [ TaD ] Course Sections - add
//          manage_sections_edit         -- [ TaD ] Course Sections - edit
//          manage_sections_delete       -- [ TaD ] Course Sections - delete
//      manage_students                  -- [ TAD ] Users - manage students in courses
//      manage_user_notes                -- [ TA  ] Faculty Journal - manage entries
//      manage_rubrics                   -- [ TAD ] Rubrics - add / edit / delete
//      Manage Pages granular permissions
//          manage_wiki_create           -- [ TADo] Pages - create
//          manage_wiki_delete           -- [ TADo] Pages - delete
//          manage_wiki_update           -- [ TADo] Pages - update
//      moderate_forum                   -- [sTADo] Discussions - moderate
//      post_to_forum                    -- [STADo] Discussions - post
//      read_announcements               -- [STADO] Announcements - view
//      read_email_addresses             -- [sTAdo] Users - view primary email address
//      read_forum                       -- [STADO] Discussions - view
//      read_question_banks              -- [ TADo] Question banks - view and link
//      read_reports                     -- [ TAD ] Courses - view usage reports
//      read_roster                      -- [STADo] Users - view list
//      read_sis                         -- [sTa  ] SIS Data - read
//      select_final_grade               -- [ TA  ] Grades - select final grade for moderation
//      send_messages                    -- [STADo] Conversations - send messages to individual course members
//      send_messages_all                -- [sTADo] Conversations - send messages to entire class
//      Users - Teacher granular permissions
//          add_teacher_to_course        -- [ Tad ] Add a teacher enrollment to a course
//          remove_teacher_from_course   -- [ Tad ] Remove a Teacher enrollment from a course
//      Users - TA granular permissions
//          add_ta_to_course             -- [ Tad ] Add a TA enrollment to a course
//          remove_ta_from_course        -- [ Tad ] Remove a TA enrollment from a course
//      Users - Designer granular permissions
//          add_designer_to_course       -- [ Tad ] Add a designer enrollment to a course
//          remove_designer_from_course  -- [ Tad ] Remove a designer enrollment from a course
//      Users - Observer granular permissions
//          add_observer_to_course       -- [ Tad ] Add an observer enrollment to a course
//          remove_observer_from_course  -- [ Tad ] Remove an observer enrollment from a course
//      Users - Student granular permissions
//          add_student_to_course        -- [ Tad ] Add a student enrollment to a course
//          remove_student_from_course   -- [ Tad ] Remove a student enrollment from a course
//      view_all_grades                  -- [ TAd ] Grades - view all grades
//      view_analytics                   -- [sTA  ] Analytics - view pages
//      view_audit_trail                 -- [ t   ] Grades - view audit trail
//      view_group_pages                 -- [sTADo] Groups - view all student groups
//      view_user_logins                 -- [ TA  ] Users - view login IDs
//
//    Some of these permissions are applicable only for roles on the site admin
//    account, on a root account, or for course-level roles with a particular base role type;
//    if a specified permission is inapplicable, it will be ignored.
//
//    Additional permissions may exist based on installed plugins.
//
//    A comprehensive list of all permissions are available:
//
//    Course Permissions PDF: http://bit.ly/cnvs-course-permissions
//
//    Account Permissions PDF: http://bit.ly/cnvs-acct-permissions
// # Form.Permissions (Optional) If the value is 1, permission <X> will be locked downstream (new roles in
//    subaccounts cannot override the setting). For any other value, permission
//    <X> is left unlocked. Ignored if permission <X> is already locked
//    upstream. May occur multiple times with unique values for <X>.
// # Form.Permissions (Optional) If the value is 1, permission <X> applies to the account this role is in.
//    The default value is 1. Must be true if applies_to_descendants is false.
//    This value is only returned if enabled is true.
// # Form.Permissions (Optional) If the value is 1, permission <X> cascades down to sub accounts of the
//    account this role is in. The default value is 1.  Must be true if
//    applies_to_self is false.This value is only returned if enabled is true.
//
type CreateNewRole struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Label        string `json:"label" url:"label,omitempty"`                   //  (Required)
		Role         string `json:"role" url:"role,omitempty"`                     //  (Optional)
		BaseRoleType string `json:"base_role_type" url:"base_role_type,omitempty"` //  (Optional) . Must be one of AccountMembership, StudentEnrollment, TeacherEnrollment, TaEnrollment, ObserverEnrollment, DesignerEnrollment
		Permissions  map[string]CreateNewRolePermissions
	} `json:"form"`
}

func (t *CreateNewRole) GetMethod() string {
	return "POST"
}

func (t *CreateNewRole) GetURLPath() string {
	path := "accounts/{account_id}/roles"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateNewRole) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateNewRole) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateNewRole) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateNewRole) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Form.Label == "" {
		errs = append(errs, "'Form.Label' is required")
	}
	if t.Form.BaseRoleType != "" && !string_utils.Include([]string{"AccountMembership", "StudentEnrollment", "TeacherEnrollment", "TaEnrollment", "ObserverEnrollment", "DesignerEnrollment"}, t.Form.BaseRoleType) {
		errs = append(errs, "BaseRoleType must be one of AccountMembership, StudentEnrollment, TeacherEnrollment, TaEnrollment, ObserverEnrollment, DesignerEnrollment")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateNewRole) Do(c *canvasapi.Canvas) (*models.Role, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Role{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type CreateNewRolePermissions struct {
	Explicit             bool `json:"explicit" url:"explicit,omitempty"`                             //  (Optional)
	Enabled              bool `json:"enabled" url:"enabled,omitempty"`                               //  (Optional)
	Locked               bool `json:"locked" url:"locked,omitempty"`                                 //  (Optional)
	AppliesToSelf        bool `json:"applies_to_self" url:"applies_to_self,omitempty"`               //  (Optional)
	AppliesToDescendants bool `json:"applies_to_descendants" url:"applies_to_descendants,omitempty"` //  (Optional)
}
