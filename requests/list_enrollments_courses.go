package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListEnrollmentsCourses Depending on the URL given, return a paginated list of either (1) all of
// the enrollments in a course, (2) all of the enrollments in a section or (3)
// all of a user's enrollments. This includes student, teacher, TA, and
// observer enrollments.
//
// If a user has multiple enrollments in a context (e.g. as a teacher
// and a student or in multiple course sections), each enrollment will be
// listed separately.
//
// note: Currently, only a root level admin user can return other users' enrollments. A
// user can, however, return his/her own enrollments.
// https://canvas.instructure.com/doc/api/enrollments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.Type (Optional) A list of enrollment types to return. Accepted values are
//    'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment',
//    'DesignerEnrollment', and 'ObserverEnrollment.' If omitted, all enrollment
//    types are returned. This argument is ignored if `role` is given.
// # Query.Role (Optional) A list of enrollment roles to return. Accepted values include course-level
//    roles created by the {api:RoleOverridesController#add_role Add Role API}
//    as well as the base enrollment types accepted by the `type` argument above.
// # Query.State (Optional) . Must be one of active, invited, creation_pending, deleted, rejected, completed, inactive, current_and_invited, current_and_future, current_and_concludedFilter by enrollment state. If omitted, 'active' and 'invited' enrollments
//    are returned. The following synthetic states are supported only when
//    querying a user's enrollments (either via user_id argument or via user
//    enrollments endpoint): +current_and_invited+, +current_and_future+, +current_and_concluded+
// # Query.Include (Optional) . Must be one of avatar_url, group_ids, locked, observed_users, can_be_removed, uuid, current_pointsArray of additional information to include on the enrollment or user records.
//    "avatar_url" and "group_ids" will be returned on the user record. If "current_points"
//    is specified, the fields "current_points" and (if the caller has
//    permissions to manage grades) "unposted_current_points" will be included
//    in the "grades" hash for student enrollments.
// # Query.UserID (Optional) Filter by user_id (only valid for course or section enrollment
//    queries). If set to the current user's id, this is a way to
//    determine if the user has any enrollments in the course or section,
//    independent of whether the user has permission to view other people
//    on the roster.
// # Query.GradingPeriodID (Optional) Return grades for the given grading_period.  If this parameter is not
//    specified, the returned grades will be for the whole course.
// # Query.EnrollmentTermID (Optional) Returns only enrollments for the specified enrollment term. This parameter
//    only applies to the user enrollments path. May pass the ID from the
//    enrollment terms api or the SIS id prepended with 'sis_term_id:'.
// # Query.SISAccountID (Optional) Returns only enrollments for the specified SIS account ID(s). Does not
//    look into sub_accounts. May pass in array or string.
// # Query.SISCourseID (Optional) Returns only enrollments matching the specified SIS course ID(s).
//    May pass in array or string.
// # Query.SISSectionID (Optional) Returns only section enrollments matching the specified SIS section ID(s).
//    May pass in array or string.
// # Query.SISUserID (Optional) Returns only enrollments for the specified SIS user ID(s). May pass in
//    array or string.
// # Query.CreatedForSISID (Optional) If sis_user_id is present and created_for_sis_id is true, Returns only
//    enrollments for the specified SIS ID(s).
//    If a user has two sis_id's, one enrollment may be created using one of the
//    two ids. This would limit the enrollments returned from the endpoint to
//    enrollments that were created from a sis_import with that sis_user_id
//
type ListEnrollmentsCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Type             []string `json:"type" url:"type,omitempty"`                             //  (Optional)
		Role             []string `json:"role" url:"role,omitempty"`                             //  (Optional)
		State            []string `json:"state" url:"state,omitempty"`                           //  (Optional) . Must be one of active, invited, creation_pending, deleted, rejected, completed, inactive, current_and_invited, current_and_future, current_and_concluded
		Include          []string `json:"include" url:"include,omitempty"`                       //  (Optional) . Must be one of avatar_url, group_ids, locked, observed_users, can_be_removed, uuid, current_points
		UserID           string   `json:"user_id" url:"user_id,omitempty"`                       //  (Optional)
		GradingPeriodID  int64    `json:"grading_period_id" url:"grading_period_id,omitempty"`   //  (Optional)
		EnrollmentTermID int64    `json:"enrollment_term_id" url:"enrollment_term_id,omitempty"` //  (Optional)
		SISAccountID     []string `json:"sis_account_id" url:"sis_account_id,omitempty"`         //  (Optional)
		SISCourseID      []string `json:"sis_course_id" url:"sis_course_id,omitempty"`           //  (Optional)
		SISSectionID     []string `json:"sis_section_id" url:"sis_section_id,omitempty"`         //  (Optional)
		SISUserID        []string `json:"sis_user_id" url:"sis_user_id,omitempty"`               //  (Optional)
		CreatedForSISID  []string `json:"created_for_sis_id" url:"created_for_sis_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListEnrollmentsCourses) GetMethod() string {
	return "GET"
}

func (t *ListEnrollmentsCourses) GetURLPath() string {
	path := "courses/{course_id}/enrollments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListEnrollmentsCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListEnrollmentsCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEnrollmentsCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEnrollmentsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	for _, v := range t.Query.State {
		if v != "" && !string_utils.Include([]string{"active", "invited", "creation_pending", "deleted", "rejected", "completed", "inactive", "current_and_invited", "current_and_future", "current_and_concluded"}, v) {
			errs = append(errs, "State must be one of active, invited, creation_pending, deleted, rejected, completed, inactive, current_and_invited, current_and_future, current_and_concluded")
		}
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"avatar_url", "group_ids", "locked", "observed_users", "can_be_removed", "uuid", "current_points"}, v) {
			errs = append(errs, "Include must be one of avatar_url, group_ids, locked, observed_users, can_be_removed, uuid, current_points")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEnrollmentsCourses) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Enrollment, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Enrollment{}
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
