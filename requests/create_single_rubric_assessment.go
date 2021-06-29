package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateSingleRubricAssessment Returns the rubric assessment with the given id.
// The returned object also provides the information of
//   :ratings, :assessor_name, :related_group_submissions_and_assessments, :artifact
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # Path.CourseID (Required) The id of the course
// # Path.RubricAssociationID (Required) The id of the object with which this rubric assessment is associated
//
// Form Parameters:
// # Form.Provisional (Optional) (optional) Indicates whether this assessment is provisional, defaults to false.
// # Form.Final (Optional) (optional) Indicates a provisional grade will be marked as final. It only takes effect if the provisional param is passed as true. Defaults to false.
// # Form.GradedAnonymously (Optional) (optional) Defaults to false
// # Form.RubricAssessment (Optional) A Hash of data to complement the rubric assessment:
//    The user id that refers to the person being assessed
//      rubric_assessment[user_id]
//    Assessment type. There are only three valid types:  'grading', 'peer_review', or 'provisional_grade'
//      rubric_assessment[assessment_type]
//    The points awarded for this row.
//      rubric_assessment[criterion_id][points]
//    Comments to add for this row.
//      rubric_assessment[criterion_id][comments]
//    For each criterion_id, change the id by the criterion number, ex: criterion_123
//    If the criterion_id is not specified it defaults to false, and nothing is updated.
//
type CreateSingleRubricAssessment struct {
	Path struct {
		CourseID            int64 `json:"course_id" url:"course_id,omitempty"`                         //  (Required)
		RubricAssociationID int64 `json:"rubric_association_id" url:"rubric_association_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Provisional       string                   `json:"provisional" url:"provisional,omitempty"`               //  (Optional)
		Final             string                   `json:"final" url:"final,omitempty"`                           //  (Optional)
		GradedAnonymously bool                     `json:"graded_anonymously" url:"graded_anonymously,omitempty"` //  (Optional)
		RubricAssessment  map[string](interface{}) `json:"rubric_assessment" url:"rubric_assessment,omitempty"`   //  (Optional)
	} `json:"form"`
}

func (t *CreateSingleRubricAssessment) GetMethod() string {
	return "POST"
}

func (t *CreateSingleRubricAssessment) GetURLPath() string {
	path := "courses/{course_id}/rubric_associations/{rubric_association_id}/rubric_assessments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{rubric_association_id}", fmt.Sprintf("%v", t.Path.RubricAssociationID))
	return path
}

func (t *CreateSingleRubricAssessment) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSingleRubricAssessment) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateSingleRubricAssessment) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateSingleRubricAssessment) HasErrors() error {
	return nil
}

func (t *CreateSingleRubricAssessment) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
