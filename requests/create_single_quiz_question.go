package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateSingleQuizQuestion Create a new quiz question for this quiz
// https://canvas.instructure.com/doc/api/quiz_questions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Form Parameters:
// # Question (Optional) The name of the question.
// # Question (Optional) The text of the question.
// # Question (Optional) The id of the quiz group to assign the question to.
// # Question (Optional) . Must be one of calculated_question, essay_question, file_upload_question, fill_in_multiple_blanks_question, matching_question, multiple_answers_question, multiple_choice_question, multiple_dropdowns_question, numerical_question, short_answer_question, text_only_question, true_false_questionThe type of question. Multiple optional fields depend upon the type of question to be used.
// # Question (Optional) The order in which the question will be displayed in the quiz in relation to other questions.
// # Question (Optional) The maximum amount of points received for answering this question correctly.
// # Question (Optional) The comment to display if the student answers the question correctly.
// # Question (Optional) The comment to display if the student answers incorrectly.
// # Question (Optional) The comment to display regardless of how the student answered.
// # Question (Optional) no description
// # Question (Optional) no description
//
type CreateSingleQuizQuestion struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
	} `json:"path"`

	Form struct {
		Question struct {
			QuestionName      string `json:"question_name"`      //  (Optional)
			QuestionText      string `json:"question_text"`      //  (Optional)
			QuizGroupID       int64  `json:"quiz_group_id"`      //  (Optional)
			QuestionType      string `json:"question_type"`      //  (Optional) . Must be one of calculated_question, essay_question, file_upload_question, fill_in_multiple_blanks_question, matching_question, multiple_answers_question, multiple_choice_question, multiple_dropdowns_question, numerical_question, short_answer_question, text_only_question, true_false_question
			Position          int64  `json:"position"`           //  (Optional)
			PointsPossible    int64  `json:"points_possible"`    //  (Optional)
			CorrectComments   string `json:"correct_comments"`   //  (Optional)
			IncorrectComments string `json:"incorrect_comments"` //  (Optional)
			NeutralComments   string `json:"neutral_comments"`   //  (Optional)
			TextAfterAnswers  string `json:"text_after_answers"` //  (Optional)
			Answers           string `json:"answers"`            //  (Optional)
		} `json:"question"`
	} `json:"form"`
}

func (t *CreateSingleQuizQuestion) GetMethod() string {
	return "POST"
}

func (t *CreateSingleQuizQuestion) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/questions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *CreateSingleQuizQuestion) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSingleQuizQuestion) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateSingleQuizQuestion) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if !string_utils.Include([]string{"calculated_question", "essay_question", "file_upload_question", "fill_in_multiple_blanks_question", "matching_question", "multiple_answers_question", "multiple_choice_question", "multiple_dropdowns_question", "numerical_question", "short_answer_question", "text_only_question", "true_false_question"}, t.Form.Question.QuestionType) {
		errs = append(errs, "Question must be one of calculated_question, essay_question, file_upload_question, fill_in_multiple_blanks_question, matching_question, multiple_answers_question, multiple_choice_question, multiple_dropdowns_question, numerical_question, short_answer_question, text_only_question, true_false_question")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSingleQuizQuestion) Do(c *canvasapi.Canvas) (*models.QuizQuestion, error) {
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
