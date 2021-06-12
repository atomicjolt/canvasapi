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

// ListCustomGradebookColumns A paginated list of all custom gradebook columns for a course
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # IncludeHidden (Optional) Include hidden parameters (defaults to false)
//
type ListCustomGradebookColumns struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		IncludeHidden bool `json:"include_hidden"` //  (Optional)
	} `json:"query"`
}

func (t *ListCustomGradebookColumns) GetMethod() string {
	return "GET"
}

func (t *ListCustomGradebookColumns) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListCustomGradebookColumns) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListCustomGradebookColumns) GetBody() (string, error) {
	return "", nil
}

func (t *ListCustomGradebookColumns) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCustomGradebookColumns) Do(c *canvasapi.Canvas) ([]*models.CustomColumn, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CustomColumn{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}