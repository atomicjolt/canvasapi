package models

type RubricCriteria struct {
	Points            int64           `json:"points"`              // Example: 10
	ID                string          `json:"id"`                  // The id of rubric criteria..Example: crit1
	LearningOutcomeID string          `json:"learning_outcome_id"` // (Optional) The id of the learning outcome this criteria uses, if any..Example: 1234
	VendorGuid        string          `json:"vendor_guid"`         // (Optional) The 3rd party vendor's GUID for the outcome this criteria references, if any..Example: abdsfjasdfne3jsdfn2
	Description       string          `json:"description"`         // Example: Criterion 1
	LongDescription   string          `json:"long_description"`    // Example: Criterion 1 more details
	CriterionUseRange bool            `json:"criterion_use_range"` // Example: true
	Ratings           []*RubricRating `json:"ratings"`             //
	IgnoreForScoring  bool            `json:"ignore_for_scoring"`  // Example: true
}

func (t *RubricCriteria) HasError() error {
	return nil
}
