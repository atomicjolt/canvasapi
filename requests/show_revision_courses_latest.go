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

// ShowRevisionCoursesLatest Retrieve the metadata and optionally content of a revision of the page.
// Note that retrieving historic versions of pages requires edit rights.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # CourseID (Required) ID
// # Url (Required) ID
//
// Query Parameters:
// # Summary (Optional) If set, exclude page content from results
//
type ShowRevisionCoursesLatest struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		Url      string `json:"url" url:"url,omitempty"`             //  (Required)
	} `json:"path"`

	Query struct {
		Summary bool `json:"summary" url:"summary,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ShowRevisionCoursesLatest) GetMethod() string {
	return "GET"
}

func (t *ShowRevisionCoursesLatest) GetURLPath() string {
	path := "courses/{course_id}/pages/{url}/revisions/latest"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	return path
}

func (t *ShowRevisionCoursesLatest) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ShowRevisionCoursesLatest) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowRevisionCoursesLatest) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowRevisionCoursesLatest) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowRevisionCoursesLatest) Do(c *canvasapi.Canvas) (*models.PageRevision, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PageRevision{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
