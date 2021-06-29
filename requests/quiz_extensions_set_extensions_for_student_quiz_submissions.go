package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// QuizExtensionsSetExtensionsForStudentQuizSubmissions <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// * <b>403 Forbidden</b> if you are not allowed to extend quizzes for this course
// https://canvas.instructure.com/doc/api/quiz_extensions.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
//
// Form Parameters:
// # Form.QuizExtensions.UserID (Required) The ID of the user we want to add quiz extensions for.
// # Form.QuizExtensions.ExtraAttempts (Optional) Number of times the student is allowed to re-take the quiz over the
//    multiple-attempt limit. This is limited to 1000 attempts or less.
// # Form.QuizExtensions.ExtraTime (Optional) The number of extra minutes to allow for all attempts. This will
//    add to the existing time limit on the submission. This is limited to
//    10080 minutes (1 week)
// # Form.QuizExtensions.ManuallyUnlocked (Optional) Allow the student to take the quiz even if it's locked for
//    everyone else.
// # Form.QuizExtensions.ExtendFromNow (Optional) The number of minutes to extend the quiz from the current time. This is
//    mutually exclusive to extend_from_end_at. This is limited to 1440
//    minutes (24 hours)
// # Form.QuizExtensions.ExtendFromEndAt (Optional) The number of minutes to extend the quiz beyond the quiz's current
//    ending time. This is mutually exclusive to extend_from_now. This is
//    limited to 1440 minutes (24 hours)
//
type QuizExtensionsSetExtensionsForStudentQuizSubmissions struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Form struct {
		QuizExtensions struct {
			UserID           []string `json:"user_id" url:"user_id,omitempty"`                       //  (Required)
			ExtraAttempts    []string `json:"extra_attempts" url:"extra_attempts,omitempty"`         //  (Optional)
			ExtraTime        []string `json:"extra_time" url:"extra_time,omitempty"`                 //  (Optional)
			ManuallyUnlocked []string `json:"manually_unlocked" url:"manually_unlocked,omitempty"`   //  (Optional)
			ExtendFromNow    []string `json:"extend_from_now" url:"extend_from_now,omitempty"`       //  (Optional)
			ExtendFromEndAt  []string `json:"extend_from_end_at" url:"extend_from_end_at,omitempty"` //  (Optional)
		} `json:"quiz_extensions" url:"quiz_extensions,omitempty"`
	} `json:"form"`
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) GetMethod() string {
	return "POST"
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/extensions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) GetQuery() (string, error) {
	return "", nil
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'Path.QuizID' is required")
	}
	if t.Form.QuizExtensions.UserID == nil {
		errs = append(errs, "'Form.QuizExtensions.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QuizExtensionsSetExtensionsForStudentQuizSubmissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
