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

// ListEntriesForColumn This does not list entries for students without associated data.
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # IncludeHidden (Optional) If true, hidden columns will be included in the
//    result. If false or absent, only visible columns
//    will be returned.
//
type ListEntriesForColumn struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Query struct {
		IncludeHidden bool `json:"include_hidden"` //  (Optional)
	} `json:"query"`
}

func (t *ListEntriesForColumn) GetMethod() string {
	return "GET"
}

func (t *ListEntriesForColumn) GetURLPath() string {
	path := "courses/{course_id}/custom_gradebook_columns/{id}/data"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListEntriesForColumn) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListEntriesForColumn) GetBody() (string, error) {
	return "", nil
}

func (t *ListEntriesForColumn) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEntriesForColumn) Do(c *canvasapi.Canvas) ([]*models.ColumnDatum, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.ColumnDatum{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
