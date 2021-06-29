package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// UpdateExistingQuizQuestion Updates an existing quiz question for this quiz
// https://canvas.instructure.com/doc/api/quiz_questions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) The associated quiz's unique identifier.
// # Path.ID (Required) The quiz question's unique identifier.
//
// Form Parameters:
// # Form.Question.QuestionName (Optional) The name of the question.
// # Form.Question.QuestionText (Optional) The text of the question.
// # Form.Question.QuizGroupID (Optional) The id of the quiz group to assign the question to.
// # Form.Question.QuestionType (Optional) . Must be one of calculated_question, essay_question, file_upload_question, fill_in_multiple_blanks_question, matching_question, multiple_answers_question, multiple_choice_question, multiple_dropdowns_question, numerical_question, short_answer_question, text_only_question, true_false_questionThe type of question. Multiple optional fields depend upon the type of question to be used.
// # Form.Question.Position (Optional) The order in which the question will be displayed in the quiz in relation to other questions.
// # Form.Question.PointsPossible (Optional) The maximum amount of points received for answering this question correctly.
// # Form.Question.CorrectComments (Optional) The comment to display if the student answers the question correctly.
// # Form.Question.IncorrectComments (Optional) The comment to display if the student answers incorrectly.
// # Form.Question.NeutralComments (Optional) The comment to display regardless of how the student answered.
// # Form.Question.TextAfterAnswers (Optional) no description
// # Form.Question.Answers (Optional) no description
//
type UpdateExistingQuizQuestion struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   int64  `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       int64  `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		Question struct {
			QuestionName      string           `json:"question_name" url:"question_name,omitempty"`           //  (Optional)
			QuestionText      string           `json:"question_text" url:"question_text,omitempty"`           //  (Optional)
			QuizGroupID       int64            `json:"quiz_group_id" url:"quiz_group_id,omitempty"`           //  (Optional)
			QuestionType      string           `json:"question_type" url:"question_type,omitempty"`           //  (Optional) . Must be one of calculated_question, essay_question, file_upload_question, fill_in_multiple_blanks_question, matching_question, multiple_answers_question, multiple_choice_question, multiple_dropdowns_question, numerical_question, short_answer_question, text_only_question, true_false_question
			Position          int64            `json:"position" url:"position,omitempty"`                     //  (Optional)
			PointsPossible    int64            `json:"points_possible" url:"points_possible,omitempty"`       //  (Optional)
			CorrectComments   string           `json:"correct_comments" url:"correct_comments,omitempty"`     //  (Optional)
			IncorrectComments string           `json:"incorrect_comments" url:"incorrect_comments,omitempty"` //  (Optional)
			NeutralComments   string           `json:"neutral_comments" url:"neutral_comments,omitempty"`     //  (Optional)
			TextAfterAnswers  string           `json:"text_after_answers" url:"text_after_answers,omitempty"` //  (Optional)
			Answers           []*models.Answer `json:"answers" url:"answers,omitempty"`                       //  (Optional)
		} `json:"question" url:"question,omitempty"`
	} `json:"form"`
}

func (t *UpdateExistingQuizQuestion) GetMethod() string {
	return "PUT"
}

func (t *UpdateExistingQuizQuestion) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/questions/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateExistingQuizQuestion) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateExistingQuizQuestion) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateExistingQuizQuestion) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateExistingQuizQuestion) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.Question.QuestionType != "" && !string_utils.Include([]string{"calculated_question", "essay_question", "file_upload_question", "fill_in_multiple_blanks_question", "matching_question", "multiple_answers_question", "multiple_choice_question", "multiple_dropdowns_question", "numerical_question", "short_answer_question", "text_only_question", "true_false_question"}, t.Form.Question.QuestionType) {
		errs = append(errs, "Question must be one of calculated_question, essay_question, file_upload_question, fill_in_multiple_blanks_question, matching_question, multiple_answers_question, multiple_choice_question, multiple_dropdowns_question, numerical_question, short_answer_question, text_only_question, true_false_question")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateExistingQuizQuestion) Do(c *canvasapi.Canvas) (*models.QuizQuestion, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.QuizQuestion{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
