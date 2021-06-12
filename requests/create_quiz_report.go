package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateQuizReport Create and return a new report for this quiz. If a previously
// generated report matches the arguments and is still current (i.e.
// there have been no new submissions), it will be returned.
//
// *Responses*
//
// * <code>400 Bad Request</code> if the specified report type is invalid
// * <code>409 Conflict</code> if a quiz report of the specified type is already being
//   generated
// https://canvas.instructure.com/doc/api/quiz_reports.html
//
// Path Parameters:
// # CourseID (Required) ID
// # QuizID (Required) ID
//
// Form Parameters:
// # QuizReport (Required) . Must be one of student_analysis, item_analysisThe type of report to be generated.
// # QuizReport (Optional) Whether the report should consider all submissions or only the most
//    recent. Defaults to false, ignored for item_analysis.
// # Include (Optional) . Must be one of file, progressWhether the output should include documents for the file and/or progress
//    objects associated with this report. (Note: JSON-API only)
//
type CreateQuizReport struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		QuizID   string `json:"quiz_id"`   //  (Required)
	} `json:"path"`

	Form struct {
		QuizReport struct {
			ReportType          string `json:"report_type"`           //  (Required) . Must be one of student_analysis, item_analysis
			IncludesAllVersions bool   `json:"includes_all_versions"` //  (Optional)
		} `json:"quiz_report"`

		Include string `json:"include"` //  (Optional) . Must be one of file, progress
	} `json:"form"`
}

func (t *CreateQuizReport) GetMethod() string {
	return "POST"
}

func (t *CreateQuizReport) GetURLPath() string {
	path := "courses/{course_id}/quizzes/{quiz_id}/reports"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{quiz_id}", fmt.Sprintf("%v", t.Path.QuizID))
	return path
}

func (t *CreateQuizReport) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateQuizReport) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateQuizReport) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.QuizID == "" {
		errs = append(errs, "'QuizID' is required")
	}
	if t.Form.QuizReport.ReportType == "" {
		errs = append(errs, "'QuizReport' is required")
	}
	if !string_utils.Include([]string{"student_analysis", "item_analysis"}, t.Form.QuizReport.ReportType) {
		errs = append(errs, "QuizReport must be one of student_analysis, item_analysis")
	}
	if !string_utils.Include([]string{"file", "progress"}, t.Form.Include) {
		errs = append(errs, "Include must be one of file, progress")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateQuizReport) Do(c *canvasapi.Canvas) (*models.QuizReport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.QuizReport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
