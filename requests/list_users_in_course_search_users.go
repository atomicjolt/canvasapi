package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListUsersInCourseSearchUsers Returns the paginated list of users in this course. And optionally the user's enrollments in the course.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # SearchTerm (Optional) The partial name or full ID of the users to match and return in the results list.
// # Sort (Optional) . Must be one of username, last_login, email, sis_idWhen set, sort the results of the search based on the given field.
// # EnrollmentType (Optional) . Must be one of teacher, student, student_view, ta, observer, designerWhen set, only return users where the user is enrolled as this type.
//    "student_view" implies include[]=test_student.
//    This argument is ignored if enrollment_role is given.
// # EnrollmentRole (Optional) Deprecated
//    When set, only return users enrolled with the specified course-level role.  This can be
//    a role created with the {api:RoleOverridesController#add_role Add Role API} or a
//    base role type of 'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment',
//    'ObserverEnrollment', or 'DesignerEnrollment'.
// # EnrollmentRoleID (Optional) When set, only return courses where the user is enrolled with the specified
//    course-level role.  This can be a role created with the
//    {api:RoleOverridesController#add_role Add Role API} or a built_in role id with type
//    'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment', 'ObserverEnrollment',
//    or 'DesignerEnrollment'.
// # Include (Optional) . Must be one of enrollments, locked, avatar_url, test_student, bio, custom_links, current_grading_period_scores, uuid- "enrollments":
//    Optionally include with each Course the user's current and invited
//    enrollments. If the user is enrolled as a student, and the account has
//    permission to manage or view all grades, each enrollment will include a
//    'grades' key with 'current_score', 'final_score', 'current_grade' and
//    'final_grade' values.
//    - "locked": Optionally include whether an enrollment is locked.
//    - "avatar_url": Optionally include avatar_url.
//    - "bio": Optionally include each user's bio.
//    - "test_student": Optionally include the course's Test Student,
//    if present. Default is to not include Test Student.
//    - "custom_links": Optionally include plugin-supplied custom links for each student,
//    such as analytics information
//    - "current_grading_period_scores": if enrollments is included as
//    well as this directive, the scores returned in the enrollment
//    will be for the current grading period if there is one. A
//    'grading_period_id' value will also be included with the
//    scores. if grading_period_id is nil there is no current grading
//    period and the score is a total score.
//    - "uuid": Optionally include the users uuid
// # UserID (Optional) If this parameter is given and it corresponds to a user in the course,
//    the +page+ parameter will be ignored and the page containing the specified user
//    will be returned instead.
// # UserIDs (Optional) If included, the course users set will only include users with IDs
//    specified by the param. Note: this will not work in conjunction
//    with the "user_id" argument but multiple user_ids can be included.
// # EnrollmentState (Optional) . Must be one of active, invited, rejected, completed, inactiveWhen set, only return users where the enrollment workflow state is of one of the given types.
//    "active" and "invited" enrollments are returned by default.
//
type ListUsersInCourseSearchUsers struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm       string   `json:"search_term"`        //  (Optional)
		Sort             string   `json:"sort"`               //  (Optional) . Must be one of username, last_login, email, sis_id
		EnrollmentType   []string `json:"enrollment_type"`    //  (Optional) . Must be one of teacher, student, student_view, ta, observer, designer
		EnrollmentRole   string   `json:"enrollment_role"`    //  (Optional)
		EnrollmentRoleID int64    `json:"enrollment_role_id"` //  (Optional)
		Include          []string `json:"include"`            //  (Optional) . Must be one of enrollments, locked, avatar_url, test_student, bio, custom_links, current_grading_period_scores, uuid
		UserID           string   `json:"user_id"`            //  (Optional)
		UserIDs          []int64  `json:"user_ids"`           //  (Optional)
		EnrollmentState  []string `json:"enrollment_state"`   //  (Optional) . Must be one of active, invited, rejected, completed, inactive
	} `json:"query"`
}

func (t *ListUsersInCourseSearchUsers) GetMethod() string {
	return "GET"
}

func (t *ListUsersInCourseSearchUsers) GetURLPath() string {
	path := "courses/{course_id}/search_users"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListUsersInCourseSearchUsers) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListUsersInCourseSearchUsers) GetBody() (string, error) {
	return "", nil
}

func (t *ListUsersInCourseSearchUsers) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"username", "last_login", "email", "sis_id"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of username, last_login, email, sis_id")
	}
	for _, v := range t.Query.EnrollmentType {
		if !string_utils.Include([]string{"teacher", "student", "student_view", "ta", "observer", "designer"}, v) {
			errs = append(errs, "EnrollmentType must be one of teacher, student, student_view, ta, observer, designer")
		}
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"enrollments", "locked", "avatar_url", "test_student", "bio", "custom_links", "current_grading_period_scores", "uuid"}, v) {
			errs = append(errs, "Include must be one of enrollments, locked, avatar_url, test_student, bio, custom_links, current_grading_period_scores, uuid")
		}
	}
	for _, v := range t.Query.EnrollmentState {
		if !string_utils.Include([]string{"active", "invited", "rejected", "completed", "inactive"}, v) {
			errs = append(errs, "EnrollmentState must be one of active, invited, rejected, completed, inactive")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUsersInCourseSearchUsers) Do(c *canvasapi.Canvas) ([]*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}