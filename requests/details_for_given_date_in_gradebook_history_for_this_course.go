package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DetailsForGivenDateInGradebookHistoryForThisCourse Returns the graders who worked on this day, along with the assignments they worked on.
// More details can be obtained by selecting a grader and assignment and calling the
// 'submissions' api endpoint for a given date.
// https://canvas.instructure.com/doc/api/gradebook_history.html
//
// Path Parameters:
// # Path.CourseID (Required) The id of the contextual course for this API call
// # Path.Date (Required) The date for which you would like to see detailed information
//
type DetailsForGivenDateInGradebookHistoryForThisCourse struct {
	Path struct {
		CourseID int64  `json:"course_id" url:"course_id,omitempty"` //  (Required)
		Date     string `json:"date" url:"date,omitempty"`           //  (Required)
	} `json:"path"`
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) GetMethod() string {
	return "GET"
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) GetURLPath() string {
	path := "courses/{course_id}/gradebook_history/{date}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{date}", fmt.Sprintf("%v", t.Path.Date))
	return path
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) GetQuery() (string, error) {
	return "", nil
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) HasErrors() error {
	errs := []string{}
	if t.Path.Date == "" {
		errs = append(errs, "'Path.Date' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) Do(c *canvasapi.Canvas) ([]*models.Grader, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Grader{}
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
