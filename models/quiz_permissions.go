package models

type QuizPermissions struct {
	Read           bool `json:"read" url:"read,omitempty"`                       // whether the user can view the quiz.Example: true
	Submit         bool `json:"submit" url:"submit,omitempty"`                   // whether the user may submit a submission for the quiz.Example: true
	Create         bool `json:"create" url:"create,omitempty"`                   // whether the user may create a new quiz.Example: true
	Manage         bool `json:"manage" url:"manage,omitempty"`                   // whether the user may edit, update, or delete the quiz.Example: true
	ReadStatistics bool `json:"read_statistics" url:"read_statistics,omitempty"` // whether the user may view quiz statistics for this quiz.Example: true
	ReviewGrades   bool `json:"review_grades" url:"review_grades,omitempty"`     // whether the user may review grades for all quiz submissions for this quiz.Example: true
	Update         bool `json:"update" url:"update,omitempty"`                   // whether the user may update the quiz.Example: true
}

func (t *QuizPermissions) HasError() error {
	return nil
}
