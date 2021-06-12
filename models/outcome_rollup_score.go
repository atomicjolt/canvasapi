package models

type OutcomeRollupScore struct {
	Score int64                    `json:"score"` // The rollup score for the outcome, based on the student alignment scores related to the outcome. This could be null if the student has no related scores..Example: 3
	Count int64                    `json:"count"` // The number of alignment scores included in this rollup..Example: 6
	Links *OutcomeRollupScoreLinks `json:"links"` // Example: 42
}

func (t *OutcomeRollupScore) HasError() error {
	return nil
}
