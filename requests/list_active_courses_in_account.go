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

// ListActiveCoursesInAccount Retrieve a paginated list of courses in this account.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # WithEnrollments (Optional) If true, include only courses with at least one enrollment.  If false,
//    include only courses with no enrollments.  If not present, do not filter
//    on course enrollment status.
// # EnrollmentType (Optional) . Must be one of teacher, student, ta, observer, designerIf set, only return courses that have at least one user enrolled in
//    in the course with one of the specified enrollment types.
// # Published (Optional) If true, include only published courses.  If false, exclude published
//    courses.  If not present, do not filter on published status.
// # Completed (Optional) If true, include only completed courses (these may be in state
//    'completed', or their enrollment term may have ended).  If false, exclude
//    completed courses.  If not present, do not filter on completed status.
// # Blueprint (Optional) If true, include only blueprint courses. If false, exclude them.
//    If not present, do not filter on this basis.
// # BlueprintAssociated (Optional) If true, include only courses that inherit content from a blueprint course.
//    If false, exclude them. If not present, do not filter on this basis.
// # ByTeachers (Optional) List of User IDs of teachers; if supplied, include only courses taught by
//    one of the referenced users.
// # BySubaccounts (Optional) List of Account IDs; if supplied, include only courses associated with one
//    of the referenced subaccounts.
// # HideEnrollmentlessCourses (Optional) If present, only return courses that have at least one enrollment.
//    Equivalent to 'with_enrollments=true'; retained for compatibility.
// # State (Optional) . Must be one of created, claimed, available, completed, deleted, allIf set, only return courses that are in the given state(s). By default,
//    all states but "deleted" are returned.
// # EnrollmentTermID (Optional) If set, only includes courses from the specified term.
// # SearchTerm (Optional) The partial course name, code, or full ID to match and return in the results list. Must be at least 3 characters.
// # Include (Optional) . Must be one of syllabus_body, term, course_progress, storage_quota_used_mb, total_students, teachers, account_name, concluded- All explanations can be seen in the {api:CoursesController#index Course API index documentation}
//    - "sections", "needs_grading_count" and "total_scores" are not valid options at the account level
// # Sort (Optional) . Must be one of course_name, sis_course_id, teacher, account_nameThe column to sort results by.
// # Order (Optional) . Must be one of asc, descThe order to sort the given column by.
// # SearchBy (Optional) . Must be one of course, teacherThe filter to search by. "course" searches for course names, course codes,
//    and SIS IDs. "teacher" searches for teacher names
// # StartsBefore (Optional) If set, only return courses that start before the value (inclusive)
//    or their enrollment term starts before the value (inclusive)
//    or both the course's start_at and the enrollment term's start_at are set to null.
//    The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # EndsAfter (Optional) If set, only return courses that end after the value (inclusive)
//    or their enrollment term ends after the value (inclusive)
//    or both the course's end_at and the enrollment term's end_at are set to null.
//    The value should be formatted as: yyyy-mm-dd or ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # Homeroom (Optional) If set, only return homeroom courses.
//
type ListActiveCoursesInAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		WithEnrollments           bool      `json:"with_enrollments" url:"with_enrollments,omitempty"`                       //  (Optional)
		EnrollmentType            []string  `json:"enrollment_type" url:"enrollment_type,omitempty"`                         //  (Optional) . Must be one of teacher, student, ta, observer, designer
		Published                 bool      `json:"published" url:"published,omitempty"`                                     //  (Optional)
		Completed                 bool      `json:"completed" url:"completed,omitempty"`                                     //  (Optional)
		Blueprint                 bool      `json:"blueprint" url:"blueprint,omitempty"`                                     //  (Optional)
		BlueprintAssociated       bool      `json:"blueprint_associated" url:"blueprint_associated,omitempty"`               //  (Optional)
		ByTeachers                []int64   `json:"by_teachers" url:"by_teachers,omitempty"`                                 //  (Optional)
		BySubaccounts             []int64   `json:"by_subaccounts" url:"by_subaccounts,omitempty"`                           //  (Optional)
		HideEnrollmentlessCourses bool      `json:"hide_enrollmentless_courses" url:"hide_enrollmentless_courses,omitempty"` //  (Optional)
		State                     []string  `json:"state" url:"state,omitempty"`                                             //  (Optional) . Must be one of created, claimed, available, completed, deleted, all
		EnrollmentTermID          int64     `json:"enrollment_term_id" url:"enrollment_term_id,omitempty"`                   //  (Optional)
		SearchTerm                string    `json:"search_term" url:"search_term,omitempty"`                                 //  (Optional)
		Include                   []string  `json:"include" url:"include,omitempty"`                                         //  (Optional) . Must be one of syllabus_body, term, course_progress, storage_quota_used_mb, total_students, teachers, account_name, concluded
		Sort                      string    `json:"sort" url:"sort,omitempty"`                                               //  (Optional) . Must be one of course_name, sis_course_id, teacher, account_name
		Order                     string    `json:"order" url:"order,omitempty"`                                             //  (Optional) . Must be one of asc, desc
		SearchBy                  string    `json:"search_by" url:"search_by,omitempty"`                                     //  (Optional) . Must be one of course, teacher
		StartsBefore              time.Time `json:"starts_before" url:"starts_before,omitempty"`                             //  (Optional)
		EndsAfter                 time.Time `json:"ends_after" url:"ends_after,omitempty"`                                   //  (Optional)
		Homeroom                  bool      `json:"homeroom" url:"homeroom,omitempty"`                                       //  (Optional)
	} `json:"query"`
}

func (t *ListActiveCoursesInAccount) GetMethod() string {
	return "GET"
}

func (t *ListActiveCoursesInAccount) GetURLPath() string {
	path := "accounts/{account_id}/courses"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListActiveCoursesInAccount) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListActiveCoursesInAccount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListActiveCoursesInAccount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListActiveCoursesInAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	for _, v := range t.Query.EnrollmentType {
		if v != "" && !string_utils.Include([]string{"teacher", "student", "ta", "observer", "designer"}, v) {
			errs = append(errs, "EnrollmentType must be one of teacher, student, ta, observer, designer")
		}
	}
	for _, v := range t.Query.State {
		if v != "" && !string_utils.Include([]string{"created", "claimed", "available", "completed", "deleted", "all"}, v) {
			errs = append(errs, "State must be one of created, claimed, available, completed, deleted, all")
		}
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"syllabus_body", "term", "course_progress", "storage_quota_used_mb", "total_students", "teachers", "account_name", "concluded"}, v) {
			errs = append(errs, "Include must be one of syllabus_body, term, course_progress, storage_quota_used_mb, total_students, teachers, account_name, concluded")
		}
	}
	if t.Query.Sort != "" && !string_utils.Include([]string{"course_name", "sis_course_id", "teacher", "account_name"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of course_name, sis_course_id, teacher, account_name")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	if t.Query.SearchBy != "" && !string_utils.Include([]string{"course", "teacher"}, t.Query.SearchBy) {
		errs = append(errs, "SearchBy must be one of course, teacher")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListActiveCoursesInAccount) Do(c *canvasapi.Canvas) ([]*models.Course, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Course{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
