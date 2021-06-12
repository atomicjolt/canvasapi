package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetOutcomeImportStatusCourses Get the status of an already created Outcome import. Pass 'latest' for the outcome import id
// for the latest import.
//
//   Examples:
//     curl 'https://<canvas>/api/v1/accounts/<account_id>/outcome_imports/<outcome_import_id>' \
//         -H "Authorization: Bearer <token>"
//     curl 'https://<canvas>/api/v1/courses/<course_id>/outcome_imports/<outcome_import_id>' \
//         -H "Authorization: Bearer <token>"
// https://canvas.instructure.com/doc/api/outcome_imports.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type GetOutcomeImportStatusCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`
}

func (t *GetOutcomeImportStatusCourses) GetMethod() string {
	return "GET"
}

func (t *GetOutcomeImportStatusCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_imports/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetOutcomeImportStatusCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetOutcomeImportStatusCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetOutcomeImportStatusCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetOutcomeImportStatusCourses) Do(c *canvasapi.Canvas) (*models.OutcomeImport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeImport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
