package models

type NeedsGradingCount struct {
	SectionID         string `json:"section_id"`          // The section ID.Example: 123456
	NeedsGradingCount int64  `json:"needs_grading_count"` // Number of submissions that need grading.Example: 5
}

func (t *NeedsGradingCount) HasError() error {
	return nil
}
