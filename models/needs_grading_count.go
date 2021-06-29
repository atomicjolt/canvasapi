package models

type NeedsGradingCount struct {
	SectionID         string `json:"section_id" url:"section_id,omitempty"`                   // The section ID.Example: 123456
	NeedsGradingCount int64  `json:"needs_grading_count" url:"needs_grading_count,omitempty"` // Number of submissions that need grading.Example: 5
}

func (t *NeedsGradingCount) HasErrors() error {
	return nil
}
