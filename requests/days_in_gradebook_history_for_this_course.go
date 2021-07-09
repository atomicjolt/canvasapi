package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DaysInGradebookHistoryForThisCourse Returns a map of dates to grader/assignment groups
// https://canvas.instructure.com/doc/api/gradebook_history.html
//
// Path Parameters:
// # Path.CourseID (Required) The id of the contextual course for this API call
//
type DaysInGradebookHistoryForThisCourse struct {
	Path struct {
		CourseID int64 `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DaysInGradebookHistoryForThisCourse) GetMethod() string {
	return "GET"
}

func (t *DaysInGradebookHistoryForThisCourse) GetURLPath() string {
	path := "courses/{course_id}/gradebook_history/days"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *DaysInGradebookHistoryForThisCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *DaysInGradebookHistoryForThisCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DaysInGradebookHistoryForThisCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DaysInGradebookHistoryForThisCourse) HasErrors() error {
	return nil
}

func (t *DaysInGradebookHistoryForThisCourse) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Day, *canvasapi.PagedResource, error) {
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
	ret := []*models.Day{}
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
