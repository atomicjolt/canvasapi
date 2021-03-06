package models

type SISImportError struct {
	SISImportID int64  `json:"sis_import_id" url:"sis_import_id,omitempty"` // The unique identifier for the SIS import..Example: 1
	File        string `json:"file" url:"file,omitempty"`                   // The file where the error message occurred..Example: courses.csv
	Message     string `json:"message" url:"message,omitempty"`             // The error message that from the record..Example: No short_name given for course C001
	RowInfo     string `json:"row_info" url:"row_info,omitempty"`           // The contents of the line that had the error..Example: account_1, Sub account 1,, active
	Row         int64  `json:"row" url:"row,omitempty"`                     // The line number where the error occurred. Some Importers do not yet support this. This is a 1 based index starting with the header row..Example: 34
}

func (t *SISImportError) HasErrors() error {
	return nil
}
