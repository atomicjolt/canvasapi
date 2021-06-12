package models

type TurnitinSettings struct {
	OriginalityReportVisibility string `json:"originality_report_visibility"` // Example: after_grading
	SPaperCheck                 bool   `json:"s_paper_check"`                 //
	InternetCheck               bool   `json:"internet_check"`                //
	JournalCheck                bool   `json:"journal_check"`                 //
	ExcludeBiblio               bool   `json:"exclude_biblio"`                //
	ExcludeQuoted               bool   `json:"exclude_quoted"`                //
	ExcludeSmallMatchesType     string `json:"exclude_small_matches_type"`    // Example: percent
	ExcludeSmallMatchesValue    int64  `json:"exclude_small_matches_value"`   // Example: 50
}

func (t *TurnitinSettings) HasError() error {
	return nil
}
