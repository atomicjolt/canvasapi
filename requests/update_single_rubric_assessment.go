package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateSingleRubricAssessment Returns the rubric assessment with the given id.
// The returned object also provides the information of
//   :ratings, :assessor_name, :related_group_submissions_and_assessments, :artifact
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # ID (Required) The id of the rubric assessment
// # CourseID (Required) The id of the course
// # RubricAssociationID (Required) The id of the object with which this rubric assessment is associated
//
// Form Parameters:
// # Provisional (Optional) (optional) Indicates whether this assessment is provisional, defaults to false.
// # Final (Optional) (optional) Indicates a provisional grade will be marked as final. It only takes effect if the provisional param is passed as true. Defaults to false.
// # GradedAnonymously (Optional) (optional) Defaults to false
// # RubricAssessment (Optional) A Hash of data to complement the rubric assessment:
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
type UpdateSingleRubricAssessment struct {
	Path struct {
		ID                  int64 `json:"id"`                    //  (Required)
		CourseID            int64 `json:"course_id"`             //  (Required)
		RubricAssociationID int64 `json:"rubric_association_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Provisional       string `json:"provisional"`        //  (Optional)
		Final             string `json:"final"`              //  (Optional)
		GradedAnonymously bool   `json:"graded_anonymously"` //  (Optional)
		RubricAssessment  string `json:"rubric_assessment"`  //  (Optional)
	} `json:"form"`
}

func (t *UpdateSingleRubricAssessment) GetMethod() string {
	return "PUT"
}

func (t *UpdateSingleRubricAssessment) GetURLPath() string {
	path := "courses/{course_id}/rubric_associations/{rubric_association_id}/rubric_assessments/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{rubric_association_id}", fmt.Sprintf("%v", t.Path.RubricAssociationID))
	return path
}

func (t *UpdateSingleRubricAssessment) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSingleRubricAssessment) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateSingleRubricAssessment) HasErrors() error {
	return nil
}

func (t *UpdateSingleRubricAssessment) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
