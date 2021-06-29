package models

type RubricCriterion struct {
	ID                string          `json:"id" url:"id,omitempty"`                                   // the ID of the criterion.Example: _10
	Description       string          `json:"description" url:"description,omitempty"`                 //
	LongDescription   string          `json:"long_description" url:"long_description,omitempty"`       //
	Points            int64           `json:"points" url:"points,omitempty"`                           // Example: 5
	CriterionUseRange bool            `json:"criterion_use_range" url:"criterion_use_range,omitempty"` // Example: false
	Ratings           []*RubricRating `json:"ratings" url:"ratings,omitempty"`                         // the possible ratings for this Criterion.
}

func (t *RubricCriterion) HasErrors() error {
	return nil
}
