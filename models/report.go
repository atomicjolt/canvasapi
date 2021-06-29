package models

import (
	"time"
)

type Report struct {
	ID          int64             `json:"id" url:"id,omitempty"`                     // The unique identifier for the report..Example: 1
	Report      string            `json:"report" url:"report,omitempty"`             // The type of report..Example: sis_export_csv
	FileUrl     string            `json:"file_url" url:"file_url,omitempty"`         // The url to the report download..Example: https://example.com/some/path
	Attachment  *File             `json:"attachment" url:"attachment,omitempty"`     // The attachment api object of the report. Only available after the report has completed..
	Status      string            `json:"status" url:"status,omitempty"`             // The status of the report.Example: complete
	CreatedAt   time.Time         `json:"created_at" url:"created_at,omitempty"`     // The date and time the report was created..Example: 2013-12-01T23:59:00-06:00
	StartedAt   time.Time         `json:"started_at" url:"started_at,omitempty"`     // The date and time the report started processing..Example: 2013-12-02T00:03:21-06:00
	EndedAt     time.Time         `json:"ended_at" url:"ended_at,omitempty"`         // The date and time the report finished processing..Example: 2013-12-02T00:03:21-06:00
	Parameters  *ReportParameters `json:"parameters" url:"parameters,omitempty"`     // The report parameters.Example: 2, 2012-07-13T10:55:20-06:00, 2012-07-13T10:55:20-06:00
	Progress    int64             `json:"progress" url:"progress,omitempty"`         // The progress of the report.Example: 100
	CurrentLine int64             `json:"current_line" url:"current_line,omitempty"` // This is the current line count being written to the report. It updates every 1000 records..Example: 12000
}

func (t *Report) HasErrors() error {
	return nil
}
