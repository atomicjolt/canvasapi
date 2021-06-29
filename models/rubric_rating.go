package models

type RubricRating struct {
	ID              string `json:"id" url:"id,omitempty"`                             // Example: name_2
	CriterionID     string `json:"criterion_id" url:"criterion_id,omitempty"`         // Example: _10
	Description     string `json:"description" url:"description,omitempty"`           //
	LongDescription string `json:"long_description" url:"long_description,omitempty"` //
	Points          int64  `json:"points" url:"points,omitempty"`                     // Example: 5
}

func (t *RubricRating) HasErrors() error {
	return nil
}
