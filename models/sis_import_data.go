package models

type SISImportData struct {
	ImportType      string           `json:"import_type" url:"import_type,omitempty"`           // The type of SIS import.Example: instructure_csv
	SuppliedBatches []string         `json:"supplied_batches" url:"supplied_batches,omitempty"` // Which files were included in the SIS import.Example: term, course, section, user, enrollment
	Counts          *SISImportCounts `json:"counts" url:"counts,omitempty"`                     // The number of rows processed for each type of import.
}

func (t *SISImportData) HasErrors() error {
	return nil
}
