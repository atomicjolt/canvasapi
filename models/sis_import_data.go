package models

type SISImportData struct {
	ImportType      string           `json:"import_type"`      // The type of SIS import.Example: instructure_csv
	SuppliedBatches []string         `json:"supplied_batches"` // Which files were included in the SIS import.Example: term, course, section, user, enrollment
	Counts          *SISImportCounts `json:"counts"`           // The number of rows processed for each type of import.
}

func (t *SISImportData) HasError() error {
	return nil
}
