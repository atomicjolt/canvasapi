package models

type RubricCriteria struct {
	Points            int64           `json:"points" url:"points,omitempty"`                           // Example: 10
	ID                string          `json:"id" url:"id,omitempty"`                                   // The id of rubric criteria..Example: crit1
	LearningOutcomeID string          `json:"learning_outcome_id" url:"learning_outcome_id,omitempty"` // (Optional) The id of the learning outcome this criteria uses, if any..Example: 1234
	VendorGuid        string          `json:"vendor_guid" url:"vendor_guid,omitempty"`                 // (Optional) The 3rd party vendor's GUID for the outcome this criteria references, if any..Example: abdsfjasdfne3jsdfn2
	Description       string          `json:"description" url:"description,omitempty"`                 // Example: Criterion 1
	LongDescription   string          `json:"long_description" url:"long_description,omitempty"`       // Example: Criterion 1 more details
	CriterionUseRange bool            `json:"criterion_use_range" url:"criterion_use_range,omitempty"` // Example: true
	Ratings           []*RubricRating `json:"ratings" url:"ratings,omitempty"`                         //
	IgnoreForScoring  bool            `json:"ignore_for_scoring" url:"ignore_for_scoring,omitempty"`   // Example: true
}

func (t *RubricCriteria) HasErrors() error {
	return nil
}
