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

// UpdateColumnData Set the content of a custom column
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
// # Path.UserID (Required) ID
//
// Form Parameters:
// # Form.ColumnData.Content (Required) Column content.  Setting this to blank will delete the datum object.
//
type UpdateColumnData struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
		UserID   string `json:"user_id" url:"user_id,omitempty"`     //  (Required)
	} `json:"path"`

	Form struct {
		ColumnData struct {
			Content string `json:"content" url:"content,omitempty"` //  (Required)
		} `json:"column_data" url:"column_data,omitempty"`
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

func (t *UpdateColumnData) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateColumnData) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateColumnData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Form.ColumnData.Content == "" {
		errs = append(errs, "'Form.ColumnData.Content' is required")
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
