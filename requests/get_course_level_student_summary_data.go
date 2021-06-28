package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetCourseLevelStudentSummaryData Returns a summary of per-user access information for all students in
// a course. This includes total page views, total participations, and a
// breakdown of on-time/late status for all homework submissions in the course.
//
// Each student's summary also includes the maximum number of page views and
// participations by any student in the course, which may be useful for some
// visualizations (since determining maximums client side can be tricky with
// pagination).
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # SortColumn (Optional) . Must be one of name, name_descending, score, score_descending, participations, participations_descending, page_views, page_views_descendingThe order results in which results are returned.  Defaults to "name".
// # StudentID (Optional) If set, returns only the specified student.
//
type GetCourseLevelStudentSummaryData struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SortColumn string `json:"sort_column" url:"sort_column,omitempty"` //  (Optional) . Must be one of name, name_descending, score, score_descending, participations, participations_descending, page_views, page_views_descending
		StudentID  string `json:"student_id" url:"student_id,omitempty"`   //  (Optional)
	} `json:"query"`
}

func (t *GetCourseLevelStudentSummaryData) GetMethod() string {
	return "GET"
}

func (t *GetCourseLevelStudentSummaryData) GetURLPath() string {
	path := "courses/{course_id}/analytics/student_summaries"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCourseLevelStudentSummaryData) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetCourseLevelStudentSummaryData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCourseLevelStudentSummaryData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCourseLevelStudentSummaryData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Query.SortColumn != "" && !string_utils.Include([]string{"name", "name_descending", "score", "score_descending", "participations", "participations_descending", "page_views", "page_views_descending"}, t.Query.SortColumn) {
		errs = append(errs, "SortColumn must be one of name, name_descending, score, score_descending, participations, participations_descending, page_views, page_views_descending")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseLevelStudentSummaryData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
