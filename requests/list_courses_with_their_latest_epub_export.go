package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListCoursesWithTheirLatestEpubExport A paginated list of all courses a user is actively participating in, and
// the latest ePub export associated with the user & course.
// https://canvas.instructure.com/doc/api/e_pub_exports.html
//
type ListCoursesWithTheirLatestEpubExport struct {
}

func (t *ListCoursesWithTheirLatestEpubExport) GetMethod() string {
	return "GET"
}

func (t *ListCoursesWithTheirLatestEpubExport) GetURLPath() string {
	return ""
}

func (t *ListCoursesWithTheirLatestEpubExport) GetQuery() (string, error) {
	return "", nil
}

func (t *ListCoursesWithTheirLatestEpubExport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCoursesWithTheirLatestEpubExport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCoursesWithTheirLatestEpubExport) HasErrors() error {
	return nil
}

func (t *ListCoursesWithTheirLatestEpubExport) Do(c *canvasapi.Canvas) ([]*models.CourseEpubExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CourseEpubExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
