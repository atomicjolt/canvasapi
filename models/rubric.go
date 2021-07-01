package models

type Rubric struct {
	ID                        int64                `json:"id" url:"id,omitempty"`                                                     // the ID of the rubric.Example: 1
	Title                     string               `json:"title" url:"title,omitempty"`                                               // title of the rubric.Example: some title
	ContextID                 int64                `json:"context_id" url:"context_id,omitempty"`                                     // the context owning the rubric.Example: 1
	ContextType               string               `json:"context_type" url:"context_type,omitempty"`                                 // Example: Course
	PointsPossible            float64              `json:"points_possible" url:"points_possible,omitempty"`                           // Example: 10.0
	Reusable                  bool                 `json:"reusable" url:"reusable,omitempty"`                                         // Example: false
	ReadOnly                  bool                 `json:"read_only" url:"read_only,omitempty"`                                       // Example: true
	FreeFormCriterionComments bool                 `json:"free_form_criterion_comments" url:"free_form_criterion_comments,omitempty"` // whether or not free-form comments are used.Example: true
	HideScoreTotal            bool                 `json:"hide_score_total" url:"hide_score_total,omitempty"`                         // Example: true
	Data                      []*RubricCriterion   `json:"data" url:"data,omitempty"`                                                 // An array with all of this Rubric's grading Criteria.
	Assessments               []*RubricAssessment  `json:"assessments" url:"assessments,omitempty"`                                   // If an assessment type is included in the 'include' parameter, includes an array of rubric assessment objects for a given rubric, based on the assessment type requested. If the user does not request an assessment type this key will be absent..
	Associations              []*RubricAssociation `json:"associations" url:"associations,omitempty"`                                 // If an association type is included in the 'include' parameter, includes an array of rubric association objects for a given rubric, based on the association type requested. If the user does not request an association type this key will be absent..
}

func (t *Rubric) HasErrors() error {
	return nil
}
