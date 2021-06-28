package models

type RubricAssociation struct {
	ID                 int64  `json:"id" url:"id,omitempty"`                                     // the ID of the association.Example: 1
	RubricID           int64  `json:"rubric_id" url:"rubric_id,omitempty"`                       // the ID of the rubric.Example: 1
	AssociationID      int64  `json:"association_id" url:"association_id,omitempty"`             // the ID of the object this association links to.Example: 1
	AssociationType    string `json:"association_type" url:"association_type,omitempty"`         // the type of object this association links to.Example: Course
	UseForGrading      bool   `json:"use_for_grading" url:"use_for_grading,omitempty"`           // Whether or not the associated rubric is used for grade calculation.Example: true
	SummaryData        string `json:"summary_data" url:"summary_data,omitempty"`                 //
	Purpose            string `json:"purpose" url:"purpose,omitempty"`                           // Whether or not the association is for grading (and thus linked to an assignment) or if it's to indicate the rubric should appear in its context. Values will be grading or bookmark..Example: grading
	HideScoreTotal     bool   `json:"hide_score_total" url:"hide_score_total,omitempty"`         // Whether or not the score total is displayed within the rubric. This option is only available if the rubric is not used for grading..Example: true
	HidePoints         bool   `json:"hide_points" url:"hide_points,omitempty"`                   // Example: true
	HideOutcomeResults bool   `json:"hide_outcome_results" url:"hide_outcome_results,omitempty"` // Example: true
}

func (t *RubricAssociation) HasError() error {
	return nil
}
