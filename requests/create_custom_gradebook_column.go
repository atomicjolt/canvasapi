package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateCustomGradebookColumn Create a custom gradebook column
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Column (Required) no description
// # Column (Optional) The position of the column relative to other custom columns
// # Column (Optional) Hidden columns are not displayed in the gradebook
// # Column (Optional) Set this if the column is created by a teacher.  The gradebook only
//    supports one teacher_notes column.
// # Column (Optional) Set this to prevent the column from being editable in the gradebook ui
//
type CreateCustomGradebookColumn struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Column struct {
			Title        string `json:"title"`         //  (Required)
			Position     int64  `json:"position"`      //  (Optional)
			Hidden       bool   `json:"hidden"`        //  (Optional)
			TeacherNotes bool   `json:"teacher_notes"` //  (Optional)
			ReadOnly     bool   `json:"read_only"`     //  (Optional)
		} `json:"column"`
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

func (t *CreateCustomGradebookColumn) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateCustomGradebookColumn) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Column.Title == "" {
		errs = append(errs, "'Column' is required")
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
