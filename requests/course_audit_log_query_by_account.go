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

// CourseAuditLogQueryByAccount List course change events for a given account.
// https://canvas.instructure.com/doc/api/course_audit_log.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.StartTime (Optional) The beginning of the time range from which you want events.
// # Query.EndTime (Optional) The end of the time range from which you want events.
//
type CourseAuditLogQueryByAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *CourseAuditLogQueryByAccount) GetMethod() string {
	return "GET"
}

func (t *CourseAuditLogQueryByAccount) GetURLPath() string {
	path := "audit/course/accounts/{account_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CourseAuditLogQueryByAccount) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *CourseAuditLogQueryByAccount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *CourseAuditLogQueryByAccount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *CourseAuditLogQueryByAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CourseAuditLogQueryByAccount) Do(c *canvasapi.Canvas) ([]*models.CourseEvent, error) {
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
