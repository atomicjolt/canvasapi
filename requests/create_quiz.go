package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateQuiz Create a new quiz for this course.
// https://canvas.instructure.com/doc/api/quizzes.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Quiz (Required) The quiz title.
// # Quiz (Optional) A description of the quiz.
// # Quiz (Optional) . Must be one of practice_quiz, assignment, graded_survey, surveyThe type of quiz.
// # Quiz (Optional) The assignment group id to put the assignment in. Defaults to the top
//    assignment group in the course. Only valid if the quiz is graded, i.e. if
//    quiz_type is "assignment" or "graded_survey".
// # Quiz (Optional) Time limit to take this quiz, in minutes. Set to null for no time limit.
//    Defaults to null.
// # Quiz (Optional) If true, quiz answers for multiple choice questions will be randomized for
//    each student. Defaults to false.
// # Quiz (Optional) . Must be one of always, until_after_last_attemptDictates whether or not quiz results are hidden from students.
//    If null, students can see their results after any attempt.
//    If "always", students can never see their results.
//    If "until_after_last_attempt", students can only see results after their
//    last attempt. (Only valid if allowed_attempts > 1). Defaults to null.
// # Quiz (Optional) Only valid if hide_results=null
//    If false, hides correct answers from students when quiz results are viewed.
//    Defaults to true.
// # Quiz (Optional) Only valid if show_correct_answers=true and allowed_attempts > 1
//    If true, hides correct answers from students when quiz results are viewed
//    until they submit the last attempt for the quiz.
//    Defaults to false.
// # Quiz (Optional) Only valid if show_correct_answers=true
//    If set, the correct answers will be visible by students only after this
//    date, otherwise the correct answers are visible once the student hands in
//    their quiz submission.
// # Quiz (Optional) Only valid if show_correct_answers=true
//    If set, the correct answers will stop being visible once this date has
//    passed. Otherwise, the correct answers will be visible indefinitely.
// # Quiz (Optional) Number of times a student is allowed to take a quiz.
//    Set to -1 for unlimited attempts.
//    Defaults to 1.
// # Quiz (Optional) . Must be one of keep_highest, keep_latestRequired and only valid if allowed_attempts > 1.
//    Scoring policy for a quiz that students can take multiple times.
//    Defaults to "keep_highest".
// # Quiz (Optional) If true, shows quiz to student one question at a time.
//    Defaults to false.
// # Quiz (Optional) Only valid if one_question_at_a_time=true
//    If true, questions are locked after answering.
//    Defaults to false.
// # Quiz (Optional) Restricts access to the quiz with a password.
//    For no access code restriction, set to null.
//    Defaults to null.
// # Quiz (Optional) Restricts access to the quiz to computers in a specified IP range.
//    Filters can be a comma-separated list of addresses, or an address followed by a mask
//
//    Examples:
//      "192.168.217.1"
//      "192.168.217.1/24"
//      "192.168.217.1/255.255.255.0"
//
//    For no IP filter restriction, set to null.
//    Defaults to null.
// # Quiz (Optional) The day/time the quiz is due.
//    Accepts times in ISO 8601 format, e.g. 2011-10-21T18:48Z.
// # Quiz (Optional) The day/time the quiz is locked for students.
//    Accepts times in ISO 8601 format, e.g. 2011-10-21T18:48Z.
// # Quiz (Optional) The day/time the quiz is unlocked for students.
//    Accepts times in ISO 8601 format, e.g. 2011-10-21T18:48Z.
// # Quiz (Optional) Whether the quiz should have a draft state of published or unpublished.
//    NOTE: If students have started taking the quiz, or there are any
//    submissions for the quiz, you may not unpublish a quiz and will recieve
//    an error.
// # Quiz (Optional) Whether students should be prevented from viewing their quiz results past
//    the first time (right after they turn the quiz in.)
//    Only valid if "hide_results" is not set to "always".
//    Defaults to false.
// # Quiz (Optional) Whether this quiz is only visible to overrides (Only useful if
//    'differentiated assignments' account setting is on)
//    Defaults to false.
//
type CreateQuiz struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Quiz struct {
			Title                         string    `json:"title"`                             //  (Required)
			Description                   string    `json:"description"`                       //  (Optional)
			QuizType                      string    `json:"quiz_type"`                         //  (Optional) . Must be one of practice_quiz, assignment, graded_survey, survey
			AssignmentGroupID             int64     `json:"assignment_group_id"`               //  (Optional)
			TimeLimit                     int64     `json:"time_limit"`                        //  (Optional)
			ShuffleAnswers                bool      `json:"shuffle_answers"`                   //  (Optional)
			HideResults                   string    `json:"hide_results"`                      //  (Optional) . Must be one of always, until_after_last_attempt
			ShowCorrectAnswers            bool      `json:"show_correct_answers"`              //  (Optional)
			ShowCorrectAnswersLastAttempt bool      `json:"show_correct_answers_last_attempt"` //  (Optional)
			ShowCorrectAnswersAt          time.Time `json:"show_correct_answers_at"`           //  (Optional)
			HideCorrectAnswersAt          time.Time `json:"hide_correct_answers_at"`           //  (Optional)
			AllowedAttempts               int64     `json:"allowed_attempts"`                  //  (Optional)
			ScoringPolicy                 string    `json:"scoring_policy"`                    //  (Optional) . Must be one of keep_highest, keep_latest
			OneQuestionAtATime            bool      `json:"one_question_at_a_time"`            //  (Optional)
			CantGoBack                    bool      `json:"cant_go_back"`                      //  (Optional)
			AccessCode                    string    `json:"access_code"`                       //  (Optional)
			IpFilter                      string    `json:"ip_filter"`                         //  (Optional)
			DueAt                         time.Time `json:"due_at"`                            //  (Optional)
			LockAt                        time.Time `json:"lock_at"`                           //  (Optional)
			UnlockAt                      time.Time `json:"unlock_at"`                         //  (Optional)
			Published                     bool      `json:"published"`                         //  (Optional)
			OneTimeResults                bool      `json:"one_time_results"`                  //  (Optional)
			OnlyVisibleToOverrides        bool      `json:"only_visible_to_overrides"`         //  (Optional)
		} `json:"quiz"`
	} `json:"form"`
}

func (t *CreateQuiz) GetMethod() string {
	return "POST"
}

func (t *CreateQuiz) GetURLPath() string {
	path := "courses/{course_id}/quizzes"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateQuiz) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateQuiz) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateQuiz) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Quiz.Title == "" {
		errs = append(errs, "'Quiz' is required")
	}
	if !string_utils.Include([]string{"practice_quiz", "assignment", "graded_survey", "survey"}, t.Form.Quiz.QuizType) {
		errs = append(errs, "Quiz must be one of practice_quiz, assignment, graded_survey, survey")
	}
	if !string_utils.Include([]string{"always", "until_after_last_attempt"}, t.Form.Quiz.HideResults) {
		errs = append(errs, "Quiz must be one of always, until_after_last_attempt")
	}
	if !string_utils.Include([]string{"keep_highest", "keep_latest"}, t.Form.Quiz.ScoringPolicy) {
		errs = append(errs, "Quiz must be one of keep_highest, keep_latest")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateQuiz) Do(c *canvasapi.Canvas) (*models.Quiz, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Quiz{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
