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

// CreateAssignmentOverride One of student_ids, group_id, or course_section_id must be present. At most
// one should be present; if multiple are present only the most specific
// (student_ids first, then group_id, then course_section_id) is used and any
// others are ignored.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
// # AssignmentID (Required) ID
//
// Form Parameters:
// # AssignmentOverride (Optional) The IDs of
//    the override's target students. If present, the IDs must each identify a
//    user with an active student enrollment in the course that is not already
//    targetted by a different adhoc override.
// # AssignmentOverride (Optional) The title of the adhoc
//    assignment override. Required if student_ids is present, ignored
//    otherwise (the title is set to the name of the targetted group or section
//    instead).
// # AssignmentOverride (Optional) The ID of the
//    override's target group. If present, the following conditions must be met
//    for the override to be successful:
//
//    1. the assignment MUST be a group assignment (a group_category_id is assigned to it)
//    2. the ID must identify an active group in the group set the assignment is in
//    3. the ID must not be targetted by a different override
//
//    See {Appendix: Group assignments} for more info.
// # AssignmentOverride (Optional) The ID
//    of the override's target section. If present, must identify an active
//    section of the assignment's course not already targetted by a different
//    override.
// # AssignmentOverride (Optional) The day/time
//    the overridden assignment is due. Accepts times in ISO 8601 format, e.g.
//    2014-10-21T18:48:00Z. If absent, this override will not affect due date.
//    May be present but null to indicate the override removes any previous due
//    date.
// # AssignmentOverride (Optional) The day/time
//    the overridden assignment becomes unlocked. Accepts times in ISO 8601
//    format, e.g. 2014-10-21T18:48:00Z. If absent, this override will not
//    affect the unlock date. May be present but null to indicate the override
//    removes any previous unlock date.
// # AssignmentOverride (Optional) The day/time
//    the overridden assignment becomes locked. Accepts times in ISO 8601
//    format, e.g. 2014-10-21T18:48:00Z. If absent, this override will not
//    affect the lock date. May be present but null to indicate the override
//    removes any previous lock date.
//
type CreateAssignmentOverride struct {
	Path struct {
		CourseID     string `json:"course_id"`     //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`

	Form struct {
		AssignmentOverride struct {
			StudentIDs      []int64   `json:"student_ids"`       //  (Optional)
			Title           string    `json:"title"`             //  (Optional)
			GroupID         int64     `json:"group_id"`          //  (Optional)
			CourseSectionID int64     `json:"course_section_id"` //  (Optional)
			DueAt           time.Time `json:"due_at"`            //  (Optional)
			UnlockAt        time.Time `json:"unlock_at"`         //  (Optional)
			LockAt          time.Time `json:"lock_at"`           //  (Optional)
		} `json:"assignment_override"`
	} `json:"form"`
}

func (t *CreateAssignmentOverride) GetMethod() string {
	return "POST"
}

func (t *CreateAssignmentOverride) GetURLPath() string {
	path := "courses/{course_id}/assignments/{assignment_id}/overrides"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *CreateAssignmentOverride) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateAssignmentOverride) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateAssignmentOverride) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateAssignmentOverride) Do(c *canvasapi.Canvas) (*models.AssignmentOverride, error) {
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
