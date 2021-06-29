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

// ShowRevisionCoursesRevisionID Retrieve the metadata and optionally content of a revision of the page.
// Note that retrieving historic versions of pages requires edit rights.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # CourseID (Required) ID
// # Url (Required) ID
// # RevisionID (Required) ID
//
// Query Parameters:
// # Summary (Optional) If set, exclude page content from results
//
type ShowRevisionCoursesRevisionID struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`     //  (Required)
		Url        string `json:"url" url:"url,omitempty"`                 //  (Required)
		RevisionID string `json:"revision_id" url:"revision_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Summary bool `json:"summary" url:"summary,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ShowRevisionCoursesRevisionID) GetMethod() string {
	return "GET"
}

func (t *ShowRevisionCoursesRevisionID) GetURLPath() string {
	path := "courses/{course_id}/pages/{url}/revisions/{revision_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	path = strings.ReplaceAll(path, "{revision_id}", fmt.Sprintf("%v", t.Path.RevisionID))
	return path
}

func (t *ShowRevisionCoursesRevisionID) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ShowRevisionCoursesRevisionID) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowRevisionCoursesRevisionID) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowRevisionCoursesRevisionID) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if t.Path.RevisionID == "" {
		errs = append(errs, "'RevisionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowRevisionCoursesRevisionID) Do(c *canvasapi.Canvas) (*models.PageRevision, error) {
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
