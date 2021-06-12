package requests

import (
	"fmt"
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
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		UserID           int64 `json:"user_id"`            //  (Required)
		ExtraAttempts    int64 `json:"extra_attempts"`     //  (Optional)
		ExtraTime        int64 `json:"extra_time"`         //  (Optional)
		ManuallyUnlocked bool  `json:"manually_unlocked"`  //  (Optional)
		ExtendFromNow    int64 `json:"extend_from_now"`    //  (Optional)
		ExtendFromEndAt  int64 `json:"extend_from_end_at"` //  (Optional)
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

func (t *CourseQuizExtensionsSetExtensionsForStudentQuizSubmissions) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
