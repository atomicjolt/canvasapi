package models

type RubricCriterion struct {
	ID                string          `json:"id"`                  // the ID of the criterion.Example: _10
	Description       string          `json:"description"`         //
	LongDescription   string          `json:"long_description"`    //
	Points            int64           `json:"points"`              // Example: 5
	CriterionUseRange bool            `json:"criterion_use_range"` // Example: false
	Ratings           []*RubricRating `json:"ratings"`             // the possible ratings for this Criterion.
}

func (t *RubricCriterion) HasError() error {
	return nil
}
