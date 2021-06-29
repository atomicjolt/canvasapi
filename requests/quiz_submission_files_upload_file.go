package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
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
// # Path.CourseID (Required) ID
// # Path.QuizID (Required) ID
//
// Form Parameters:
// # Form.Name (Optional) The name of the quiz submission file
// # Form.OnDuplicate (Optional) How to handle duplicate names
//
type QuizSubmissionFilesUploadFile struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		QuizID   string `json:"quiz_id" url:"quiz_id,omitempty"`     //  (Required)
	} `json:"path"`

	Form struct {
		Name        string `json:"name" url:"name,omitempty"`                 //  (Optional)
		OnDuplicate string `json:"on_duplicate" url:"on_duplicate,omitempty"` //  (Optional)
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

func (t *QuizSubmissionFilesUploadFile) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *QuizSubmissionFilesUploadFile) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *QuizSubmissionFilesUploadFile) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'Path.QuizID' is required")
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
