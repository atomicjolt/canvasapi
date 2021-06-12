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

// ListEnrollmentsSections Depending on the URL given, return a paginated list of either (1) all of
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
// # SectionID (Required) ID
//
// Query Parameters:
// # Type (Optional) A list of enrollment types to return. Accepted values are
//    'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment',
//    'DesignerEnrollment', and 'ObserverEnrollment.' If omitted, all enrollment
//    types are returned. This argument is ignored if `role` is given.
// # Role (Optional) A list of enrollment roles to return. Accepted values include course-level
//    roles created by the {api:RoleOverridesController#add_role Add Role API}
//    as well as the base enrollment types accepted by the `type` argument above.
// # State (Optional) . Must be one of active, invited, creation_pending, deleted, rejected, completed, inactive, current_and_invited, current_and_future, current_and_concludedFilter by enrollment state. If omitted, 'active' and 'invited' enrollments
//    are returned. The following synthetic states are supported only when
//    querying a user's enrollments (either via user_id argument or via user
//    enrollments endpoint): +current_and_invited+, +current_and_future+, +current_and_concluded+
// # Include (Optional) . Must be one of avatar_url, group_ids, locked, observed_users, can_be_removed, uuid, current_pointsArray of additional information to include on the enrollment or user records.
//    "avatar_url" and "group_ids" will be returned on the user record. If "current_points"
//    is specified, the fields "current_points" and (if the caller has
//    permissions to manage grades) "unposted_current_points" will be included
//    in the "grades" hash for student enrollments.
// # UserID (Optional) Filter by user_id (only valid for course or section enrollment
//    queries). If set to the current user's id, this is a way to
//    determine if the user has any enrollments in the course or section,
//    independent of whether the user has permission to view other people
//    on the roster.
// # GradingPeriodID (Optional) Return grades for the given grading_period.  If this parameter is not
//    specified, the returned grades will be for the whole course.
// # EnrollmentTermID (Optional) Returns only enrollments for the specified enrollment term. This parameter
//    only applies to the user enrollments path. May pass the ID from the
//    enrollment terms api or the SIS id prepended with 'sis_term_id:'.
// # SISAccountID (Optional) Returns only enrollments for the specified SIS account ID(s). Does not
//    look into sub_accounts. May pass in array or string.
// # SISCourseID (Optional) Returns only enrollments matching the specified SIS course ID(s).
//    May pass in array or string.
// # SISSectionID (Optional) Returns only section enrollments matching the specified SIS section ID(s).
//    May pass in array or string.
// # SISUserID (Optional) Returns only enrollments for the specified SIS user ID(s). May pass in
//    array or string.
// # CreatedForSISID (Optional) If sis_user_id is present and created_for_sis_id is true, Returns only
//    enrollments for the specified SIS ID(s).
//    If a user has two sis_id's, one enrollment may be created using one of the
//    two ids. This would limit the enrollments returned from the endpoint to
//    enrollments that were created from a sis_import with that sis_user_id
//
type ListEnrollmentsSections struct {
	Path struct {
		SectionID string `json:"section_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Type             []string `json:"type"`               //  (Optional)
		Role             []string `json:"role"`               //  (Optional)
		State            []string `json:"state"`              //  (Optional) . Must be one of active, invited, creation_pending, deleted, rejected, completed, inactive, current_and_invited, current_and_future, current_and_concluded
		Include          []string `json:"include"`            //  (Optional) . Must be one of avatar_url, group_ids, locked, observed_users, can_be_removed, uuid, current_points
		UserID           string   `json:"user_id"`            //  (Optional)
		GradingPeriodID  int64    `json:"grading_period_id"`  //  (Optional)
		EnrollmentTermID int64    `json:"enrollment_term_id"` //  (Optional)
		SISAccountID     []string `json:"sis_account_id"`     //  (Optional)
		SISCourseID      []string `json:"sis_course_id"`      //  (Optional)
		SISSectionID     []string `json:"sis_section_id"`     //  (Optional)
		SISUserID        []string `json:"sis_user_id"`        //  (Optional)
		CreatedForSISID  []bool   `json:"created_for_sis_id"` //  (Optional)
	} `json:"query"`
}

func (t *ListEnrollmentsSections) GetMethod() string {
	return "GET"
}

func (t *ListEnrollmentsSections) GetURLPath() string {
	path := "sections/{section_id}/enrollments"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	return path
}

func (t *ListEnrollmentsSections) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListEnrollmentsSections) GetBody() (string, error) {
	return "", nil
}

func (t *ListEnrollmentsSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'SectionID' is required")
	}
	for _, v := range t.Query.State {
		if !string_utils.Include([]string{"active", "invited", "creation_pending", "deleted", "rejected", "completed", "inactive", "current_and_invited", "current_and_future", "current_and_concluded"}, v) {
			errs = append(errs, "State must be one of active, invited, creation_pending, deleted, rejected, completed, inactive, current_and_invited, current_and_future, current_and_concluded")
		}
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"avatar_url", "group_ids", "locked", "observed_users", "can_be_removed", "uuid", "current_points"}, v) {
			errs = append(errs, "Include must be one of avatar_url, group_ids, locked, observed_users, can_be_removed, uuid, current_points")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEnrollmentsSections) Do(c *canvasapi.Canvas) ([]*models.Enrollment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Enrollment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
