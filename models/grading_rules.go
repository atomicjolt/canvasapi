package models

type GradingRules struct {
	DropLowest  int64   `json:"drop_lowest" url:"drop_lowest,omitempty"`   // Number of lowest scores to be dropped for each user..Example: 1
	DropHighest int64   `json:"drop_highest" url:"drop_highest,omitempty"` // Number of highest scores to be dropped for each user..Example: 1
	NeverDrop   []int64 `json:"never_drop" url:"never_drop,omitempty"`     // Assignment IDs that should never be dropped..Example: 33, 17, 24
}

func (t *GradingRules) HasError() error {
	return nil
}
