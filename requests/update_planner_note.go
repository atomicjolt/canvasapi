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

// UpdatePlannerNote Update a planner note for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # Title (Optional) The title of the planner note.
// # Details (Optional) Text of the planner note.
// # TodoDate (Optional) The date where this planner note should appear in the planner.
//    The value should be formatted as: yyyy-mm-dd.
// # CourseID (Optional) The ID of the course to associate with the planner note. The caller must be able to view the course in order to
//    associate it with a planner note. Use a null or empty value to remove a planner note from a course. Note that if
//    the planner note is linked to a learning object, its course_id cannot be changed.
//
type UpdatePlannerNote struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Title    string    `json:"title" url:"title,omitempty"`         //  (Optional)
		Details  string    `json:"details" url:"details,omitempty"`     //  (Optional)
		TodoDate time.Time `json:"todo_date" url:"todo_date,omitempty"` //  (Optional)
		CourseID int64     `json:"course_id" url:"course_id,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *UpdatePlannerNote) GetMethod() string {
	return "PUT"
}

func (t *UpdatePlannerNote) GetURLPath() string {
	path := "planner_notes/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdatePlannerNote) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdatePlannerNote) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdatePlannerNote) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdatePlannerNote) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdatePlannerNote) Do(c *canvasapi.Canvas) (*models.PlannerNote, error) {
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
