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

// UpdateAssignmentOverride All current overridden values must be supplied if they are to be retained;
// e.g. if due_at was overridden, but this PUT omits a value for due_at,
// due_at will no longer be overridden. If the override is adhoc and
// student_ids is not supplied, the target override set is unchanged. Target
// override sets cannot be changed for group or section overrides.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.AssignmentOverride.StudentIDs (Optional) The IDs of the
//    override's target students. If present, the IDs must each identify a
//    user with an active student enrollment in the course that is not already
//    targetted by a different adhoc override. Ignored unless the override
//    being updated is adhoc.
// # Form.AssignmentOverride.Title (Optional) The title of an adhoc
//    assignment override. Ignored unless the override being updated is adhoc.
// # Form.AssignmentOverride.DueAt (Optional) The day/time
//    the overridden assignment is due. Accepts times in ISO 8601 format, e.g.
//    2014-10-21T18:48:00Z. If absent, this override will not affect due date.
//    May be present but null to indicate the override removes any previous due
//    date.
// # Form.AssignmentOverride.UnlockAt (Optional) The day/time
//    the overridden assignment becomes unlocked. Accepts times in ISO 8601
//    format, e.g. 2014-10-21T18:48:00Z. If absent, this override will not
//    affect the unlock date. May be present but null to indicate the override
//    removes any previous unlock date.
// # Form.AssignmentOverride.LockAt (Optional) The day/time
//    the overridden assignment becomes locked. Accepts times in ISO 8601
//    format, e.g. 2014-10-21T18:48:00Z. If absent, this override will not
//    affect the lock date. May be present but null to indicate the override
//    removes any previous lock date.
//
type UpdateAssignmentOverride struct {
	Path struct {
		CourseID     string `json:"course_id" url:"course_id,omitempty"`         //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		ID           string `json:"id" url:"id,omitempty"`                       //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentOverride struct {
			StudentIDs []string  `json:"student_ids" url:"student_ids,omitempty"` //  (Optional)
			Title      string    `json:"title" url:"title,omitempty"`             //  (Optional)
			DueAt      time.Time `json:"due_at" url:"due_at,omitempty"`           //  (Optional)
			UnlockAt   time.Time `json:"unlock_at" url:"unlock_at,omitempty"`     //  (Optional)
			LockAt     time.Time `json:"lock_at" url:"lock_at,omitempty"`         //  (Optional)
		} `json:"assignment_override" url:"assignment_override,omitempty"`
	} `json:"form"`
}

func (t *UpdateAssignmentOverride) GetMethod() string {
	return "PUT"
}

func (t *UpdateAssignmentOverride) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/overrides/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateAssignmentOverride) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateAssignmentOverride) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateAssignmentOverride) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateAssignmentOverride) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateAssignmentOverride) Do(c *canvasapi.Canvas) (*models.AssignmentOverride, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.AssignmentOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
