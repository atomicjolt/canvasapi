package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// UpdateCourse Update an existing course.
//
// Arguments are the same as Courses#create, with a few exceptions (enroll_me).
//
// If a user has content management rights, but not full course editing rights, the only attribute
// editable through this endpoint will be "syllabus_body"
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Course.AccountID (Optional) The unique ID of the account to move the course to.
// # Form.Course.Name (Optional) The name of the course. If omitted, the course will be named "Unnamed
//    Course."
// # Form.Course.CourseCode (Optional) The course code for the course.
// # Form.Course.StartAt (Optional) Course start date in ISO8601 format, e.g. 2011-01-01T01:00Z
// # Form.Course.EndAt (Optional) Course end date in ISO8601 format. e.g. 2011-01-01T01:00Z
// # Form.Course.License (Optional) The name of the licensing. Should be one of the following abbreviations
//    (a descriptive name is included in parenthesis for reference):
//    - 'private' (Private Copyrighted)
//    - 'cc_by_nc_nd' (CC Attribution Non-Commercial No Derivatives)
//    - 'cc_by_nc_sa' (CC Attribution Non-Commercial Share Alike)
//    - 'cc_by_nc' (CC Attribution Non-Commercial)
//    - 'cc_by_nd' (CC Attribution No Derivatives)
//    - 'cc_by_sa' (CC Attribution Share Alike)
//    - 'cc_by' (CC Attribution)
//    - 'public_domain' (Public Domain).
// # Form.Course.IsPublic (Optional) Set to true if course is public to both authenticated and unauthenticated users.
// # Form.Course.IsPublicToAuthUsers (Optional) Set to true if course is public only to authenticated users.
// # Form.Course.PublicSyllabus (Optional) Set to true to make the course syllabus public.
// # Form.Course.PublicSyllabusToAuth (Optional) Set to true to make the course syllabus to public for authenticated users.
// # Form.Course.PublicDescription (Optional) A publicly visible description of the course.
// # Form.Course.AllowStudentWikiEdits (Optional) If true, students will be able to modify the course wiki.
// # Form.Course.AllowWikiComments (Optional) If true, course members will be able to comment on wiki pages.
// # Form.Course.AllowStudentForumAttachments (Optional) If true, students can attach files to forum posts.
// # Form.Course.OpenEnrollment (Optional) Set to true if the course is open enrollment.
// # Form.Course.SelfEnrollment (Optional) Set to true if the course is self enrollment.
// # Form.Course.RestrictEnrollmentsToCourseDates (Optional) Set to true to restrict user enrollments to the start and end dates of the
//    course. This parameter is required when using the API, as this option is
//    not displayed in the Course Settings page.
// # Form.Course.TermID (Optional) The unique ID of the term to create to course in.
// # Form.Course.SISCourseID (Optional) The unique SIS identifier.
// # Form.Course.IntegrationID (Optional) The unique Integration identifier.
// # Form.Course.HideFinalGrades (Optional) If this option is set to true, the totals in student grades summary will
//    be hidden.
// # Form.Course.TimeZone (Optional) The time zone for the course. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # Form.Course.ApplyAssignmentGroupWeights (Optional) Set to true to weight final grade based on assignment groups percentages.
// # Form.Course.StorageQuotaMb (Optional) Set the storage quota for the course, in megabytes. The caller must have
//    the "Manage storage quotas" account permission.
// # Form.Offer (Optional) If this option is set to true, the course will be available to students
//    immediately.
// # Form.Course.Event (Optional) . Must be one of claim, offer, conclude, delete, undeleteThe action to take on each course.
//    * 'claim' makes a course no longer visible to students. This action is also called "unpublish" on the web site.
//      A course cannot be unpublished if students have received graded submissions.
//    * 'offer' makes a course visible to students. This action is also called "publish" on the web site.
//    * 'conclude' prevents future enrollments and makes a course read-only for all participants. The course still appears
//      in prior-enrollment lists.
//    * 'delete' completely removes the course from the web site (including course menus and prior-enrollment lists).
//      All enrollments are deleted. Course content may be physically deleted at a future date.
//    * 'undelete' attempts to recover a course that has been deleted. This action requires account administrative rights.
//      (Recovery is not guaranteed; please conclude rather than delete a course if there is any possibility the course
//      will be used again.) The recovered course will be unpublished. Deleted enrollments will not be recovered.
// # Form.Course.DefaultView (Optional) . Must be one of feed, wiki, modules, syllabus, assignmentsThe type of page that users will see when they first visit the course
//    * 'feed' Recent Activity Dashboard
//    * 'wiki' Wiki Front Page
//    * 'modules' Course Modules/Sections Page
//    * 'assignments' Course Assignments List
//    * 'syllabus' Course Syllabus Page
//    other types may be added in the future
// # Form.Course.SyllabusBody (Optional) The syllabus body for the course
// # Form.Course.SyllabusCourseSummary (Optional) Optional. Indicates whether the Course Summary (consisting of the course's assignments and calendar events) is displayed on the syllabus page. Defaults to +true+.
// # Form.Course.GradingStandardID (Optional) The grading standard id to set for the course.  If no value is provided for this argument the current grading_standard will be un-set from this course.
// # Form.Course.GradePassbackSetting (Optional) Optional. The grade_passback_setting for the course. Only 'nightly_sync' and '' are allowed
// # Form.Course.CourseFormat (Optional) Optional. Specifies the format of the course. (Should be either 'on_campus' or 'online')
// # Form.Course.ImageID (Optional) This is a file ID corresponding to an image file in the course that will
//    be used as the course image.
//    This will clear the course's image_url setting if set.  If you attempt
//    to provide image_url and image_id in a request it will fail.
// # Form.Course.ImageUrl (Optional) This is a URL to an image to be used as the course image.
//    This will clear the course's image_id setting if set.  If you attempt
//    to provide image_url and image_id in a request it will fail.
// # Form.Course.RemoveImage (Optional) If this option is set to true, the course image url and course image
//    ID are both set to nil
// # Form.Course.Blueprint (Optional) Sets the course as a blueprint course.
// # Form.Course.BlueprintRestrictions (Optional) Sets a default set to apply to blueprint course objects when restricted,
//    unless _use_blueprint_restrictions_by_object_type_ is enabled.
//    See the {api:Blueprint_Courses:BlueprintRestriction Blueprint Restriction} documentation
// # Form.Course.UseBlueprintRestrictionsByObjectType (Optional) When enabled, the _blueprint_restrictions_ parameter will be ignored in favor of
//    the _blueprint_restrictions_by_object_type_ parameter
// # Form.Course.BlueprintRestrictionsByObjectType (Optional) Allows setting multiple {api:Blueprint_Courses:BlueprintRestriction Blueprint Restriction}
//    to apply to blueprint course objects of the matching type when restricted.
//    The possible object types are "assignment", "attachment", "discussion_topic", "quiz" and "wiki_page".
//    Example usage:
//      course[blueprint_restrictions_by_object_type][assignment][content]=1
// # Form.Course.HomeroomCourse (Optional) Sets the course as a homeroom course. The setting takes effect only when the course is associated
//    with a Canvas for Elementary-enabled account.
// # Form.Course.SyncEnrollmentsFromHomeroom (Optional) Syncs enrollments from the homeroom that is set in homeroom_course_id. The setting only takes effect when the
//    course is associated with a Canvas for Elementary-enabled account and sync_enrollments_from_homeroom is enabled.
// # Form.Course.HomeroomCourseID (Optional) Sets the Homeroom Course id to be used with sync_enrollments_from_homeroom. The setting only takes effect when the
//    course is associated with a Canvas for Elementary-enabled account and sync_enrollments_from_homeroom is enabled.
// # Form.Course.Template (Optional) Enable or disable the course as a template that can be selected by an account
// # Form.Course.CourseColor (Optional) Sets a color in hex code format to be associated with the course. The setting takes effect only when the course
//    is associated with a Canvas for Elementary-enabled account.
//
type UpdateCourse struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Course struct {
			AccountID                            int64                          `json:"account_id" url:"account_id,omitempty"`                                                               //  (Optional)
			Name                                 string                         `json:"name" url:"name,omitempty"`                                                                           //  (Optional)
			CourseCode                           string                         `json:"course_code" url:"course_code,omitempty"`                                                             //  (Optional)
			StartAt                              time.Time                      `json:"start_at" url:"start_at,omitempty"`                                                                   //  (Optional)
			EndAt                                time.Time                      `json:"end_at" url:"end_at,omitempty"`                                                                       //  (Optional)
			License                              string                         `json:"license" url:"license,omitempty"`                                                                     //  (Optional)
			IsPublic                             bool                           `json:"is_public" url:"is_public,omitempty"`                                                                 //  (Optional)
			IsPublicToAuthUsers                  bool                           `json:"is_public_to_auth_users" url:"is_public_to_auth_users,omitempty"`                                     //  (Optional)
			PublicSyllabus                       bool                           `json:"public_syllabus" url:"public_syllabus,omitempty"`                                                     //  (Optional)
			PublicSyllabusToAuth                 bool                           `json:"public_syllabus_to_auth" url:"public_syllabus_to_auth,omitempty"`                                     //  (Optional)
			PublicDescription                    string                         `json:"public_description" url:"public_description,omitempty"`                                               //  (Optional)
			AllowStudentWikiEdits                bool                           `json:"allow_student_wiki_edits" url:"allow_student_wiki_edits,omitempty"`                                   //  (Optional)
			AllowWikiComments                    bool                           `json:"allow_wiki_comments" url:"allow_wiki_comments,omitempty"`                                             //  (Optional)
			AllowStudentForumAttachments         bool                           `json:"allow_student_forum_attachments" url:"allow_student_forum_attachments,omitempty"`                     //  (Optional)
			OpenEnrollment                       bool                           `json:"open_enrollment" url:"open_enrollment,omitempty"`                                                     //  (Optional)
			SelfEnrollment                       bool                           `json:"self_enrollment" url:"self_enrollment,omitempty"`                                                     //  (Optional)
			RestrictEnrollmentsToCourseDates     bool                           `json:"restrict_enrollments_to_course_dates" url:"restrict_enrollments_to_course_dates,omitempty"`           //  (Optional)
			TermID                               int64                          `json:"term_id" url:"term_id,omitempty"`                                                                     //  (Optional)
			SISCourseID                          string                         `json:"sis_course_id" url:"sis_course_id,omitempty"`                                                         //  (Optional)
			IntegrationID                        string                         `json:"integration_id" url:"integration_id,omitempty"`                                                       //  (Optional)
			HideFinalGrades                      bool                           `json:"hide_final_grades" url:"hide_final_grades,omitempty"`                                                 //  (Optional)
			TimeZone                             string                         `json:"time_zone" url:"time_zone,omitempty"`                                                                 //  (Optional)
			ApplyAssignmentGroupWeights          bool                           `json:"apply_assignment_group_weights" url:"apply_assignment_group_weights,omitempty"`                       //  (Optional)
			StorageQuotaMb                       int64                          `json:"storage_quota_mb" url:"storage_quota_mb,omitempty"`                                                   //  (Optional)
			Event                                string                         `json:"event" url:"event,omitempty"`                                                                         //  (Optional) . Must be one of claim, offer, conclude, delete, undelete
			DefaultView                          string                         `json:"default_view" url:"default_view,omitempty"`                                                           //  (Optional) . Must be one of feed, wiki, modules, syllabus, assignments
			SyllabusBody                         string                         `json:"syllabus_body" url:"syllabus_body,omitempty"`                                                         //  (Optional)
			SyllabusCourseSummary                bool                           `json:"syllabus_course_summary" url:"syllabus_course_summary,omitempty"`                                     //  (Optional)
			GradingStandardID                    int64                          `json:"grading_standard_id" url:"grading_standard_id,omitempty"`                                             //  (Optional)
			GradePassbackSetting                 string                         `json:"grade_passback_setting" url:"grade_passback_setting,omitempty"`                                       //  (Optional)
			CourseFormat                         string                         `json:"course_format" url:"course_format,omitempty"`                                                         //  (Optional)
			ImageID                              int64                          `json:"image_id" url:"image_id,omitempty"`                                                                   //  (Optional)
			ImageUrl                             string                         `json:"image_url" url:"image_url,omitempty"`                                                                 //  (Optional)
			RemoveImage                          bool                           `json:"remove_image" url:"remove_image,omitempty"`                                                           //  (Optional)
			Blueprint                            bool                           `json:"blueprint" url:"blueprint,omitempty"`                                                                 //  (Optional)
			BlueprintRestrictions                *models.BlueprintRestriction   `json:"blueprint_restrictions" url:"blueprint_restrictions,omitempty"`                                       //  (Optional)
			UseBlueprintRestrictionsByObjectType bool                           `json:"use_blueprint_restrictions_by_object_type" url:"use_blueprint_restrictions_by_object_type,omitempty"` //  (Optional)
			BlueprintRestrictionsByObjectType    []*models.BlueprintRestriction `json:"blueprint_restrictions_by_object_type" url:"blueprint_restrictions_by_object_type,omitempty"`         //  (Optional)
			HomeroomCourse                       bool                           `json:"homeroom_course" url:"homeroom_course,omitempty"`                                                     //  (Optional)
			SyncEnrollmentsFromHomeroom          string                         `json:"sync_enrollments_from_homeroom" url:"sync_enrollments_from_homeroom,omitempty"`                       //  (Optional)
			HomeroomCourseID                     string                         `json:"homeroom_course_id" url:"homeroom_course_id,omitempty"`                                               //  (Optional)
			Template                             bool                           `json:"template" url:"template,omitempty"`                                                                   //  (Optional)
			CourseColor                          string                         `json:"course_color" url:"course_color,omitempty"`                                                           //  (Optional)
		} `json:"course" url:"course,omitempty"`

		Offer bool `json:"offer" url:"offer,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateCourse) GetMethod() string {
	return "PUT"
}

func (t *UpdateCourse) GetURLPath() string {
	path := "courses/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCourse) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateCourse) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateCourse) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.Course.Event != "" && !string_utils.Include([]string{"claim", "offer", "conclude", "delete", "undelete"}, t.Form.Course.Event) {
		errs = append(errs, "Course must be one of claim, offer, conclude, delete, undelete")
	}
	if t.Form.Course.DefaultView != "" && !string_utils.Include([]string{"feed", "wiki", "modules", "syllabus", "assignments"}, t.Form.Course.DefaultView) {
		errs = append(errs, "Course must be one of feed, wiki, modules, syllabus, assignments")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCourse) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
