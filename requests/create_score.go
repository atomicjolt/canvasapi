package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateScore Create a new Result from the score params. If this is for the first created line_item for a
// resourceLinkId, or it is a line item that is not attached to a resourceLinkId, then a submission
// record will be created for the associated assignment when gradingProgress is set to
// FullyGraded or PendingManual.
//
// The submission score will also be updated when a score object is sent with either of those
// two values for gradingProgress. If a score object is sent with either of FullyGraded or
// PendingManual as the value for gradingProgress and scoreGiven is missing, the assignment
// will not be graded. This also supposes the line_item meets the condition to create a submission.
//
// A submission comment with an unknown author will be created when the comment value is included.
// This also supposes the line_item meets the condition to create a submission.
//
// NOTE: Upcoming Feature
// It will soon be possible to submit a file along with this score, which will attach the file to the
// submission that is created. Files should be formatted as Content Items, with the correct syntax
// below.
//
// Returns a url pointing to the Result. If any files were submitted, also returns the Content Items
// which were sent in the request, each with a url pointing to the Progress of the file upload.
// https://canvas.instructure.com/doc/api/score.html
//
// Path Parameters:
// # CourseID (Required) ID
// # LineItemID (Required) ID
//
// Form Parameters:
// # UserID (Required) The lti_user_id or the Canvas user_id.
//    Returns a 422 if user not found in Canvas or is not a student.
// # ActivityProgress (Required) Indicate to Canvas the status of the user towards the activity's completion.
//    Must be one of Initialized, Started, InProgress, Submitted, Completed.
// # GradingProgress (Required) Indicate to Canvas the status of the grading process.
//    A value of PendingManual will require intervention by a grader.
//    Values of NotReady, Failed, and Pending will cause the scoreGiven to be ignored.
//    FullyGraded values will require no action.
//    Possible values are NotReady, Failed, Pending, PendingManual, FullyGraded.
// # Timestamp (Required) Date and time when the score was modified in the tool. Should use subsecond precision.
//    Returns a 400 if the timestamp is earlier than the updated_at time of the Result.
// # ScoreGiven (Optional) The Current score received in the tool for this line item and user,
//    scaled to the scoreMaximum
// # ScoreMaximum (Optional) Maximum possible score for this result; it must be present if scoreGiven is present.
//    Returns 412 if not present when scoreGiven is present.
// # Comment (Optional) Comment visible to the student about this score.
// # CanvasLTISubmission (Optional) (EXTENSION) Optional submission type and data.
//    new_submission [Boolean] flag to indicate that this is a new submission. Defaults to true unless submission_type is none.
//    submission_type [String] permissible values are: none, basic_lti_launch, online_text_entry, external_tool, online_upload, or online_url. Defaults to external_tool. Ignored if content_items are provided.
//    submission_data [String] submission data (URL or body text)
//    submitted_at [String] Date and time that the submission was originally created. Should use subsecond precision. This should match the data and time that the original submission happened in Canvas.
//    content_items [Array] Files that should be included with the submission. Each item should contain `type: file`, a url pointing to the file, a title, and a progress url that Canvas can report to. If present, submission_type will be online_upload.
//
type CreateScore struct {
	Path struct {
		CourseID   string `json:"course_id"`    //  (Required)
		LineItemID string `json:"line_item_id"` //  (Required)
	} `json:"path"`

	Form struct {
		UserID              string  `json:"user_id"`                                       //  (Required)
		ActivityProgress    string  `json:"activity_progress"`                             //  (Required)
		GradingProgress     string  `json:"grading_progress"`                              //  (Required)
		Timestamp           string  `json:"timestamp"`                                     //  (Required)
		ScoreGiven          float64 `json:"score_given"`                                   //  (Optional)
		ScoreMaximum        float64 `json:"score_maximum"`                                 //  (Optional)
		Comment             string  `json:"comment"`                                       //  (Optional)
		CanvasLTISubmission string  `json:"https://canvas.instructure.com/lti/submission"` //  (Optional)
	} `json:"form"`
}

func (t *CreateScore) GetMethod() string {
	return "POST"
}

func (t *CreateScore) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items/{line_item_id}/scores"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{line_item_id}", fmt.Sprintf("%v", t.Path.LineItemID))
	return path
}

func (t *CreateScore) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateScore) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateScore) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.LineItemID == "" {
		errs = append(errs, "'LineItemID' is required")
	}
	if t.Form.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Form.ActivityProgress == "" {
		errs = append(errs, "'ActivityProgress' is required")
	}
	if t.Form.GradingProgress == "" {
		errs = append(errs, "'GradingProgress' is required")
	}
	if t.Form.Timestamp == "" {
		errs = append(errs, "'Timestamp' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateScore) Do(c *canvasapi.Canvas) ([]string, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []string{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
