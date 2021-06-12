package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # CourseID (Required) The id of the contextual course for this API call
// # Date (Required) The date for which you would like to see detailed information
//
type DetailsForGivenDateInGradebookHistoryForThisCourse struct {
	Path struct {
		CourseID int64  `json:"course_id"` //  (Required)
		Date     string `json:"date"`      //  (Required)
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

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) GetBody() (string, error) {
	return "", nil
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) HasErrors() error {
	errs := []string{}
	if t.Path.Date == "" {
		errs = append(errs, "'Date' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DetailsForGivenDateInGradebookHistoryForThisCourse) Do(c *canvasapi.Canvas) ([]*models.Grader, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Grader{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
