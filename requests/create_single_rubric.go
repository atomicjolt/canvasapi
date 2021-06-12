package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateSingleRubric Returns the rubric with the given id.
//
// Unfortuantely this endpoint does not return a standard Rubric object,
// instead it returns a hash that looks like
//   { 'rubric': Rubric, 'rubric_association': RubricAssociation }
//
// This may eventually be deprecated in favor of a more standardized return
// value, but that is not currently planned.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # ID (Optional) The id of the rubric
// # RubricAssociationID (Optional) The id of the object with which this rubric is associated
// # Rubric (Optional) The title of the rubric
// # Rubric (Optional) Whether or not you can write custom comments in the ratings field for a rubric
// # RubricAssociation (Optional) The id of the object with which this rubric is associated
// # RubricAssociation (Optional) . Must be one of Assignment, Course, AccountThe type of object this rubric is associated with
// # RubricAssociation (Optional) Whether or not the associated rubric is used for grade calculation
// # RubricAssociation (Optional) Whether or not the score total is displayed within the rubric.
//    This option is only available if the rubric is not used for grading.
// # RubricAssociation (Optional) Whether or not the association is for grading (and thus linked to an assignment)
//    or if it's to indicate the rubric should appear in its context
// # Rubric (Optional) An indexed Hash of RubricCriteria objects where the keys are integer ids and the values are the RubricCriteria objects
//
type CreateSingleRubric struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		ID                  int64 `json:"id"`                    //  (Optional)
		RubricAssociationID int64 `json:"rubric_association_id"` //  (Optional)
		Rubric              struct {
			Title                     string `json:"title"`                        //  (Optional)
			FreeFormCriterionComments bool   `json:"free_form_criterion_comments"` //  (Optional)
			Criteria                  string `json:"criteria"`                     //  (Optional)
		} `json:"rubric"`

		RubricAssociation struct {
			AssociationID   int64  `json:"association_id"`   //  (Optional)
			AssociationType string `json:"association_type"` //  (Optional) . Must be one of Assignment, Course, Account
			UseForGrading   bool   `json:"use_for_grading"`  //  (Optional)
			HideScoreTotal  bool   `json:"hide_score_total"` //  (Optional)
			Purpose         string `json:"purpose"`          //  (Optional)
		} `json:"rubric_association"`
	} `json:"form"`
}

func (t *CreateSingleRubric) GetMethod() string {
	return "POST"
}

func (t *CreateSingleRubric) GetURLPath() string {
	path := "courses/{course_id}/rubrics"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateSingleRubric) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSingleRubric) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateSingleRubric) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"Assignment", "Course", "Account"}, t.Form.RubricAssociation.AssociationType) {
		errs = append(errs, "RubricAssociation must be one of Assignment, Course, Account")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSingleRubric) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
