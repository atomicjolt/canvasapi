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

// ListYourCourses Returns the paginated list of active courses for the current user.
// https://canvas.instructure.com/doc/api/courses.html
//
// Query Parameters:
// # Query.EnrollmentType (Optional) . Must be one of teacher, student, ta, observer, designerWhen set, only return courses where the user is enrolled as this type. For
//    example, set to "teacher" to return only courses where the user is
//    enrolled as a Teacher.  This argument is ignored if enrollment_role is given.
// # Query.EnrollmentRole (Optional) Deprecated
//    When set, only return courses where the user is enrolled with the specified
//    course-level role.  This can be a role created with the
//    {api:RoleOverridesController#add_role Add Role API} or a base role type of
//    'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment', 'ObserverEnrollment',
//    or 'DesignerEnrollment'.
// # Query.EnrollmentRoleID (Optional) When set, only return courses where the user is enrolled with the specified
//    course-level role.  This can be a role created with the
//    {api:RoleOverridesController#add_role Add Role API} or a built_in role type of
//    'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment', 'ObserverEnrollment',
//    or 'DesignerEnrollment'.
// # Query.EnrollmentState (Optional) . Must be one of active, invited_or_pending, completedWhen set, only return courses where the user has an enrollment with the given state.
//    This will respect section/course/term date overrides.
// # Query.ExcludeBlueprintCourses (Optional) When set, only return courses that are not configured as blueprint courses.
// # Query.Include (Optional) . Must be one of needs_grading_count, syllabus_body, public_description, total_scores, current_grading_period_scores, grading_periods, term, account, course_progress, sections, storage_quota_used_mb, total_students, passback_status, favorites, teachers, observed_users, course_image, concluded- "needs_grading_count": Optional information to include with each Course.
//      When needs_grading_count is given, and the current user has grading
//      rights, the total number of submissions needing grading for all
//      assignments is returned.
//    - "syllabus_body": Optional information to include with each Course.
//      When syllabus_body is given the user-generated html for the course
//      syllabus is returned.
//    - "public_description": Optional information to include with each Course.
//      When public_description is given the user-generated text for the course
//      public description is returned.
//    - "total_scores": Optional information to include with each Course.
//      When total_scores is given, any student enrollments will also
//      include the fields 'computed_current_score', 'computed_final_score',
//      'computed_current_grade', and 'computed_final_grade', as well as (if
//      the user has permission) 'unposted_current_score',
//      'unposted_final_score', 'unposted_current_grade', and
//      'unposted_final_grade' (see Enrollment documentation for more
//      information on these fields). This argument is ignored if the course is
//      configured to hide final grades.
//    - "current_grading_period_scores": Optional information to include with
//      each Course. When current_grading_period_scores is given and total_scores
//      is given, any student enrollments will also include the fields
//      'has_grading_periods',
//      'totals_for_all_grading_periods_option', 'current_grading_period_title',
//      'current_grading_period_id', current_period_computed_current_score',
//      'current_period_computed_final_score',
//      'current_period_computed_current_grade', and
//      'current_period_computed_final_grade', as well as (if the user has permission)
//      'current_period_unposted_current_score',
//      'current_period_unposted_final_score',
//      'current_period_unposted_current_grade', and
//      'current_period_unposted_final_grade' (see Enrollment documentation for
//      more information on these fields). In addition, when this argument is
//      passed, the course will have a 'has_grading_periods' attribute
//      on it. This argument is ignored if the total_scores argument is not
//      included. If the course is configured to hide final grades, the
//      following fields are not returned:
//      'totals_for_all_grading_periods_option',
//      'current_period_computed_current_score',
//      'current_period_computed_final_score',
//      'current_period_computed_current_grade',
//      'current_period_computed_final_grade',
//      'current_period_unposted_current_score',
//      'current_period_unposted_final_score',
//      'current_period_unposted_current_grade', and
//      'current_period_unposted_final_grade'
//    - "grading_periods": Optional information to include with each Course. When
//      grading_periods is given, a list of the grading periods associated with
//      each course is returned.
//    - "term": Optional information to include with each Course. When
//      term is given, the information for the enrollment term for each course
//      is returned.
//    - "account": Optional information to include with each Course. When
//      account is given, the account json for each course is returned.
//    - "course_progress": Optional information to include with each Course.
//      When course_progress is given, each course will include a
//      'course_progress' object with the fields: 'requirement_count', an integer
//      specifying the total number of requirements in the course,
//      'requirement_completed_count', an integer specifying the total number of
//      requirements in this course that have been completed, and
//      'next_requirement_url', a string url to the next requirement item, and
//      'completed_at', the date the course was completed (null if incomplete).
//      'next_requirement_url' will be null if all requirements have been
//      completed or the current module does not require sequential progress.
//      "course_progress" will return an error message if the course is not
//      module based or the user is not enrolled as a student in the course.
//    - "sections": Section enrollment information to include with each Course.
//      Returns an array of hashes containing the section ID (id), section name
//      (name), start and end dates (start_at, end_at), as well as the enrollment
//      type (enrollment_role, e.g. 'StudentEnrollment').
//    - "storage_quota_used_mb": The amount of storage space used by the files in this course
//    - "total_students": Optional information to include with each Course.
//      Returns an integer for the total amount of active and invited students.
//    - "passback_status": Include the grade passback_status
//    - "favorites": Optional information to include with each Course.
//      Indicates if the user has marked the course as a favorite course.
//    - "teachers": Teacher information to include with each Course.
//      Returns an array of hashes containing the {api:Users:UserDisplay UserDisplay} information
//      for each teacher in the course.
//    - "observed_users": Optional information to include with each Course.
//      Will include data for observed users if the current user has an
//      observer enrollment.
//    - "tabs": Optional information to include with each Course.
//      Will include the list of tabs configured for each course.  See the
//      {api:TabsController#index List available tabs API} for more information.
//    - "course_image": Optional course image data for when there is a course image
//      and the course image feature flag has been enabled
//    - "concluded": Optional information to include with each Course. Indicates whether
//      the course has been concluded, taking course and term dates into account.
// # Query.State (Optional) . Must be one of unpublished, available, completed, deletedIf set, only return courses that are in the given state(s).
//    By default, "available" is returned for students and observers, and
//    anything except "deleted", for all other enrollment types
//
type ListYourCourses struct {
	Query struct {
		EnrollmentType          string   `json:"enrollment_type" url:"enrollment_type,omitempty"`                     //  (Optional) . Must be one of teacher, student, ta, observer, designer
		EnrollmentRole          string   `json:"enrollment_role" url:"enrollment_role,omitempty"`                     //  (Optional)
		EnrollmentRoleID        int64    `json:"enrollment_role_id" url:"enrollment_role_id,omitempty"`               //  (Optional)
		EnrollmentState         string   `json:"enrollment_state" url:"enrollment_state,omitempty"`                   //  (Optional) . Must be one of active, invited_or_pending, completed
		ExcludeBlueprintCourses bool     `json:"exclude_blueprint_courses" url:"exclude_blueprint_courses,omitempty"` //  (Optional)
		Include                 []string `json:"include" url:"include,omitempty"`                                     //  (Optional) . Must be one of needs_grading_count, syllabus_body, public_description, total_scores, current_grading_period_scores, grading_periods, term, account, course_progress, sections, storage_quota_used_mb, total_students, passback_status, favorites, teachers, observed_users, course_image, concluded
		State                   []string `json:"state" url:"state,omitempty"`                                         //  (Optional) . Must be one of unpublished, available, completed, deleted
	} `json:"query"`
}

