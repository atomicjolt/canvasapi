package models

type Assessment struct {
	ID    string `json:"id"`    // A unique identifier for this live assessment.Example: 42
	Key   string `json:"key"`   // A client specified unique identifier for the assessment.Example: 2014-05-27,outcome_52
	Title string `json:"title"` // A human readable title for the assessment.Example: May 27th Reading Assessment
}

func (t *Assessment) HasError() error {
	return nil
}
