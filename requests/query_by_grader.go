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

// QueryByGrader List grade change events for a given grader.
// https://canvas.instructure.com/doc/api/grade_change_log.html
//
// Path Parameters:
// # Path.GraderID (Required) ID
//
// Query Parameters:
// # Query.StartTime (Optional) The beginning of the time range from which you want events.
// # Query.EndTime (Optional) The end of the time range from which you want events.
//
type QueryByGrader struct {
	Path struct {
		GraderID string `json:"grader_id" url:"grader_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *QueryByGrader) GetMethod() string {
	return "GET"
}

func (t *QueryByGrader) GetURLPath() string {
	path := "audit/grade_change/graders/{grader_id}"
	path = strings.ReplaceAll(path, "{grader_id}", fmt.Sprintf("%v", t.Path.GraderID))
	return path
}

func (t *QueryByGrader) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *QueryByGrader) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *QueryByGrader) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *QueryByGrader) HasErrors() error {
	errs := []string{}
	if t.Path.GraderID == "" {
		errs = append(errs, "'Path.GraderID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QueryByGrader) Do(c *canvasapi.Canvas) ([]*models.GradeChangeEvent, error) {
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
