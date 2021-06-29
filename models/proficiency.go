package models

type Proficiency struct {
	Ratings []string `json:"ratings" url:"ratings,omitempty"` // An array of proficiency ratings. See the ProficiencyRating specification above..
}

func (t *Proficiency) HasErrors() error {
	return nil
}
