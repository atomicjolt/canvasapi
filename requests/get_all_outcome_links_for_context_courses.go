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

// GetAllOutcomeLinksForContextCourses
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # OutcomeStyle (Optional) The detail level of the outcomes. Defaults to "abbrev".
//    Specify "full" for more information.
// # OutcomeGroupStyle (Optional) The detail level of the outcome groups. Defaults to "abbrev".
//    Specify "full" for more information.
//
type GetAllOutcomeLinksForContextCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		OutcomeStyle      string `json:"outcome_style" url:"outcome_style,omitempty"`             //  (Optional)
		OutcomeGroupStyle string `json:"outcome_group_style" url:"outcome_group_style,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetAllOutcomeLinksForContextCourses) GetMethod() string {
	return "GET"
}

func (t *GetAllOutcomeLinksForContextCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_group_links"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetAllOutcomeLinksForContextCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetAllOutcomeLinksForContextCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAllOutcomeLinksForContextCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAllOutcomeLinksForContextCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetAllOutcomeLinksForContextCourses) Do(c *canvasapi.Canvas) ([]*models.OutcomeLink, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.OutcomeLink{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
