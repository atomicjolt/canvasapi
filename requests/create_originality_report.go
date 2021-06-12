package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateOriginalityReport Create a new OriginalityReport for the specified file
// https://canvas.instructure.com/doc/api/originality_reports.html
//
// Path Parameters:
// # AssignmentID (Required) ID
// # SubmissionID (Required) ID
//
// Form Parameters:
// # OriginalityReport (Optional) The id of the file being given an originality score. Required
//    if creating a report associated with a file.
// # OriginalityReport (Required) A number between 0 and 100 representing the measure of the
//    specified file's originality.
// # OriginalityReport (Optional) The URL where the originality report for the specified
//    file may be found.
// # OriginalityReport (Optional) The ID of the file within Canvas that contains the originality
//    report for the submitted file provided in the request URL.
// # OriginalityReport (Optional) The resource type code of the resource handler Canvas should use for the
//    LTI launch for viewing originality reports. If set Canvas will launch
//    to the message with type 'basic-lti-launch-request' in the specified
//    resource handler rather than using the originality_report_url.
// # OriginalityReport (Optional) The URL Canvas should launch to when showing an LTI originality report.
//    Note that this value is inferred from the specified resource handler's
//    message "path" value (See `resource_type_code`) unless
//    it is specified. If this parameter is used a `resource_type_code`
//    must also be specified.
// # OriginalityReport (Optional) May be set to "pending", "error", or "scored". If an originality score
//    is provided a workflow state of "scored" will be inferred.
// # OriginalityReport (Optional) A message describing the error. If set, the "workflow_state"
//    will be set to "error."
// # OriginalityReport (Optional) If no `file_id` is given, and no file is required for the assignment
//    (that is, the assignment allows an online text entry), this parameter
//    may be given to clarify which attempt number the report is for (in the
//    case of resubmissions). If this field is omitted and no `file_id` is
//    given, the report will be created (or updated, if it exists) for the
//    first submission attempt with no associated file.
//
type CreateOriginalityReport struct {
	Path struct {
		AssignmentID string `json:"assignment_id"` //  (Required)
		SubmissionID string `json:"submission_id"` //  (Required)
	} `json:"path"`

	Form struct {
		OriginalityReport struct {
			FileID                  int64   `json:"file_id"`                    //  (Optional)
			OriginalityScore        float64 `json:"originality_score"`          //  (Required)
			OriginalityReportUrl    string  `json:"originality_report_url"`     //  (Optional)
			OriginalityReportFileID int64   `json:"originality_report_file_id"` //  (Optional)
			ToolSetting             struct {
				ResourceTypeCode string `json:"resource_type_code"` //  (Optional)
				ResourceUrl      string `json:"resource_url"`       //  (Optional)
			} `json:"tool_setting"`

			WorkflowState string `json:"workflow_state"` //  (Optional)
			ErrorMessage  string `json:"error_message"`  //  (Optional)
			Attempt       int64  `json:"attempt"`        //  (Optional)
		} `json:"originality_report"`
	} `json:"form"`
}

func (t *CreateOriginalityReport) GetMethod() string {
	return "POST"
}

func (t *CreateOriginalityReport) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}/submissions/{submission_id}/originality_report"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{submission_id}", fmt.Sprintf("%v", t.Path.SubmissionID))
	return path
}

func (t *CreateOriginalityReport) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateOriginalityReport) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateOriginalityReport) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'SubmissionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateOriginalityReport) Do(c *canvasapi.Canvas) (*models.OriginalityReport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OriginalityReport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
