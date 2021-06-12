package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Quiz struct {
	ID                            int64             `json:"id"`                                // the ID of the quiz.Example: 5
	Title                         string            `json:"title"`                             // the title of the quiz.Example: Hamlet Act 3 Quiz
	HtmlUrl                       string            `json:"html_url"`                          // the HTTP/HTTPS URL to the quiz.Example: http://canvas.example.edu/courses/1/quizzes/2
	MobileUrl                     string            `json:"mobile_url"`                        // a url suitable for loading the quiz in a mobile webview.  it will persiste the headless session and, for quizzes in public courses, will force the user to login.Example: http://canvas.example.edu/courses/1/quizzes/2?persist_healdess=1&force_user=1
	PreviewUrl                    string            `json:"preview_url"`                       // A url that can be visited in the browser with a POST request to preview a quiz as the teacher. Only present when the user may grade.Example: http://canvas.example.edu/courses/1/quizzes/2/take?preview=1
	Description                   string            `json:"description"`                       // the description of the quiz.Example: This is a quiz on Act 3 of Hamlet
	QuizType                      string            `json:"quiz_type"`                         // type of quiz possible values: 'practice_quiz', 'assignment', 'graded_survey', 'survey'.Example: assignment
	AssignmentGroupID             int64             `json:"assignment_group_id"`               // the ID of the quiz's assignment group:.Example: 3
	TimeLimit                     int64             `json:"time_limit"`                        // quiz time limit in minutes.Example: 5
	ShuffleAnswers                bool              `json:"shuffle_answers"`                   // shuffle answers for students?.
	HideResults                   string            `json:"hide_results"`                      // let students see their quiz responses? possible values: null, 'always', 'until_after_last_attempt'.Example: always
	ShowCorrectAnswers            bool              `json:"show_correct_answers"`              // show which answers were correct when results are shown? only valid if hide_results=null.Example: true
	ShowCorrectAnswersLastAttempt bool              `json:"show_correct_answers_last_attempt"` // restrict the show_correct_answers option above to apply only to the last submitted attempt of a quiz that allows multiple attempts. only valid if show_correct_answers=true and allowed_attempts > 1.Example: true
	ShowCorrectAnswersAt          time.Time         `json:"show_correct_answers_at"`           // when should the correct answers be visible by students? only valid if show_correct_answers=true.Example: 2013-01-23T23:59:00-07:00
	HideCorrectAnswersAt          time.Time         `json:"hide_correct_answers_at"`           // prevent the students from seeing correct answers after the specified date has passed. only valid if show_correct_answers=true.Example: 2013-01-23T23:59:00-07:00
	OneTimeResults                bool              `json:"one_time_results"`                  // prevent the students from seeing their results more than once (right after they submit the quiz).Example: true
	ScoringPolicy                 string            `json:"scoring_policy"`                    // which quiz score to keep (only if allowed_attempts != 1) possible values: 'keep_highest', 'keep_latest'.Example: keep_highest
	AllowedAttempts               int64             `json:"allowed_attempts"`                  // how many times a student can take the quiz -1 = unlimited attempts.Example: 3
	OneQuestionAtATime            bool              `json:"one_question_at_a_time"`            // show one question at a time?.
	QuestionCount                 int64             `json:"question_count"`                    // the number of questions in the quiz.Example: 12
	PointsPossible                int64             `json:"points_possible"`                   // The total point value given to the quiz.Example: 20
	CantGoBack                    bool              `json:"cant_go_back"`                      // lock questions after answering? only valid if one_question_at_a_time=true.
	AccessCode                    string            `json:"access_code"`                       // access code to restrict quiz access.Example: 2beornot2be
	IpFilter                      string            `json:"ip_filter"`                         // IP address or range that quiz access is limited to.Example: 123.123.123.123
	DueAt                         time.Time         `json:"due_at"`                            // when the quiz is due.Example: 2013-01-23T23:59:00-07:00
	LockAt                        time.Time         `json:"lock_at"`                           // when to lock the quiz.
	UnlockAt                      time.Time         `json:"unlock_at"`                         // when to unlock the quiz.Example: 2013-01-21T23:59:00-07:00
	Published                     bool              `json:"published"`                         // whether the quiz has a published or unpublished draft state..Example: true
	Unpublishable                 bool              `json:"unpublishable"`                     // Whether the assignment's 'published' state can be changed to false. Will be false if there are student submissions for the quiz..Example: true
	LockedForUser                 bool              `json:"locked_for_user"`                   // Whether or not this is locked for the user..
	LockInfo                      *LockInfo         `json:"lock_info"`                         // (Optional) Information for the user about the lock. Present when locked_for_user is true..
	LockExplanation               string            `json:"lock_explanation"`                  // (Optional) An explanation of why this is locked for the user. Present when locked_for_user is true..Example: This quiz is locked until September 1 at 12:00am
	SpeedgraderUrl                string            `json:"speedgrader_url"`                   // Link to Speed Grader for this quiz. Will not be present if quiz is unpublished.Example: http://canvas.instructure.com/courses/1/speed_grader?assignment_id=1
	QuizExtensionsUrl             string            `json:"quiz_extensions_url"`               // Link to endpoint to send extensions for this quiz..Example: http://canvas.instructure.com/courses/1/quizzes/2/quiz_extensions
	Permissions                   *QuizPermissions  `json:"permissions"`                       // Permissions the user has for the quiz.
	AllDates                      []*AssignmentDate `json:"all_dates"`                         // list of due dates for the quiz.
	VersionNumber                 int64             `json:"version_number"`                    // Current version number of the quiz.Example: 3
	QuestionTypes                 []string          `json:"question_types"`                    // List of question types in the quiz.Example: multiple_choice, essay
	AnonymousSubmissions          bool              `json:"anonymous_submissions"`             // Whether survey submissions will be kept anonymous (only applicable to 'graded_survey', 'survey' quiz types).
}

func (t *Quiz) HasError() error {
	var s []string
	s = []string{"practice_quiz", "assignment", "graded_survey", "survey"}
	if !string_utils.Include(s, t.QuizType) {
		return fmt.Errorf("expected 'quiz_type' to be one of %v", s)
	}
	s = []string{"always", "until_after_last_attempt"}
	if !string_utils.Include(s, t.HideResults) {
		return fmt.Errorf("expected 'hide_results' to be one of %v", s)
	}
	s = []string{"keep_highest", "keep_latest"}
	if !string_utils.Include(s, t.ScoringPolicy) {
		return fmt.Errorf("expected 'scoring_policy' to be one of %v", s)
	}
	return nil
}
