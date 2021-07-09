package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListEntriesForColumn This does not list entries for students without associated data.
// https://canvas.instructure.com/doc/api/custom_gradebook_columns.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.IncludeHidden (Optional) If true, hidden columns will be included in the
//    result. If false or absent, only visible columns
//    will be returned.
//
type ListEntriesForColumn struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		IncludeHidden bool `json:"include_hidden" url:"include_hidden,omitempty"` //  (Optional)
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
	return v.Encode(), nil
}

func (t *ListEntriesForColumn) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListEntriesForColumn) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListEntriesForColumn) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListEntriesForColumn) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.ColumnDatum, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.ColumnDatum{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