func (t *ListYourCourses) GetMethod() string {
	return "GET"
}

func (t *ListYourCourses) GetURLPath() string {
	return ""
}

func (t *ListYourCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListYourCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListYourCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListYourCourses) HasErrors() error {
	errs := []string{}
	if t.Query.EnrollmentType != "" && !string_utils.Include([]string{"teacher", "student", "ta", "observer", "designer"}, t.Query.EnrollmentType) {
		errs = append(errs, "EnrollmentType must be one of teacher, student, ta, observer, designer")
	}
	if t.Query.EnrollmentState != "" && !string_utils.Include([]string{"active", "invited_or_pending", "completed"}, t.Query.EnrollmentState) {
		errs = append(errs, "EnrollmentState must be one of active, invited_or_pending, completed")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"needs_grading_count", "syllabus_body", "public_description", "total_scores", "current_grading_period_scores", "grading_periods", "term", "account", "course_progress", "sections", "storage_quota_used_mb", "total_students", "passback_status", "favorites", "teachers", "observed_users", "course_image", "concluded"}, v) {
			errs = append(errs, "Include must be one of needs_grading_count, syllabus_body, public_description, total_scores, current_grading_period_scores, grading_periods, term, account, course_progress, sections, storage_quota_used_mb, total_students, passback_status, favorites, teachers, observed_users, course_image, concluded")
		}
	}
	for _, v := range t.Query.State {
		if v != "" && !string_utils.Include([]string{"unpublished", "available", "completed", "deleted"}, v) {
			errs = append(errs, "State must be one of unpublished, available, completed, deleted")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListYourCourses) Do(c *canvasapi.Canvas) ([]*models.Course, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Course{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
