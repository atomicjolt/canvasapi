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

// GetSingleCourseCourses Return information on a single course.
//
// Accepts the same include[] parameters as the list action plus:
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of needs_grading_count, syllabus_body, public_description, total_scores, current_grading_period_scores, term, account, course_progress, sections, storage_quota_used_mb, total_students, passback_status, favorites, teachers, observed_users, all_courses, permissions, observed_users, course_image, concluded- "all_courses": Also search recently deleted courses.
//    - "permissions": Include permissions the current user has
//      for the course.
//    - "observed_users": include observed users in the enrollments
//    - "course_image": Optional course image data for when there is a course image
//      and the course image feature flag has been enabled
//    - "concluded": Optional information to include with each Course. Indicates whether
//      the course has been concluded, taking course and term dates into account.
// # Query.TeacherLimit (Optional) The maximum number of teacher enrollments to show.
//    If the course contains more teachers than this, instead of giving the teacher
//    enrollments, the count of teachers will be given under a _teacher_count_ key.
//
type GetSingleCourseCourses struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include      []string `json:"include" url:"include,omitempty"`             //  (Optional) . Must be one of needs_grading_count, syllabus_body, public_description, total_scores, current_grading_period_scores, term, account, course_progress, sections, storage_quota_used_mb, total_students, passback_status, favorites, teachers, observed_users, all_courses, permissions, observed_users, course_image, concluded
		TeacherLimit int64    `json:"teacher_limit" url:"teacher_limit,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetSingleCourseCourses) GetMethod() string {
	return "GET"
}

func (t *GetSingleCourseCourses) GetURLPath() string {
	path := "courses/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleCourseCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetSingleCourseCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleCourseCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleCourseCourses) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"needs_grading_count", "syllabus_body", "public_description", "total_scores", "current_grading_period_scores", "term", "account", "course_progress", "sections", "storage_quota_used_mb", "total_students", "passback_status", "favorites", "teachers", "observed_users", "all_courses", "permissions", "observed_users", "course_image", "concluded"}, v) {
			errs = append(errs, "Include must be one of needs_grading_count, syllabus_body, public_description, total_scores, current_grading_period_scores, term, account, course_progress, sections, storage_quota_used_mb, total_students, passback_status, favorites, teachers, observed_users, all_courses, permissions, observed_users, course_image, concluded")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleCourseCourses) Do(c *canvasapi.Canvas) (*models.Course, error) {
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
