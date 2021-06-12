package models

type GradingRules struct {
	DropLowest  int64   `json:"drop_lowest"`  // Number of lowest scores to be dropped for each user..Example: 1
	DropHighest int64   `json:"drop_highest"` // Number of highest scores to be dropped for each user..Example: 1
	NeverDrop   []int64 `json:"never_drop"`   // Assignment IDs that should never be dropped..Example: 33, 17, 24
}

func (t *GradingRules) HasError() error {
	return nil
}
