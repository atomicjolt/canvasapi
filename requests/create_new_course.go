package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateNewCourse Create a new course
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
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
// # Course (Optional) Set to true to make the course syllabus public for authenticated users.
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
// # Course (Optional) Set to true to weight final grade based on assignment groups percentages.
// # Course (Optional) The time zone for the course. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # Offer (Optional) If this option is set to true, the course will be available to students
//    immediately.
// # EnrollMe (Optional) Set to true to enroll the current user as the teacher.
// # Course (Optional) . Must be one of feed, wiki, modules, syllabus, assignmentsThe type of page that users will see when they first visit the course
//    * 'feed' Recent Activity Dashboard
//    * 'modules' Course Modules/Sections Page
//    * 'assignments' Course Assignments List
//    * 'syllabus' Course Syllabus Page
//    other types may be added in the future
// # Course (Optional) The syllabus body for the course
// # Course (Optional) The grading standard id to set for the course.  If no value is provided for this argument the current grading_standard will be un-set from this course.
// # Course (Optional) Optional. The grade_passback_setting for the course. Only 'nightly_sync', 'disabled', and '' are allowed
// # Course (Optional) Optional. Specifies the format of the course. (Should be 'on_campus', 'online', or 'blended')
// # EnableSISReactivation (Optional) When true, will first try to re-activate a deleted course with matching sis_course_id if possible.
//
type CreateNewCourse struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Course struct {
			Name                             string    `json:"name" url:"name,omitempty"`                                                                 //  (Optional)
			CourseCode                       string    `json:"course_code" url:"course_code,omitempty"`                                                   //  (Optional)
			StartAt                          time.Time `json:"start_at" url:"start_at,omitempty"`                                                         //  (Optional)
			EndAt                            time.Time `json:"end_at" url:"end_at,omitempty"`                                                             //  (Optional)
			License                          string    `json:"license" url:"license,omitempty"`                                                           //  (Optional)
			IsPublic                         bool      `json:"is_public" url:"is_public,omitempty"`                                                       //  (Optional)
			IsPublicToAuthUsers              bool      `json:"is_public_to_auth_users" url:"is_public_to_auth_users,omitempty"`                           //  (Optional)
			PublicSyllabus                   bool      `json:"public_syllabus" url:"public_syllabus,omitempty"`                                           //  (Optional)
			PublicSyllabusToAuth             bool      `json:"public_syllabus_to_auth" url:"public_syllabus_to_auth,omitempty"`                           //  (Optional)
			PublicDescription                string    `json:"public_description" url:"public_description,omitempty"`                                     //  (Optional)
			AllowStudentWikiEdits            bool      `json:"allow_student_wiki_edits" url:"allow_student_wiki_edits,omitempty"`                         //  (Optional)
			AllowWikiComments                bool      `json:"allow_wiki_comments" url:"allow_wiki_comments,omitempty"`                                   //  (Optional)
			AllowStudentForumAttachments     bool      `json:"allow_student_forum_attachments" url:"allow_student_forum_attachments,omitempty"`           //  (Optional)
			OpenEnrollment                   bool      `json:"open_enrollment" url:"open_enrollment,omitempty"`                                           //  (Optional)
			SelfEnrollment                   bool      `json:"self_enrollment" url:"self_enrollment,omitempty"`                                           //  (Optional)
			RestrictEnrollmentsToCourseDates bool      `json:"restrict_enrollments_to_course_dates" url:"restrict_enrollments_to_course_dates,omitempty"` //  (Optional)
			TermID                           string    `json:"term_id" url:"term_id,omitempty"`                                                           //  (Optional)
			SISCourseID                      string    `json:"sis_course_id" url:"sis_course_id,omitempty"`                                               //  (Optional)
			IntegrationID                    string    `json:"integration_id" url:"integration_id,omitempty"`                                             //  (Optional)
			HideFinalGrades                  bool      `json:"hide_final_grades" url:"hide_final_grades,omitempty"`                                       //  (Optional)
			ApplyAssignmentGroupWeights      bool      `json:"apply_assignment_group_weights" url:"apply_assignment_group_weights,omitempty"`             //  (Optional)
			TimeZone                         string    `json:"time_zone" url:"time_zone,omitempty"`                                                       //  (Optional)
			DefaultView                      string    `json:"default_view" url:"default_view,omitempty"`                                                 //  (Optional) . Must be one of feed, wiki, modules, syllabus, assignments
			SyllabusBody                     string    `json:"syllabus_body" url:"syllabus_body,omitempty"`                                               //  (Optional)
			GradingStandardID                int64     `json:"grading_standard_id" url:"grading_standard_id,omitempty"`                                   //  (Optional)
			GradePassbackSetting             string    `json:"grade_passback_setting" url:"grade_passback_setting,omitempty"`                             //  (Optional)
			CourseFormat                     string    `json:"course_format" url:"course_format,omitempty"`                                               //  (Optional)
		} `json:"course" url:"course,omitempty"`

		Offer                 bool `json:"offer" url:"offer,omitempty"`                                     //  (Optional)
		EnrollMe              bool `json:"enroll_me" url:"enroll_me,omitempty"`                             //  (Optional)
		EnableSISReactivation bool `json:"enable_sis_reactivation" url:"enable_sis_reactivation,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *CreateNewCourse) GetMethod() string {
	return "POST"
}

func (t *CreateNewCourse) GetURLPath() string {
	path := "accounts/{account_id}/courses"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateNewCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateNewCourse) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateNewCourse) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateNewCourse) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.Course.DefaultView != "" && !string_utils.Include([]string{"feed", "wiki", "modules", "syllabus", "assignments"}, t.Form.Course.DefaultView) {
		errs = append(errs, "Course must be one of feed, wiki, modules, syllabus, assignments")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateNewCourse) Do(c *canvasapi.Canvas) (*models.Course, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Course{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
