package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Quiz struct {
	ID                            int64             `json:"id" url:"id,omitempty"`                                                               // the ID of the quiz.Example: 5
	Title                         string            `json:"title" url:"title,omitempty"`                                                         // the title of the quiz.Example: Hamlet Act 3 Quiz
	HtmlUrl                       string            `json:"html_url" url:"html_url,omitempty"`                                                   // the HTTP/HTTPS URL to the quiz.Example: http://canvas.example.edu/courses/1/quizzes/2
	MobileUrl                     string            `json:"mobile_url" url:"mobile_url,omitempty"`                                               // a url suitable for loading the quiz in a mobile webview.  it will persiste the headless session and, for quizzes in public courses, will force the user to login.Example: http://canvas.example.edu/courses/1/quizzes/2?persist_healdess=1&force_user=1
	PreviewUrl                    string            `json:"preview_url" url:"preview_url,omitempty"`                                             // A url that can be visited in the browser with a POST request to preview a quiz as the teacher. Only present when the user may grade.Example: http://canvas.example.edu/courses/1/quizzes/2/take?preview=1
	Description                   string            `json:"description" url:"description,omitempty"`                                             // the description of the quiz.Example: This is a quiz on Act 3 of Hamlet
	QuizType                      string            `json:"quiz_type" url:"quiz_type,omitempty"`                                                 // type of quiz possible values: 'practice_quiz', 'assignment', 'graded_survey', 'survey'.Example: assignment
	AssignmentGroupID             int64             `json:"assignment_group_id" url:"assignment_group_id,omitempty"`                             // the ID of the quiz's assignment group:.Example: 3
	TimeLimit                     int64             `json:"time_limit" url:"time_limit,omitempty"`                                               // quiz time limit in minutes.Example: 5
	ShuffleAnswers                bool              `json:"shuffle_answers" url:"shuffle_answers,omitempty"`                                     // shuffle answers for students?.
	HideResults                   string            `json:"hide_results" url:"hide_results,omitempty"`                                           // let students see their quiz responses? possible values: null, 'always', 'until_after_last_attempt'.Example: always
	ShowCorrectAnswers            bool              `json:"show_correct_answers" url:"show_correct_answers,omitempty"`                           // show which answers were correct when results are shown? only valid if hide_results=null.Example: true
	ShowCorrectAnswersLastAttempt bool              `json:"show_correct_answers_last_attempt" url:"show_correct_answers_last_attempt,omitempty"` // restrict the show_correct_answers option above to apply only to the last submitted attempt of a quiz that allows multiple attempts. only valid if show_correct_answers=true and allowed_attempts > 1.Example: true
	ShowCorrectAnswersAt          time.Time         `json:"show_correct_answers_at" url:"show_correct_answers_at,omitempty"`                     // when should the correct answers be visible by students? only valid if show_correct_answers=true.Example: 2013-01-23T23:59:00-07:00
	HideCorrectAnswersAt          time.Time         `json:"hide_correct_answers_at" url:"hide_correct_answers_at,omitempty"`                     // prevent the students from seeing correct answers after the specified date has passed. only valid if show_correct_answers=true.Example: 2013-01-23T23:59:00-07:00
	OneTimeResults                bool              `json:"one_time_results" url:"one_time_results,omitempty"`                                   // prevent the students from seeing their results more than once (right after they submit the quiz).Example: true
	ScoringPolicy                 string            `json:"scoring_policy" url:"scoring_policy,omitempty"`                                       // which quiz score to keep (only if allowed_attempts != 1) possible values: 'keep_highest', 'keep_latest'.Example: keep_highest
	AllowedAttempts               int64             `json:"allowed_attempts" url:"allowed_attempts,omitempty"`                                   // how many times a student can take the quiz -1 = unlimited attempts.Example: 3
	OneQuestionAtATime            bool              `json:"one_question_at_a_time" url:"one_question_at_a_time,omitempty"`                       // show one question at a time?.
	QuestionCount                 int64             `json:"question_count" url:"question_count,omitempty"`                                       // the number of questions in the quiz.Example: 12
	PointsPossible                float64           `json:"points_possible" url:"points_possible,omitempty"`                                     // The total point value given to the quiz.Example: 20
	CantGoBack                    bool              `json:"cant_go_back" url:"cant_go_back,omitempty"`                                           // lock questions after answering? only valid if one_question_at_a_time=true.
	AccessCode                    string            `json:"access_code" url:"access_code,omitempty"`                                             // access code to restrict quiz access.Example: 2beornot2be
	IpFilter                      string            `json:"ip_filter" url:"ip_filter,omitempty"`                                                 // IP address or range that quiz access is limited to.Example: 123.123.123.123
	DueAt                         time.Time         `json:"due_at" url:"due_at,omitempty"`                                                       // when the quiz is due.Example: 2013-01-23T23:59:00-07:00
	LockAt                        time.Time         `json:"lock_at" url:"lock_at,omitempty"`                                                     // when to lock the quiz.
	UnlockAt                      time.Time         `json:"unlock_at" url:"unlock_at,omitempty"`                                                 // when to unlock the quiz.Example: 2013-01-21T23:59:00-07:00
	Published                     bool              `json:"published" url:"published,omitempty"`                                                 // whether the quiz has a published or unpublished draft state..Example: true
	Unpublishable                 bool              `json:"unpublishable" url:"unpublishable,omitempty"`                                         // Whether the assignment's 'published' state can be changed to false. Will be false if there are student submissions for the quiz..Example: true
	LockedForUser                 bool              `json:"locked_for_user" url:"locked_for_user,omitempty"`                                     // Whether or not this is locked for the user..
	LockInfo                      *LockInfo         `json:"lock_info" url:"lock_info,omitempty"`                                                 // (Optional) Information for the user about the lock. Present when locked_for_user is true..
	LockExplanation               string            `json:"lock_explanation" url:"lock_explanation,omitempty"`                                   // (Optional) An explanation of why this is locked for the user. Present when locked_for_user is true..Example: This quiz is locked until September 1 at 12:00am
	SpeedgraderUrl                string            `json:"speedgrader_url" url:"speedgrader_url,omitempty"`                                     // Link to Speed Grader for this quiz. Will not be present if quiz is unpublished.Example: http://canvas.instructure.com/courses/1/speed_grader?assignment_id=1
	QuizExtensionsUrl             string            `json:"quiz_extensions_url" url:"quiz_extensions_url,omitempty"`                             // Link to endpoint to send extensions for this quiz..Example: http://canvas.instructure.com/courses/1/quizzes/2/quiz_extensions
	Permissions                   *QuizPermissions  `json:"permissions" url:"permissions,omitempty"`                                             // Permissions the user has for the quiz.
	AllDates                      []*AssignmentDate `json:"all_dates" url:"all_dates,omitempty"`                                                 // list of due dates for the quiz.
	VersionNumber                 int64             `json:"version_number" url:"version_number,omitempty"`                                       // Current version number of the quiz.Example: 3
	QuestionTypes                 []string          `json:"question_types" url:"question_types,omitempty"`                                       // List of question types in the quiz.Example: multiple_choice, essay
	AnonymousSubmissions          bool              `json:"anonymous_submissions" url:"anonymous_submissions,omitempty"`                         // Whether survey submissions will be kept anonymous (only applicable to 'graded_survey', 'survey' quiz types).
}

func (t *Quiz) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"practice_quiz", "assignment", "graded_survey", "survey"}
	if t.QuizType != "" && !string_utils.Include(s, t.QuizType) {
		errs = append(errs, fmt.Sprintf("expected 'QuizType' to be one of %v", s))
	}
	s = []string{"always", "until_after_last_attempt"}
	if t.HideResults != "" && !string_utils.Include(s, t.HideResults) {
		errs = append(errs, fmt.Sprintf("expected 'HideResults' to be one of %v", s))
	}
	s = []string{"keep_highest", "keep_latest"}
	if t.ScoringPolicy != "" && !string_utils.Include(s, t.ScoringPolicy) {
		errs = append(errs, fmt.Sprintf("expected 'ScoringPolicy' to be one of %v", s))
	}
	return nil
}
