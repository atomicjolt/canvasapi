package models

type OutcomeRollupLinks struct {
	Course  int64 `json:"course"`  // If an aggregate result was requested, the course field will be present. Otherwise, the user and section field will be present (Optional) The id of the course that this rollup applies to.Example: 42
	User    int64 `json:"user"`    // (Optional) The id of the user that this rollup applies to.Example: 42
	Section int64 `json:"section"` // (Optional) The id of the section the user is in.Example: 57
}

func (t *OutcomeRollupLinks) HasError() error {
	return nil
}
