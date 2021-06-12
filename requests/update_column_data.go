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

// UpdateColumnData Set the content of a custom column
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
// # UserID (Required) ID
//
// Form Parameters:
// # ColumnData (Required) Column content.  Setting this to blank will delete the datum object.
//
type UpdateColumnData struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
		UserID   string `json:"user_id"`   //  (Required)
	} `json:"path"`

	Form struct {
		ColumnData struct {
			Content string `json:"content"` //  (Required)
		} `json:"column_data"`
	} `json:"form"`
}

func (t *UpdateColumnData) GetMethod() string {
	return "PUT"
}

func (t *UpdateColumnData) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns/{id}/data/{user_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *UpdateColumnData) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateColumnData) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateColumnData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Form.ColumnData.Content == "" {
		errs = append(errs, "'ColumnData' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateColumnData) Do(c *canvasapi.Canvas) (*models.ColumnDatum, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ColumnDatum{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
