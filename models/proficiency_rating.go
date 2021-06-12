package models

type ProficiencyRating struct {
	Description string  `json:"description"` // The description of the rating.Example: Exceeds Mastery
	Points      float64 `json:"points"`      // A non-negative number of points for the rating.Example: 4
	Mastery     bool    `json:"mastery"`     // Indicates the rating where mastery is first achieved.
	Color       string  `json:"color"`       // The hex color code of the rating.Example: 127A1B
}

func (t *ProficiencyRating) HasError() error {
	return nil
}
