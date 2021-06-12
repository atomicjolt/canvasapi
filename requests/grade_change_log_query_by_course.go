package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GradeChangeLogQueryByCourse List grade change events for a given course.
// https://canvas.instructure.com/doc/api/grade_change_log.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want events.
// # EndTime (Optional) The end of the time range from which you want events.
//
type GradeChangeLogQueryByCourse struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time"` //  (Optional)
		EndTime   time.Time `json:"end_time"`   //  (Optional)
	} `json:"query"`
}

func (t *GradeChangeLogQueryByCourse) GetMethod() string {
	return "GET"
}

func (t *GradeChangeLogQueryByCourse) GetURLPath() string {
	path := "audit/grade_change/courses/{course_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GradeChangeLogQueryByCourse) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GradeChangeLogQueryByCourse) GetBody() (string, error) {
	return "", nil
}

func (t *GradeChangeLogQueryByCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GradeChangeLogQueryByCourse) Do(c *canvasapi.Canvas) ([]*models.GradeChangeEvent, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.GradeChangeEvent{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
