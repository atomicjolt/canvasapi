package models

type ScoreStatistic struct {
	Min  int64 `json:"min"`  // Min score.Example: 1
	Max  int64 `json:"max"`  // Max score.Example: 10
	Mean int64 `json:"mean"` // Mean score.Example: 6
}

func (t *ScoreStatistic) HasError() error {
	return nil
}
