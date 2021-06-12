package models

type Rubric struct {
	ID                        int64                `json:"id"`                           // the ID of the rubric.Example: 1
	Title                     string               `json:"title"`                        // title of the rubric.Example: some title
	ContextID                 int64                `json:"context_id"`                   // the context owning the rubric.Example: 1
	ContextType               string               `json:"context_type"`                 // Example: Course
	PointsPossible            int64                `json:"points_possible"`              // Example: 10.0
	Reusable                  bool                 `json:"reusable"`                     // Example: false
	ReadOnly                  bool                 `json:"read_only"`                    // Example: true
	FreeFormCriterionComments bool                 `json:"free_form_criterion_comments"` // whether or not free-form comments are used.Example: true
	HideScoreTotal            bool                 `json:"hide_score_total"`             // Example: true
	Data                      []*RubricCriterion   `json:"data"`                         // An array with all of this Rubric's grading Criteria.
	Assessments               []*RubricAssessment  `json:"assessments"`                  // If an assessment type is included in the 'include' parameter, includes an array of rubric assessment objects for a given rubric, based on the assessment type requested. If the user does not request an assessment type this key will be absent..
	Associations              []*RubricAssociation `json:"associations"`                 // If an association type is included in the 'include' parameter, includes an array of rubric association objects for a given rubric, based on the association type requested. If the user does not request an association type this key will be absent..
}

func (t *Rubric) HasError() error {
	return nil
}
