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

// EnrollUserCourses Create a new user enrollment for a course or section.
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Enrollment (Optional) The start time of the enrollment, in ISO8601 format. e.g. 2012-04-18T23:08:51Z
// # Enrollment (Optional) The end time of the enrollment, in ISO8601 format. e.g. 2012-04-18T23:08:51Z
// # Enrollment (Required) The ID of the user to be enrolled in the course.
// # Enrollment (Required) . Must be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, ObserverEnrollment, DesignerEnrollmentEnroll the user as a student, teacher, TA, observer, or designer. If no
//    value is given, the type will be inferred by enrollment[role] if supplied,
//    otherwise 'StudentEnrollment' will be used.
// # Enrollment (Optional) Assigns a custom course-level role to the user.
// # Enrollment (Optional) Assigns a custom course-level role to the user.
// # Enrollment (Optional) . Must be one of active, invited, inactiveIf set to 'active,' student will be immediately enrolled in the course.
//    Otherwise they will be required to accept a course invitation. Default is
//    'invited.'.
//
//    If set to 'inactive', student will be listed in the course roster for
//    teachers, but will not be able to participate in the course until
//    their enrollment is activated.
// # Enrollment (Optional) The ID of the course section to enroll the student in. If the
//    section-specific URL is used, this argument is redundant and will be
//    ignored.
// # Enrollment (Optional) If set, the enrollment will only allow the user to see and interact with
//    users enrolled in the section given by course_section_id.
//    * For teachers and TAs, this includes grading privileges.
//    * Section-limited students will not see any users (including teachers
//      and TAs) not enrolled in their sections.
//    * Users may have other enrollments that grant privileges to
//      multiple sections in the same course.
// # Enrollment (Optional) If true, a notification will be sent to the enrolled user.
//    Notifications are not sent by default.
// # Enrollment (Optional) If the current user is not allowed to manage enrollments in this
//    course, but the course allows self-enrollment, the user can self-
//    enroll as a student in the default section by passing in a valid
//    code. When self-enrolling, the user_id must be 'self'. The
//    enrollment_state will be set to 'active' and all other arguments
//    will be ignored.
// # Enrollment (Optional) If true, marks the enrollment as a self-enrollment, which gives
//    students the ability to drop the course if desired. Defaults to false.
// # Enrollment (Optional) For an observer enrollment, the ID of a student to observe.
//    This is a one-off operation; to automatically observe all a
//    student's enrollments (for example, as a parent), please use
//    the {api:UserObserveesController#create User Observees API}.
//
type EnrollUserCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Enrollment struct {
			StartAt                        time.Time `json:"start_at" url:"start_at,omitempty"`                                                     //  (Optional)
			EndAt                          time.Time `json:"end_at" url:"end_at,omitempty"`                                                         //  (Optional)
			UserID                         string    `json:"user_id" url:"user_id,omitempty"`                                                       //  (Required)
			Type                           string    `json:"type" url:"type,omitempty"`                                                             //  (Required) . Must be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, ObserverEnrollment, DesignerEnrollment
			Role                           string    `json:"role" url:"role,omitempty"`                                                             //  (Optional)
			RoleID                         int64     `json:"role_id" url:"role_id,omitempty"`                                                       //  (Optional)
			EnrollmentState                string    `json:"enrollment_state" url:"enrollment_state,omitempty"`                                     //  (Optional) . Must be one of active, invited, inactive
			CourseSectionID                int64     `json:"course_section_id" url:"course_section_id,omitempty"`                                   //  (Optional)
			LimitPrivilegesToCourseSection bool      `json:"limit_privileges_to_course_section" url:"limit_privileges_to_course_section,omitempty"` //  (Optional)
			Notify                         bool      `json:"notify" url:"notify,omitempty"`                                                         //  (Optional)
			SelfEnrollmentCode             string    `json:"self_enrollment_code" url:"self_enrollment_code,omitempty"`                             //  (Optional)
			SelfEnrolled                   bool      `json:"self_enrolled" url:"self_enrolled,omitempty"`                                           //  (Optional)
			AssociatedUserID               int64     `json:"associated_user_id" url:"associated_user_id,omitempty"`                                 //  (Optional)
		} `json:"enrollment" url:"enrollment,omitempty"`
	} `json:"form"`
}

func (t *EnrollUserCourses) GetMethod() string {
	return "POST"
}

func (t *EnrollUserCourses) GetURLPath() string {
	path := "courses/{course_id}/enrollments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *EnrollUserCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *EnrollUserCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EnrollUserCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EnrollUserCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Enrollment.UserID == "" {
		errs = append(errs, "'Enrollment' is required")
	}
	if t.Form.Enrollment.Type == "" {
		errs = append(errs, "'Enrollment' is required")
	}
	if t.Form.Enrollment.Type != "" && !string_utils.Include([]string{"StudentEnrollment", "TeacherEnrollment", "TaEnrollment", "ObserverEnrollment", "DesignerEnrollment"}, t.Form.Enrollment.Type) {
		errs = append(errs, "Enrollment must be one of StudentEnrollment, TeacherEnrollment, TaEnrollment, ObserverEnrollment, DesignerEnrollment")
	}
	if t.Form.Enrollment.EnrollmentState != "" && !string_utils.Include([]string{"active", "invited", "inactive"}, t.Form.Enrollment.EnrollmentState) {
		errs = append(errs, "Enrollment must be one of active, invited, inactive")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EnrollUserCourses) Do(c *canvasapi.Canvas) (*models.Enrollment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Enrollment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
