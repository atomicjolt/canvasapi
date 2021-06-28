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

// BulkUpdateColumnData Set the content of custom columns
//
// {
//   "column_data": [
//     {
//       "column_id": example_column_id,
//       "user_id": example_student_id,
//       "content": example_content
//       },
//       {
//       "column_id": example_column_id,
//       "user_id": example_student_id,
//       "content: example_content
//     }
//   ]
// }
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # ColumnData (Required) Column content. Setting this to an empty string will delete the data object.
//
type BulkUpdateColumnData struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ColumnData []string `json:"column_data" url:"column_data,omitempty"` //  (Required)
	} `json:"form"`
}

func (t *BulkUpdateColumnData) GetMethod() string {
	return "PUT"
}

func (t *BulkUpdateColumnData) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_column_data"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *BulkUpdateColumnData) GetQuery() (string, error) {
	return "", nil
}

func (t *BulkUpdateColumnData) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *BulkUpdateColumnData) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *BulkUpdateColumnData) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.ColumnData == nil {
		errs = append(errs, "'ColumnData' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BulkUpdateColumnData) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
