package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateOriginalityReport Create a new OriginalityReport for the specified file
// https://canvas.instructure.com/doc/api/originality_reports.html
//
// Path Parameters:
// # Path.AssignmentID (Required) ID
// # Path.SubmissionID (Required) ID
//
// Form Parameters:
// # Form.OriginalityReport.FileID (Optional) The id of the file being given an originality score. Required
//    if creating a report associated with a file.
// # Form.OriginalityReport.OriginalityScore (Required) A number between 0 and 100 representing the measure of the
//    specified file's originality.
// # Form.OriginalityReport.OriginalityReportUrl (Optional) The URL where the originality report for the specified
//    file may be found.
// # Form.OriginalityReport.OriginalityReportFileID (Optional) The ID of the file within Canvas that contains the originality
//    report for the submitted file provided in the request URL.
// # Form.OriginalityReport.ToolSetting.ResourceTypeCode (Optional) The resource type code of the resource handler Canvas should use for the
//    LTI launch for viewing originality reports. If set Canvas will launch
//    to the message with type 'basic-lti-launch-request' in the specified
//    resource handler rather than using the originality_report_url.
// # Form.OriginalityReport.ToolSetting.ResourceUrl (Optional) The URL Canvas should launch to when showing an LTI originality report.
//    Note that this value is inferred from the specified resource handler's
//    message "path" value (See `resource_type_code`) unless
//    it is specified. If this parameter is used a `resource_type_code`
//    must also be specified.
// # Form.OriginalityReport.WorkflowState (Optional) May be set to "pending", "error", or "scored". If an originality score
//    is provided a workflow state of "scored" will be inferred.
// # Form.OriginalityReport.ErrorMessage (Optional) A message describing the error. If set, the "workflow_state"
//    will be set to "error."
// # Form.OriginalityReport.Attempt (Optional) If no `file_id` is given, and no file is required for the assignment
//    (that is, the assignment allows an online text entry), this parameter
//    may be given to clarify which attempt number the report is for (in the
//    case of resubmissions). If this field is omitted and no `file_id` is
//    given, the report will be created (or updated, if it exists) for the
//    first submission attempt with no associated file.
//
type CreateOriginalityReport struct {
	Path struct {
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		SubmissionID string `json:"submission_id" url:"submission_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		OriginalityReport struct {
			FileID                  int64   `json:"file_id" url:"file_id,omitempty"`                                       //  (Optional)
			OriginalityScore        float64 `json:"originality_score" url:"originality_score,omitempty"`                   //  (Required)
			OriginalityReportUrl    string  `json:"originality_report_url" url:"originality_report_url,omitempty"`         //  (Optional)
			OriginalityReportFileID int64   `json:"originality_report_file_id" url:"originality_report_file_id,omitempty"` //  (Optional)
			ToolSetting             struct {
				ResourceTypeCode string `json:"resource_type_code" url:"resource_type_code,omitempty"` //  (Optional)
				ResourceUrl      string `json:"resource_url" url:"resource_url,omitempty"`             //  (Optional)
			} `json:"tool_setting" url:"tool_setting,omitempty"`

			WorkflowState string `json:"workflow_state" url:"workflow_state,omitempty"` //  (Optional)
			ErrorMessage  string `json:"error_message" url:"error_message,omitempty"`   //  (Optional)
			Attempt       int64  `json:"attempt" url:"attempt,omitempty"`               //  (Optional)
		} `json:"originality_report" url:"originality_report,omitempty"`
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

func (t *CreateOriginalityReport) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateOriginalityReport) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateOriginalityReport) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.SubmissionID == "" {
		errs = append(errs, "'Path.SubmissionID' is required")
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
