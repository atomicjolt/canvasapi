package models

type TurnitinSettings struct {
	OriginalityReportVisibility string `json:"originality_report_visibility" url:"originality_report_visibility,omitempty"` // Example: after_grading
	SPaperCheck                 bool   `json:"s_paper_check" url:"s_paper_check,omitempty"`                                 //
	InternetCheck               bool   `json:"internet_check" url:"internet_check,omitempty"`                               //
	JournalCheck                bool   `json:"journal_check" url:"journal_check,omitempty"`                                 //
	ExcludeBiblio               bool   `json:"exclude_biblio" url:"exclude_biblio,omitempty"`                               //
	ExcludeQuoted               bool   `json:"exclude_quoted" url:"exclude_quoted,omitempty"`                               //
	ExcludeSmallMatchesType     string `json:"exclude_small_matches_type" url:"exclude_small_matches_type,omitempty"`       // Example: percent
	ExcludeSmallMatchesValue    int64  `json:"exclude_small_matches_value" url:"exclude_small_matches_value,omitempty"`     // Example: 50
}

func (t *TurnitinSettings) HasErrors() error {
	return nil
}
