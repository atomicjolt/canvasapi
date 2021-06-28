package models

type ScoreStatistic struct {
	Min  int64 `json:"min" url:"min,omitempty"`   // Min score.Example: 1
	Max  int64 `json:"max" url:"max,omitempty"`   // Max score.Example: 10
	Mean int64 `json:"mean" url:"mean,omitempty"` // Mean score.Example: 6
}

func (t *ScoreStatistic) HasError() error {
	return nil
}
