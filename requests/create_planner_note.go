package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreatePlannerNote Create a planner note for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Form Parameters:
// # Form.Title (Optional) The title of the planner note.
// # Form.Details (Optional) Text of the planner note.
// # Form.TodoDate (Optional) The date where this planner note should appear in the planner.
//    The value should be formatted as: yyyy-mm-dd.
// # Form.CourseID (Optional) The ID of the course to associate with the planner note. The caller must be able to view the course in order to
//    associate it with a planner note.
// # Form.LinkedObjectType (Optional) The type of a learning object to link to this planner note. Must be used in conjunction wtih linked_object_id
//    and course_id. Valid linked_object_type values are:
//    'announcement', 'assignment', 'discussion_topic', 'wiki_page', 'quiz'
// # Form.LinkedObjectID (Optional) The id of a learning object to link to this planner note. Must be used in conjunction with linked_object_type
//    and course_id. The object must be in the same course as specified by course_id. If the title argument is not
//    provided, the planner note will use the learning object's title as its title. Only one planner note may be
//    linked to a specific learning object.
//
type CreatePlannerNote struct {
	Form struct {
		Title            string    `json:"title" url:"title,omitempty"`                           //  (Optional)
		Details          string    `json:"details" url:"details,omitempty"`                       //  (Optional)
		TodoDate         time.Time `json:"todo_date" url:"todo_date,omitempty"`                   //  (Optional)
		CourseID         int64     `json:"course_id" url:"course_id,omitempty"`                   //  (Optional)
		LinkedObjectType string    `json:"linked_object_type" url:"linked_object_type,omitempty"` //  (Optional)
		LinkedObjectID   int64     `json:"linked_object_id" url:"linked_object_id,omitempty"`     //  (Optional)
	} `json:"form"`
}

func (t *CreatePlannerNote) GetMethod() string {
	return "POST"
}

func (t *CreatePlannerNote) GetURLPath() string {
	return ""
}

func (t *CreatePlannerNote) GetQuery() (string, error) {
	return "", nil
}

func (t *CreatePlannerNote) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreatePlannerNote) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreatePlannerNote) HasErrors() error {
	return nil
}

func (t *CreatePlannerNote) Do(c *canvasapi.Canvas) (*models.PlannerNote, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PlannerNote{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
