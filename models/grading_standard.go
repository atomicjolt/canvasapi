package models

type GradingStandard struct {
	Title         string                `json:"title" url:"title,omitempty"`                   // the title of the grading standard.Example: Account Standard
	ID            int64                 `json:"id" url:"id,omitempty"`                         // the id of the grading standard.Example: 1
	ContextType   string                `json:"context_type" url:"context_type,omitempty"`     // the context this standard is associated with, either 'Account' or 'Course'.Example: Account
	ContextID     int64                 `json:"context_id" url:"context_id,omitempty"`         // the id for the context either the Account or Course id.Example: 1
	GradingScheme []*GradingSchemeEntry `json:"grading_scheme" url:"grading_scheme,omitempty"` // A list of GradingSchemeEntry that make up the Grading Standard as an array of values with the scheme name and value.Example: {'name'=>'A', 'value'=>0.9}, {'name'=>'B', 'value'=>0.8}, {'name'=>'C', 'value'=>0.7}, {'name'=>'D', 'value'=>0.6}
}

func (t *GradingStandard) HasError() error {
	return nil
}
