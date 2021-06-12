package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type QuizReport struct {
	ID                  int64     `json:"id"`                    // the ID of the quiz report.Example: 5
	QuizID              int64     `json:"quiz_id"`               // the ID of the quiz.Example: 4
	ReportType          string    `json:"report_type"`           // which type of report this is possible values: 'student_analysis', 'item_analysis'.Example: student_analysis
	ReadableType        string    `json:"readable_type"`         // a human-readable (and localized) version of the report_type.Example: Student Analysis
	IncludesAllVersions bool      `json:"includes_all_versions"` // boolean indicating whether the report represents all submissions or only the most recent ones for each student.Example: true
	Anonymous           bool      `json:"anonymous"`             // boolean indicating whether the report is for an anonymous survey. if true, no student names will be included in the csv.
	Generatable         bool      `json:"generatable"`           // boolean indicating whether the report can be generated, which is true unless the quiz is a survey one.Example: true
	CreatedAt           time.Time `json:"created_at"`            // when the report was created.Example: 2013-05-01T12:34:56-07:00
	UpdatedAt           time.Time `json:"updated_at"`            // when the report was last updated.Example: 2013-05-01T12:34:56-07:00
	Url                 string    `json:"url"`                   // the API endpoint for this report.Example: http://canvas.example.com/api/v1/courses/1/quizzes/1/reports/1
	File                *File     `json:"file"`                  // if the report has finished generating, a File object that represents it. refer to the Files API for more information about the format.
	ProgressUrl         string    `json:"progress_url"`          // if the report has not yet finished generating, a URL where information about its progress can be retrieved. refer to the Progress API for more information (Note: not available in JSON-API format).
	Progress            *Progress `json:"progress"`              // if the report is being generated, a Progress object that represents the operation. Refer to the Progress API for more information about the format. (Note: available only in JSON-API format).
}

func (t *QuizReport) HasError() error {
	var s []string
	s = []string{"student_analysis", "item_analysis"}
	if !string_utils.Include(s, t.ReportType) {
		return fmt.Errorf("expected 'report_type' to be one of %v", s)
	}
	return nil
}