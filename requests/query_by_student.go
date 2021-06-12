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

// QueryByStudent List grade change events for a given student.
// https://canvas.instructure.com/doc/api/grade_change_log.html
//
// Path Parameters:
// # StudentID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want events.
// # EndTime (Optional) The end of the time range from which you want events.
//
type QueryByStudent struct {
	Path struct {
		StudentID string `json:"student_id"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time"` //  (Optional)
		EndTime   time.Time `json:"end_time"`   //  (Optional)
	} `json:"query"`
}

func (t *QueryByStudent) GetMethod() string {
	return "GET"
}

func (t *QueryByStudent) GetURLPath() string {
	path := "audit/grade_change/students/{student_id}"
	path = strings.ReplaceAll(path, "{student_id}", fmt.Sprintf("%v", t.Path.StudentID))
	return path
}

func (t *QueryByStudent) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *QueryByStudent) GetBody() (string, error) {
	return "", nil
}

func (t *QueryByStudent) HasErrors() error {
	errs := []string{}
	if t.Path.StudentID == "" {
		errs = append(errs, "'StudentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QueryByStudent) Do(c *canvasapi.Canvas) ([]*models.GradeChangeEvent, error) {
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
