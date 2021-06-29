package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateCustomGradebookColumn Create a custom gradebook column
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.Column.Title (Required) no description
// # Form.Column.Position (Optional) The position of the column relative to other custom columns
// # Form.Column.Hidden (Optional) Hidden columns are not displayed in the gradebook
// # Form.Column.TeacherNotes (Optional) Set this if the column is created by a teacher.  The gradebook only
//    supports one teacher_notes column.
// # Form.Column.ReadOnly (Optional) Set this to prevent the column from being editable in the gradebook ui
//
type CreateCustomGradebookColumn struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Column struct {
			Title        string `json:"title" url:"title,omitempty"`                 //  (Required)
			Position     int64  `json:"position" url:"position,omitempty"`           //  (Optional)
			Hidden       bool   `json:"hidden" url:"hidden,omitempty"`               //  (Optional)
			TeacherNotes bool   `json:"teacher_notes" url:"teacher_notes,omitempty"` //  (Optional)
			ReadOnly     bool   `json:"read_only" url:"read_only,omitempty"`         //  (Optional)
		} `json:"column" url:"column,omitempty"`
	} `json:"form"`
}

func (t *CreateCustomGradebookColumn) GetMethod() string {
	return "POST"
}

func (t *CreateCustomGradebookColumn) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateCustomGradebookColumn) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateCustomGradebookColumn) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateCustomGradebookColumn) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateCustomGradebookColumn) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.Column.Title == "" {
		errs = append(errs, "'Form.Column.Title' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateCustomGradebookColumn) Do(c *canvasapi.Canvas) (*models.CustomColumn, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.CustomColumn{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
