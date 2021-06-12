package models

type OutcomeImportData struct {
	ImportType string `json:"import_type"` // The type of outcome import.Example: instructure_csv
}

func (t *OutcomeImportData) HasError() error {
	return nil
}
