package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// * <b>403 Forbidden</b> if you are not allowed to extend quizzes for this course
// https://canvas.instructure.com/doc/api/course_quiz_extensions.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # UserID (Required) The ID of the user we want to add quiz extensions for.
// # ExtraAttempts (Optional) Number of times the student is allowed to re-take the quiz over the
//    multiple-attempt limit. This is limited to 1000 attempts or less.
// # ExtraTime (Optional) The number of extra minutes to allow for all attempts. This will
//    add to the existing time limit on the submission. This is limited to
//    10080 minutes (1 week)
// # ManuallyUnlocked (Optional) Allow the student to take the quiz even if it's locked for
//    everyone else.
// # ExtendFromNow (Optional) The number of minutes to extend the quiz from the current time. This is
//    mutually exclusive to extend_from_end_at. This is limited to 1440
//    minutes (24 hours)
// # ExtendFromEndAt (Optional) The number of minutes to extend the quiz beyond the quiz's current
//    ending time. This is mutually exclusive to extend_from_now. This is
//    limited to 1440 minutes (24 hours)
//
type CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		UserID           int64 `json:"user_id" url:"user_id,omitempty"`                       //  (Required)
		ExtraAttempts    int64 `json:"extra_attempts" url:"extra_attempts,omitempty"`         //  (Optional)
		ExtraTime        int64 `json:"extra_time" url:"extra_time,omitempty"`                 //  (Optional)
		ManuallyUnlocked bool  `json:"manually_unlocked" url:"manually_unlocked,omitempty"`   //  (Optional)
		ExtendFromNow    int64 `json:"extend_from_now" url:"extend_from_now,omitempty"`       //  (Optional)
		ExtendFromEndAt  int64 `json:"extend_from_end_at" url:"extend_from_end_at,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) GetMethod() string {
	return "POST"
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) GetURLPath() string {
	path := "courses/{course_id}/quiz_extensions"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) GetQuery() (string, error) {
	return "", nil
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
