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

// EditOriginalityReportFiles Modify an existing originality report. An alternative to this endpoint is
// to POST the same parameters listed below to the CREATE endpoint.
// https://canvas.instructure.com/doc/api/originality_reports.html
//
// Path Parameters:
// # AssignmentID (Required) ID
// # FileID (Required) ID
//
// Form Parameters:
// # OriginalityReport (Optional) A number between 0 and 100 representing the measure of the
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
//
type EditOriginalityReportFiles struct {
	Path struct {
		AssignmentID string `json:"assignment_id"` //  (Required)
		FileID       string `json:"file_id"`       //  (Required)
	} `json:"path"`

	Form struct {
		OriginalityReport struct {
			OriginalityScore        float64 `json:"originality_score"`          //  (Optional)
			OriginalityReportUrl    string  `json:"originality_report_url"`     //  (Optional)
			OriginalityReportFileID int64   `json:"originality_report_file_id"` //  (Optional)
			ToolSetting             struct {
				ResourceTypeCode string `json:"resource_type_code"` //  (Optional)
				ResourceUrl      string `json:"resource_url"`       //  (Optional)
			} `json:"tool_setting"`

			WorkflowState string `json:"workflow_state"` //  (Optional)
			ErrorMessage  string `json:"error_message"`  //  (Optional)
		} `json:"originality_report"`
	} `json:"form"`
}

func (t *EditOriginalityReportFiles) GetMethod() string {
	return "PUT"
}

func (t *EditOriginalityReportFiles) GetURLPath() string {
	path := "/lti/assignments/{assignment_id}/files/{file_id}/originality_report"
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{file_id}", fmt.Sprintf("%v", t.Path.FileID))
	return path
}

func (t *EditOriginalityReportFiles) GetQuery() (string, error) {
	return "", nil
}

func (t *EditOriginalityReportFiles) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditOriginalityReportFiles) HasErrors() error {
	errs := []string{}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Path.FileID == "" {
		errs = append(errs, "'FileID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditOriginalityReportFiles) Do(c *canvasapi.Canvas) (*models.OriginalityReport, error) {
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
