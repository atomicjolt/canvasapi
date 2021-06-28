package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// UpdateSingleRubric Returns the rubric with the given id.
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
// # ID (Required) The id of the rubric
//
// Form Parameters:
// # RubricAssociationID (Optional) The id of the object with which this rubric is associated
// # Rubric (Optional) The title of the rubric
// # Rubric (Optional) Whether or not you can write custom comments in the ratings field for a rubric
// # Rubric (Optional) Whether or not to update the points possible
// # RubricAssociation (Optional) The id of the object with which this rubric is associated
// # RubricAssociation (Optional) . Must be one of Assignment, Course, AccountThe type of object this rubric is associated with
// # RubricAssociation (Optional) Whether or not the associated rubric is used for grade calculation
// # RubricAssociation (Optional) Whether or not the score total is displayed within the rubric.
//    This option is only available if the rubric is not used for grading.
// # RubricAssociation (Optional) . Must be one of grading, bookmarkWhether or not the association is for grading (and thus linked to an assignment)
//    or if it's to indicate the rubric should appear in its context
// # Rubric (Optional) An indexed Hash of RubricCriteria objects where the keys are integer ids and the values are the RubricCriteria objects
//
type UpdateSingleRubric struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       int64  `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		RubricAssociationID int64 `json:"rubric_association_id" url:"rubric_association_id,omitempty"` //  (Optional)
		Rubric              struct {
			Title                      string `json:"title" url:"title,omitempty"`                                                 //  (Optional)
			FreeFormCriterionComments  bool   `json:"free_form_criterion_comments" url:"free_form_criterion_comments,omitempty"`   //  (Optional)
			SkipUpdatingPointsPossible bool   `json:"skip_updating_points_possible" url:"skip_updating_points_possible,omitempty"` //  (Optional)
			Criteria                   string `json:"criteria" url:"criteria,omitempty"`                                           //  (Optional)
		} `json:"rubric" url:"rubric,omitempty"`

		RubricAssociation struct {
			AssociationID   int64  `json:"association_id" url:"association_id,omitempty"`     //  (Optional)
			AssociationType string `json:"association_type" url:"association_type,omitempty"` //  (Optional) . Must be one of Assignment, Course, Account
			UseForGrading   bool   `json:"use_for_grading" url:"use_for_grading,omitempty"`   //  (Optional)
			HideScoreTotal  bool   `json:"hide_score_total" url:"hide_score_total,omitempty"` //  (Optional)
			Purpose         string `json:"purpose" url:"purpose,omitempty"`                   //  (Optional) . Must be one of grading, bookmark
		} `json:"rubric_association" url:"rubric_association,omitempty"`
	} `json:"form"`
}

func (t *UpdateSingleRubric) GetMethod() string {
	return "PUT"
}

func (t *UpdateSingleRubric) GetURLPath() string {
	path := "courses/{course_id}/rubrics/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateSingleRubric) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSingleRubric) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateSingleRubric) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateSingleRubric) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.RubricAssociation.AssociationType != "" && !string_utils.Include([]string{"Assignment", "Course", "Account"}, t.Form.RubricAssociation.AssociationType) {
		errs = append(errs, "RubricAssociation must be one of Assignment, Course, Account")
	}
	if t.Form.RubricAssociation.Purpose != "" && !string_utils.Include([]string{"grading", "bookmark"}, t.Form.RubricAssociation.Purpose) {
		errs = append(errs, "RubricAssociation must be one of grading, bookmark")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateSingleRubric) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
