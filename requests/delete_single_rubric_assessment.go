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

// DeleteSingleRubricAssessment Deletes a rubric assessment
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.RubricAssociationID (Required) ID
// # Path.ID (Required) ID
//
type DeleteSingleRubricAssessment struct {
	Path struct {
		CourseID            string `json:"course_id" url:"course_id,omitempty"`                         //  (Required)
		RubricAssociationID string `json:"rubric_association_id" url:"rubric_association_id,omitempty"` //  (Required)
		ID                  string `json:"id" url:"id,omitempty"`                                       //  (Required)
	} `json:"path"`
}

func (t *DeleteSingleRubricAssessment) GetMethod() string {
	return "DELETE"
}

func (t *DeleteSingleRubricAssessment) GetURLPath() string {
	path := "courses/{course_id}/rubric_associations/{rubric_association_id}/rubric_assessments/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{rubric_association_id}", fmt.Sprintf("%v", t.Path.RubricAssociationID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteSingleRubricAssessment) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteSingleRubricAssessment) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteSingleRubricAssessment) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteSingleRubricAssessment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.RubricAssociationID == "" {
		errs = append(errs, "'Path.RubricAssociationID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteSingleRubricAssessment) Do(c *canvasapi.Canvas) (*models.RubricAssessment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.RubricAssessment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
