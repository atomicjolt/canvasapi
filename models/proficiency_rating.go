package models

type ProficiencyRating struct {
	Description string  `json:"description" url:"description,omitempty"` // The description of the rating.Example: Exceeds Mastery
	Points      float64 `json:"points" url:"points,omitempty"`           // A non-negative number of points for the rating.Example: 4
	Mastery     bool    `json:"mastery" url:"mastery,omitempty"`         // Indicates the rating where mastery is first achieved.
	Color       string  `json:"color" url:"color,omitempty"`             // The hex color code of the rating.Example: 127A1B
}

func (t *ProficiencyRating) HasError() error {
	return nil
}
