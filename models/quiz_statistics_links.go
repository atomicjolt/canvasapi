package models

type QuizStatisticsLinks struct {
	Quiz string `json:"quiz" url:"quiz,omitempty"` // HTTP/HTTPS API URL to the quiz this statistics describe..Example: http://canvas.example.edu/api/v1/courses/1/quizzes/2
}

func (t *QuizStatisticsLinks) HasError() error {
	return nil
}
