package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateStudentQuestionScoresAndComments Update the amount of points a student has scored for questions they've
// answered, provide comments for the student about their answer(s), or simply
// fudge the total score by a specific amount of points.
//
// <b>Responses</b>
//
// * <b>200 OK</b> if the request was successful
// * <b>403 Forbidden</b> if you are not a teacher in this course
// * <b>400 Bad Request</b> if the attempt parameter is missing or invalid
// * <b>400 Bad Request</b> if the specified QS attempt is not yet complete
// https://canvas.instructure.com/doc/api/quiz_submissions.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # QuizSubmissions (Required) The attempt number of the quiz submission that should be updated. This
//    attempt MUST be already completed.
// # QuizSubmissions (Optional) Amount of positive or negative points to fudge the total score by.
// # QuizSubmissions (Optional) A set of scores and comments for each question answered by the student.
//    The keys are the question IDs, and the values are hashes of `score` and
//    `comment` entries. See {Appendix: Manual Scoring} for more on this
//    parameter.
//
type UpdateStudentQuestionScoresAndComments struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		QuizSubmissions struct {
			Attempt     []int64                            `json:"attempt"`      //  (Required)
			FudgePoints []float64                          `json:"fudge_points"` //  (Optional)
			Questions   map[string]QuizSubmissionOverrides `json:"questions"`    //  (Optional)
		} `json:"quiz_submissions"`
	} `json:"form"`
}

func (t *UpdateStudentQuestionScoresAndComments) GetMethod() string {
	return "PUT"
}

func (t *UpdateStudentQuestionScoresAndComments) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateStudentQuestionScoresAndComments) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateStudentQuestionScoresAndComments) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateStudentQuestionScoresAndComments) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.QuizSubmissions.Attempt == nil {
		errs = append(errs, "'QuizSubmissions' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateStudentQuestionScoresAndComments) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}

type QuizSubmissionOverrides struct {
	Score   string `json:"score"`   //  (Optional)
	Comment string `json:"comment"` //  (Optional)
}
