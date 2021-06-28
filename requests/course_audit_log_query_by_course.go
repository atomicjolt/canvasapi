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
)

// CourseAuditLogQueryByCourse List course change events for a given course.
// https://canvas.instructure.com/doc/api/course_audit_log.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want events.
// # EndTime (Optional) The end of the time range from which you want events.
//
type CourseAuditLogQueryByCourse struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *CourseAuditLogQueryByCourse) GetMethod() string {
	return "GET"
}

func (t *CourseAuditLogQueryByCourse) GetURLPath() string {
	path := "audit/course/courses/{course_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CourseAuditLogQueryByCourse) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *CourseAuditLogQueryByCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CourseAuditLogQueryByCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CourseAuditLogQueryByCourse) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CourseAuditLogQueryByCourse) Do(c *canvasapi.Canvas) ([]*models.CourseEvent, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CourseEvent{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
