package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// AbortGenerationOfReportOrRemovePreviouslyGeneratedOne This API allows you to cancel a previous request you issued for a report to
// be generated. Or in the case of an already generated report, you'd like to
// remove it, perhaps to generate it another time with an updated version that
// provides new features.
//
// You must check the report's generation status before attempting to use this
// interface. See the "workflow_state" property of the QuizReport's Progress
// object for more information. Only when the progress reports itself in a
// "queued" state can the generation be aborted.
//
// *Responses*
//
// - <code>204 No Content</code> if your request was accepted
// - <code>422 Unprocessable Entity</code> if the report is not being generated
//   or can not be aborted at this stage
// https://canvas.instructure.com/doc/api/quiz_reports.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
// # ID (Required) ID
//
type AbortGenerationOfReportOrRemovePreviouslyGeneratedOne struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) GetMethod() string {
	return "DELETE"
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/reports/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) GetQuery() (string, error) {
	return "", nil
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) HasErrors() error {
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
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AbortGenerationOfReportOrRemovePreviouslyGeneratedOne) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
