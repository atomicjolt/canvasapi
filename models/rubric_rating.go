package models

type RubricRating struct {
	ID              string `json:"id"`               // Example: name_2
	CriterionID     string `json:"criterion_id"`     // Example: _10
	Description     string `json:"description"`      //
	LongDescription string `json:"long_description"` //
	Points          int64  `json:"points"`           // Example: 5
}

func (t *RubricRating) HasError() error {
	return nil
}
