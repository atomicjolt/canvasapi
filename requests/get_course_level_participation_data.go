package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCourseLevelParticipationData Returns page view hits and participation numbers grouped by day through the
// entire history of the course. Page views is returned as a hash, where the
// hash keys are dates in the format "YYYY-MM-DD". The page_views result set
// includes page views broken out by access category. Participations is
// returned as an array of dates in the format "YYYY-MM-DD".
// https://canvas.instructure.com/doc/api/analytics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
type GetCourseLevelParticipationData struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetCourseLevelParticipationData) GetMethod() string {
	return "GET"
}

func (t *GetCourseLevelParticipationData) GetURLPath() string {
	path := "courses/{course_id}/analytics/activity"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetCourseLevelParticipationData) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCourseLevelParticipationData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCourseLevelParticipationData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCourseLevelParticipationData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCourseLevelParticipationData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
