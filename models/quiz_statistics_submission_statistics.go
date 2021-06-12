package models

type QuizStatisticsSubmissionStatistics struct {
	UniqueCount           int64   `json:"unique_count"`            // The number of students who have taken the quiz..Example: 3
	ScoreAverage          float64 `json:"score_average"`           // The mean of the student submission scores..Example: 4.33333333333333
	ScoreHigh             float64 `json:"score_high"`              // The highest submission score..Example: 6
	ScoreLow              float64 `json:"score_low"`               // The lowest submission score..Example: 3
	ScoreStdev            float64 `json:"score_stdev"`             // Standard deviation of the submission scores..Example: 1.24721912892465
	Scores                string  `json:"scores"`                  // A percentile distribution of the student scores, each key is the percentile (ranges between 0 and 100%) while the value is the number of students who received that score..Example: 1, 5, 1
	CorrectCountAverage   float64 `json:"correct_count_average"`   // The mean of the number of questions answered correctly by each student..Example: 3.66666666666667
	IncorrectCountAverage float64 `json:"incorrect_count_average"` // The mean of the number of questions answered incorrectly by each student..Example: 5
	DurationAverage       float64 `json:"duration_average"`        // The average time spent by students while taking the quiz..Example: 42.333333333
}

func (t *QuizStatisticsSubmissionStatistics) HasError() error {
	return nil
}
