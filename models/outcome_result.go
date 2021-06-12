package models

import (
	"time"
)

type OutcomeResult struct {
	ID                    int64     `json:"id"`                       // A unique identifier for this result.Example: 42
	Score                 int64     `json:"score"`                    // The student's score.Example: 6
	SubmittedOrAssessedAt time.Time `json:"submitted_or_assessed_at"` // The datetime the resulting OutcomeResult was submitted at, or absent that, when it was assessed..Example: 2013-02-01T00:00:00-06:00
	Links                 string    `json:"links"`                    // Unique identifiers of objects associated with this result.Example: 3, 97, 53
	Percent               float64   `json:"percent"`                  // score's percent of maximum points possible for outcome, scaled to reflect any custom mastery levels that differ from the learning outcome.Example: 0.65
}

func (t *OutcomeResult) HasError() error {
	return nil
}
