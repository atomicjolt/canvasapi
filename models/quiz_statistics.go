package models

import (
	"time"
)

type QuizStatistics struct {
	ID     int64 `json:"id" url:"id,omitempty"`           // The ID of the quiz statistics report..Example: 1
	QuizID int64 `json:"quiz_id" url:"quiz_id,omitempty"` // The ID of the Quiz the statistics report is for.
	//NOTE: AVAILABLE ONLY IN NON-JSON-API REQUESTS..Example: 2
	MultipleAttemptsExist bool                                `json:"multiple_attempts_exist" url:"multiple_attempts_exist,omitempty"` // Whether there are any students that have made mutliple submissions for this quiz..Example: true
	IncludesAllVersions   bool                                `json:"includes_all_versions" url:"includes_all_versions,omitempty"`     // In the presence of multiple attempts, this field describes whether the statistics describe all the submission attempts and not only the latest ones..Example: true
	GeneratedAt           time.Time                           `json:"generated_at" url:"generated_at,omitempty"`                       // The time at which the statistics were generated, which is usually after the occurrence of a quiz event, like a student submitting it..Example: 2013-01-23T23:59:00-07:00
	Url                   string                              `json:"url" url:"url,omitempty"`                                         // The API HTTP/HTTPS URL to this quiz statistics..Example: http://canvas.example.edu/api/v1/courses/1/quizzes/2/statistics
	HtmlUrl               string                              `json:"html_url" url:"html_url,omitempty"`                               // The HTTP/HTTPS URL to the page where the statistics can be seen visually..Example: http://canvas.example.edu/courses/1/quizzes/2/statistics
	QuestionStatistics    *QuizStatisticsQuestionStatistics   `json:"question_statistics" url:"question_statistics,omitempty"`         // Question-specific statistics for each question and its answers..
	SubmissionStatistics  *QuizStatisticsSubmissionStatistics `json:"submission_statistics" url:"submission_statistics,omitempty"`     // Question-specific statistics for each question and its answers..
	Links                 *QuizStatisticsLinks                `json:"links" url:"links,omitempty"`                                     // JSON-API construct that contains links to media related to this quiz statistics object.
	//NOTE: AVAILABLE ONLY IN JSON-API REQUESTS..
}

func (t *QuizStatistics) HasErrors() error {
	return nil
}
