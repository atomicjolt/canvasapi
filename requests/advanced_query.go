package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// AdvancedQuery List grade change events satisfying all given parameters. Teachers may query for events in courses they teach.
// Queries without +course_id+ require account administrator rights.
//
// At least one of +course_id+, +assignment_id+, +student_id+, or +grader_id+ must be specified.
// https://canvas.instructure.com/doc/api/grade_change_log.html
//
// Query Parameters:
// # CourseID (Optional) Restrict query to events in the specified course.
// # AssignmentID (Optional) Restrict query to the given assignment. If "override" is given, query the course final grade override instead.
// # StudentID (Optional) User id of a student to search grading events for.
// # GraderID (Optional) User id of a grader to search grading events for.
// # StartTime (Optional) The beginning of the time range from which you want events.
// # EndTime (Optional) The end of the time range from which you want events.
//
type AdvancedQuery struct {
	Query struct {
		CourseID     int64     `json:"course_id"`     //  (Optional)
		AssignmentID int64     `json:"assignment_id"` //  (Optional)
		StudentID    int64     `json:"student_id"`    //  (Optional)
		GraderID     int64     `json:"grader_id"`     //  (Optional)
		StartTime    time.Time `json:"start_time"`    //  (Optional)
		EndTime      time.Time `json:"end_time"`      //  (Optional)
	} `json:"query"`
}

func (t *AdvancedQuery) GetMethod() string {
	return "GET"
}

func (t *AdvancedQuery) GetURLPath() string {
	return ""
}

func (t *AdvancedQuery) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *AdvancedQuery) GetBody() (string, error) {
	return "", nil
}

func (t *AdvancedQuery) HasErrors() error {
	return nil
}

func (t *AdvancedQuery) Do(c *canvasapi.Canvas) ([]*models.GradeChangeEvent, error) {
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
