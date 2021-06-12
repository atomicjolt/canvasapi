package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// QuizSubmissionFilesUploadFile Associate a new quiz submission file
//
// This API endpoint is the first step in uploading a quiz submission file.
// See the {file:file_uploads.html File Upload Documentation} for details on
// the file upload workflow as these parameters are interpreted as per the
// documentation there.
// https://canvas.instructure.com/doc/api/quiz_submission_files.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Form Parameters:
// # Name (Optional) The name of the quiz submission file
// # OnDuplicate (Optional) How to handle duplicate names
//
type QuizSubmissionFilesUploadFile struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
	} `json:"path"`

	Form struct {
		Name        string `json:"name"`         //  (Optional)
		OnDuplicate string `json:"on_duplicate"` //  (Optional)
	} `json:"form"`
}

func (t *QuizSubmissionFilesUploadFile) GetMethod() string {
	return "POST"
}

func (t *QuizSubmissionFilesUploadFile) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/submissions/self/files"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *QuizSubmissionFilesUploadFile) GetQuery() (string, error) {
	return "", nil
}

func (t *QuizSubmissionFilesUploadFile) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *QuizSubmissionFilesUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QuizSubmissionFilesUploadFile) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
