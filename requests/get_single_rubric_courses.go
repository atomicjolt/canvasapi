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
	"github.com/atomicjolt/string_utils"
)

// GetSingleRubricCourses Returns the rubric with the given id.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of assessments, graded_assessments, peer_assessments, associations, assignment_associations, course_associations, account_associationsRelated records to include in the response.
// # Style (Optional) . Must be one of full, comments_onlyApplicable only if assessments are being returned. If included, returns either all criteria data associated with the assessment, or just the comments. If not included, both data and comments are omitted.
//
type GetSingleRubricCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of assessments, graded_assessments, peer_assessments, associations, assignment_associations, course_associations, account_associations
		Style   string   `json:"style" url:"style,omitempty"`     //  (Optional) . Must be one of full, comments_only
	} `json:"query"`
}

func (t *GetSingleRubricCourses) GetMethod() string {
	return "GET"
}

func (t *GetSingleRubricCourses) GetURLPath() string {
	path := "courses/{course_id}/rubrics/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleRubricCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleRubricCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleRubricCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleRubricCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"assessments", "graded_assessments", "peer_assessments", "associations", "assignment_associations", "course_associations", "account_associations"}, v) {
			errs = append(errs, "Include must be one of assessments, graded_assessments, peer_assessments, associations, assignment_associations, course_associations, account_associations")
		}
	}
	if t.Query.Style != "" && !string_utils.Include([]string{"full", "comments_only"}, t.Query.Style) {
		errs = append(errs, "Style must be one of full, comments_only")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleRubricCourses) Do(c *canvasapi.Canvas) (*models.Rubric, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Rubric{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
