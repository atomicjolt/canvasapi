package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
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
// # ID (Required) ID
//
// Form Parameters:
// # Course (Optional) The unique ID of the account to move the course to.
// # Course (Optional) The name of the course. If omitted, the course will be named "Unnamed
//    Course."
// # Course (Optional) The course code for the course.
// # Course (Optional) Course start date in ISO8601 format, e.g. 2011-01-01T01:00Z
// # Course (Optional) Course end date in ISO8601 format. e.g. 2011-01-01T01:00Z
// # Course (Optional) The name of the licensing. Should be one of the following abbreviations
//    (a descriptive name is included in parenthesis for reference):
//    - 'private' (Private Copyrighted)
//    - 'cc_by_nc_nd' (CC Attribution Non-Commercial No Derivatives)
//    - 'cc_by_nc_sa' (CC Attribution Non-Commercial Share Alike)
//    - 'cc_by_nc' (CC Attribution Non-Commercial)
//    - 'cc_by_nd' (CC Attribution No Derivatives)
//    - 'cc_by_sa' (CC Attribution Share Alike)
//    - 'cc_by' (CC Attribution)
//    - 'public_domain' (Public Domain).
// # Course (Optional) Set to true if course is public to both authenticated and unauthenticated users.
// # Course (Optional) Set to true if course is public only to authenticated users.
// # Course (Optional) Set to true to make the course syllabus public.
// # Course (Optional) Set to true to make the course syllabus to public for authenticated users.
// # Course (Optional) A publicly visible description of the course.
// # Course (Optional) If true, students will be able to modify the course wiki.
// # Course (Optional) If true, course members will be able to comment on wiki pages.
// # Course (Optional) If true, students can attach files to forum posts.
// # Course (Optional) Set to true if the course is open enrollment.
// # Course (Optional) Set to true if the course is self enrollment.
// # Course (Optional) Set to true to restrict user enrollments to the start and end dates of the
//    course. This parameter is required when using the API, as this option is
//    not displayed in the Course Settings page.
// # Course (Optional) The unique ID of the term to create to course in.
// # Course (Optional) The unique SIS identifier.
// # Course (Optional) The unique Integration identifier.
// # Course (Optional) If this option is set to true, the totals in student grades summary will
//    be hidden.
// # Course (Optional) The time zone for the course. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # Course (Optional) Set to true to weight final grade based on assignment groups percentages.
// # Course (Optional) Set the storage quota for the course, in megabytes. The caller must have
//    the "Manage storage quotas" account permission.
// # Offer (Optional) If this option is set to true, the course will be available to students
//    immediately.
// # Course (Optional) . Must be one of claim, offer, conclude, delete, undeleteThe action to take on each course.
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
// # Course (Optional) . Must be one of feed, wiki, modules, syllabus, assignmentsThe type of page that users will see when they first visit the course
//    * 'feed' Recent Activity Dashboard
//    * 'wiki' Wiki Front Page
//    * 'modules' Course Modules/Sections Page
//    * 'assignments' Course Assignments List
//    * 'syllabus' Course Syllabus Page
//    other types may be added in the future
// # Course (Optional) The syllabus body for the course
// # Course (Optional) Optional. Indicates whether the Course Summary (consisting of the course's assignments and calendar events) is displayed on the syllabus page. Defaults to +true+.
// # Course (Optional) The grading standard id to set for the course.  If no value is provided for this argument the current grading_standard will be un-set from this course.
// # Course (Optional) Optional. The grade_passback_setting for the course. Only 'nightly_sync' and '' are allowed
// # Course (Optional) Optional. Specifies the format of the course. (Should be either 'on_campus' or 'online')
// # Course (Optional) This is a file ID corresponding to an image file in the course that will
//    be used as the course image.
//    This will clear the course's image_url setting if set.  If you attempt
//    to provide image_url and image_id in a request it will fail.
// # Course (Optional) This is a URL to an image to be used as the course image.
//    This will clear the course's image_id setting if set.  If you attempt
//    to provide image_url and image_id in a request it will fail.
// # Course (Optional) If this option is set to true, the course image url and course image
//    ID are both set to nil
// # Course (Optional) Sets the course as a blueprint course.
// # Course (Optional) Sets a default set to apply to blueprint course objects when restricted,
//    unless _use_blueprint_restrictions_by_object_type_ is enabled.
//    See the {api:Blueprint_Courses:BlueprintRestriction Blueprint Restriction} documentation
// # Course (Optional) When enabled, the _blueprint_restrictions_ parameter will be ignored in favor of
//    the _blueprint_restrictions_by_object_type_ parameter
// # Course (Optional) Allows setting multiple {api:Blueprint_Courses:BlueprintRestriction Blueprint Restriction}
//    to apply to blueprint course objects of the matching type when restricted.
//    The possible object types are "assignment", "attachment", "discussion_topic", "quiz" and "wiki_page".
//    Example usage:
//      course[blueprint_restrictions_by_object_type][assignment][content]=1
// # Course (Optional) Sets the course as a homeroom course. The setting takes effect only when the course is associated
//    with a Canvas for Elementary-enabled account.
// # Course (Optional) Syncs enrollments from the homeroom that is set in homeroom_course_id. The setting only takes effect when the
//    course is associated with a Canvas for Elementary-enabled account and sync_enrollments_from_homeroom is enabled.
// # Course (Optional) Sets the Homeroom Course id to be used with sync_enrollments_from_homeroom. The setting only takes effect when the
//    course is associated with a Canvas for Elementary-enabled account and sync_enrollments_from_homeroom is enabled.
// # Course (Optional) Enable or disable the course as a template that can be selected by an account
// # Course (Optional) Sets a color in hex code format to be associated with the course. The setting takes effect only when the course
//    is associated with a Canvas for Elementary-enabled account.
//
type UpdateCourse struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Course struct {
			AccountID                            int64     `json:"account_id"`                                //  (Optional)
			Name                                 string    `json:"name"`                                      //  (Optional)
			CourseCode                           string    `json:"course_code"`                               //  (Optional)
			StartAt                              time.Time `json:"start_at"`                                  //  (Optional)
			EndAt                                time.Time `json:"end_at"`                                    //  (Optional)
			License                              string    `json:"license"`                                   //  (Optional)
			IsPublic                             bool      `json:"is_public"`                                 //  (Optional)
			IsPublicToAuthUsers                  bool      `json:"is_public_to_auth_users"`                   //  (Optional)
			PublicSyllabus                       bool      `json:"public_syllabus"`                           //  (Optional)
			PublicSyllabusToAuth                 bool      `json:"public_syllabus_to_auth"`                   //  (Optional)
			PublicDescription                    string    `json:"public_description"`                        //  (Optional)
			AllowStudentWikiEdits                bool      `json:"allow_student_wiki_edits"`                  //  (Optional)
			AllowWikiComments                    bool      `json:"allow_wiki_comments"`                       //  (Optional)
			AllowStudentForumAttachments         bool      `json:"allow_student_forum_attachments"`           //  (Optional)
			OpenEnrollment                       bool      `json:"open_enrollment"`                           //  (Optional)
			SelfEnrollment                       bool      `json:"self_enrollment"`                           //  (Optional)
			RestrictEnrollmentsToCourseDates     bool      `json:"restrict_enrollments_to_course_dates"`      //  (Optional)
			TermID                               int64     `json:"term_id"`                                   //  (Optional)
			SISCourseID                          string    `json:"sis_course_id"`                             //  (Optional)
			IntegrationID                        string    `json:"integration_id"`                            //  (Optional)
			HideFinalGrades                      bool      `json:"hide_final_grades"`                         //  (Optional)
			TimeZone                             string    `json:"time_zone"`                                 //  (Optional)
			ApplyAssignmentGroupWeights          bool      `json:"apply_assignment_group_weights"`            //  (Optional)
			StorageQuotaMb                       int64     `json:"storage_quota_mb"`                          //  (Optional)
			Event                                string    `json:"event"`                                     //  (Optional) . Must be one of claim, offer, conclude, delete, undelete
			DefaultView                          string    `json:"default_view"`                              //  (Optional) . Must be one of feed, wiki, modules, syllabus, assignments
			SyllabusBody                         string    `json:"syllabus_body"`                             //  (Optional)
			SyllabusCourseSummary                bool      `json:"syllabus_course_summary"`                   //  (Optional)
			GradingStandardID                    int64     `json:"grading_standard_id"`                       //  (Optional)
			GradePassbackSetting                 string    `json:"grade_passback_setting"`                    //  (Optional)
			CourseFormat                         string    `json:"course_format"`                             //  (Optional)
			ImageID                              int64     `json:"image_id"`                                  //  (Optional)
			ImageUrl                             string    `json:"image_url"`                                 //  (Optional)
			RemoveImage                          bool      `json:"remove_image"`                              //  (Optional)
			Blueprint                            bool      `json:"blueprint"`                                 //  (Optional)
			BlueprintRestrictions                string    `json:"blueprint_restrictions"`                    //  (Optional)
			UseBlueprintRestrictionsByObjectType bool      `json:"use_blueprint_restrictions_by_object_type"` //  (Optional)
			BlueprintRestrictionsByObjectType    string    `json:"blueprint_restrictions_by_object_type"`     //  (Optional)
			HomeroomCourse                       bool      `json:"homeroom_course"`                           //  (Optional)
			SyncEnrollmentsFromHomeroom          string    `json:"sync_enrollments_from_homeroom"`            //  (Optional)
			HomeroomCourseID                     string    `json:"homeroom_course_id"`                        //  (Optional)
			Template                             bool      `json:"template"`                                  //  (Optional)
			CourseColor                          string    `json:"course_color"`                              //  (Optional)
		} `json:"course"`

		Offer bool `json:"offer"` //  (Optional)
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

func (t *UpdateCourse) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateCourse) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"claim", "offer", "conclude", "delete", "undelete"}, t.Form.Course.Event) {
		errs = append(errs, "Course must be one of claim, offer, conclude, delete, undelete")
	}
	if !string_utils.Include([]string{"feed", "wiki", "modules", "syllabus", "assignments"}, t.Form.Course.DefaultView) {
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
