package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Course struct {
	ID                                int64            `json:"id" url:"id,omitempty"`                                                                       // the unique identifier for the course.Example: 370663
	SISCourseID                       string           `json:"sis_course_id" url:"sis_course_id,omitempty"`                                                 // the SIS identifier for the course, if defined. This field is only included if the user has permission to view SIS information..
	Uuid                              string           `json:"uuid" url:"uuid,omitempty"`                                                                   // the UUID of the course.Example: WvAHhY5FINzq5IyRIJybGeiXyFkG3SqHUPb7jZY5
	IntegrationID                     string           `json:"integration_id" url:"integration_id,omitempty"`                                               // the integration identifier for the course, if defined. This field is only included if the user has permission to view SIS information..
	SISImportID                       int64            `json:"sis_import_id" url:"sis_import_id,omitempty"`                                                 // the unique identifier for the SIS import. This field is only included if the user has permission to manage SIS information..Example: 34
	Name                              string           `json:"name" url:"name,omitempty"`                                                                   // the full name of the course.Example: InstructureCon 2012
	CourseCode                        string           `json:"course_code" url:"course_code,omitempty"`                                                     // the course code.Example: INSTCON12
	WorkflowState                     string           `json:"workflow_state" url:"workflow_state,omitempty"`                                               // the current state of the course one of 'unpublished', 'available', 'completed', or 'deleted'.Example: available
	AccountID                         int64            `json:"account_id" url:"account_id,omitempty"`                                                       // the account associated with the course.Example: 81259
	RootAccountID                     int64            `json:"root_account_id" url:"root_account_id,omitempty"`                                             // the root account associated with the course.Example: 81259
	EnrollmentTermID                  int64            `json:"enrollment_term_id" url:"enrollment_term_id,omitempty"`                                       // the enrollment term associated with the course.Example: 34
	GradingPeriods                    []*GradingPeriod `json:"grading_periods" url:"grading_periods,omitempty"`                                             // A list of grading periods associated with the course.
	GradingStandardID                 int64            `json:"grading_standard_id" url:"grading_standard_id,omitempty"`                                     // the grading standard associated with the course.Example: 25
	GradePassbackSetting              string           `json:"grade_passback_setting" url:"grade_passback_setting,omitempty"`                               // the grade_passback_setting set on the course.Example: nightly_sync
	CreatedAt                         time.Time        `json:"created_at" url:"created_at,omitempty"`                                                       // the date the course was created..Example: 2012-05-01T00:00:00-06:00
	StartAt                           time.Time        `json:"start_at" url:"start_at,omitempty"`                                                           // the start date for the course, if applicable.Example: 2012-06-01T00:00:00-06:00
	EndAt                             time.Time        `json:"end_at" url:"end_at,omitempty"`                                                               // the end date for the course, if applicable.Example: 2012-09-01T00:00:00-06:00
	Locale                            string           `json:"locale" url:"locale,omitempty"`                                                               // the course-set locale, if applicable.Example: en
	Enrollments                       []*Enrollment    `json:"enrollments" url:"enrollments,omitempty"`                                                     // A list of enrollments linking the current user to the course. for student enrollments, grading information may be included if include[]=total_scores.
	TotalStudents                     int64            `json:"total_students" url:"total_students,omitempty"`                                               // optional: the total number of active and invited students in the course.Example: 32
	Calendar                          *CalendarLink    `json:"calendar" url:"calendar,omitempty"`                                                           // course calendar.
	DefaultView                       string           `json:"default_view" url:"default_view,omitempty"`                                                   // the type of page that users will see when they first visit the course - 'feed': Recent Activity Dashboard - 'wiki': Wiki Front Page - 'modules': Course Modules/Sections Page - 'assignments': Course Assignments List - 'syllabus': Course Syllabus Page other types may be added in the future.Example: feed
	SyllabusBody                      string           `json:"syllabus_body" url:"syllabus_body,omitempty"`                                                 // optional: user-generated HTML for the course syllabus.Example: <p>syllabus html goes here</p>
	NeedsGradingCount                 int64            `json:"needs_grading_count" url:"needs_grading_count,omitempty"`                                     // optional: the number of submissions needing grading returned only if the current user has grading rights and include[]=needs_grading_count.Example: 17
	Term                              *Term            `json:"term" url:"term,omitempty"`                                                                   // optional: the enrollment term object for the course returned only if include[]=term.
	CourseProgress                    *CourseProgress  `json:"course_progress" url:"course_progress,omitempty"`                                             // optional: information on progress through the course returned only if include[]=course_progress.
	ApplyAssignmentGroupWeights       bool             `json:"apply_assignment_group_weights" url:"apply_assignment_group_weights,omitempty"`               // weight final grade based on assignment group percentages.Example: true
	Permissions                       string           `json:"permissions" url:"permissions,omitempty"`                                                     // optional: the permissions the user has for the course. returned only for a single course and include[]=permissions.Example: true, true
	IsPublic                          bool             `json:"is_public" url:"is_public,omitempty"`                                                         // Example: true
	IsPublicToAuthUsers               bool             `json:"is_public_to_auth_users" url:"is_public_to_auth_users,omitempty"`                             // Example: true
	PublicSyllabus                    bool             `json:"public_syllabus" url:"public_syllabus,omitempty"`                                             // Example: true
	PublicSyllabusToAuth              bool             `json:"public_syllabus_to_auth" url:"public_syllabus_to_auth,omitempty"`                             // Example: true
	PublicDescription                 string           `json:"public_description" url:"public_description,omitempty"`                                       // optional: the public description of the course.Example: Come one, come all to InstructureCon 2012!
	StorageQuotaMb                    int64            `json:"storage_quota_mb" url:"storage_quota_mb,omitempty"`                                           // Example: 5
	StorageQuotaUsedMb                float64          `json:"storage_quota_used_mb" url:"storage_quota_used_mb,omitempty"`                                 // Example: 5
	HideFinalGrades                   bool             `json:"hide_final_grades" url:"hide_final_grades,omitempty"`                                         //
	License                           string           `json:"license" url:"license,omitempty"`                                                             // Example: Creative Commons
	AllowStudentAssignmentEdits       bool             `json:"allow_student_assignment_edits" url:"allow_student_assignment_edits,omitempty"`               //
	AllowWikiComments                 bool             `json:"allow_wiki_comments" url:"allow_wiki_comments,omitempty"`                                     //
	AllowStudentForumAttachments      bool             `json:"allow_student_forum_attachments" url:"allow_student_forum_attachments,omitempty"`             //
	OpenEnrollment                    bool             `json:"open_enrollment" url:"open_enrollment,omitempty"`                                             // Example: true
	SelfEnrollment                    bool             `json:"self_enrollment" url:"self_enrollment,omitempty"`                                             //
	RestrictEnrollmentsToCourseDates  bool             `json:"restrict_enrollments_to_course_dates" url:"restrict_enrollments_to_course_dates,omitempty"`   //
	CourseFormat                      string           `json:"course_format" url:"course_format,omitempty"`                                                 // Example: online
	AccessRestrictedByDate            bool             `json:"access_restricted_by_date" url:"access_restricted_by_date,omitempty"`                         // optional: this will be true if this user is currently prevented from viewing the course because of date restriction settings.
	TimeZone                          string           `json:"time_zone" url:"time_zone,omitempty"`                                                         // The course's IANA time zone name..Example: America/Denver
	Blueprint                         bool             `json:"blueprint" url:"blueprint,omitempty"`                                                         // optional: whether the course is set as a Blueprint Course (blueprint fields require the Blueprint Courses feature).Example: true
	BlueprintRestrictions             string           `json:"blueprint_restrictions" url:"blueprint_restrictions,omitempty"`                               // optional: Set of restrictions applied to all locked course objects.Example: true, true, false, false
	BlueprintRestrictionsByObjectType string           `json:"blueprint_restrictions_by_object_type" url:"blueprint_restrictions_by_object_type,omitempty"` // optional: Sets of restrictions differentiated by object type applied to locked course objects.Example: {'content'=>true, 'points'=>true}, {'content'=>true}
	Template                          bool             `json:"template" url:"template,omitempty"`                                                           // optional: whether the course is set as a template (requires the Course Templates feature).Example: true
}

func (t *Course) HasError() error {
	var s []string
	s = []string{"unpublished", "available", "completed", "deleted"}
	if t.WorkflowState != "" && !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	s = []string{"feed", "wiki", "modules", "syllabus", "assignments"}
	if t.DefaultView != "" && !string_utils.Include(s, t.DefaultView) {
		return fmt.Errorf("expected 'default_view' to be one of %v", s)
	}
	return nil
}
