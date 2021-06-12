package models

type QuizPermissions struct {
	Read           bool `json:"read"`            // whether the user can view the quiz.Example: true
	Submit         bool `json:"submit"`          // whether the user may submit a submission for the quiz.Example: true
	Create         bool `json:"create"`          // whether the user may create a new quiz.Example: true
	Manage         bool `json:"manage"`          // whether the user may edit, update, or delete the quiz.Example: true
	ReadStatistics bool `json:"read_statistics"` // whether the user may view quiz statistics for this quiz.Example: true
	ReviewGrades   bool `json:"review_grades"`   // whether the user may review grades for all quiz submissions for this quiz.Example: true
	Update         bool `json:"update"`          // whether the user may update the quiz.Example: true
}

func (t *QuizPermissions) HasError() error {
	return nil
}
