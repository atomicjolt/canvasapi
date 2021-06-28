package models

type OutcomeRollupScoreLinks struct {
	Outcome int64 `json:"outcome" url:"outcome,omitempty"` // The id of the related outcome.Example: 42
}

func (t *OutcomeRollupScoreLinks) HasError() error {
	return nil
}
