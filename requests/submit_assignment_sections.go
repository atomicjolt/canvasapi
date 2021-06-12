package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// SubmitAssignmentSections Make a submission for an assignment. You must be enrolled as a student in
// the course/section to do this.
//
// All online turn-in submission types are supported in this API. However,
// there are a few things that are not yet supported:
//
// * Files can be submitted based on a file ID of a user or group file or through the {api:SubmissionsApiController#create_file file upload API}. However, there is no API yet for listing the user and group files.
// * Media comments can be submitted, however, there is no API yet for creating a media comment to submit.
// * Integration with Google Docs is not yet supported.
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # SectionID (Required) ID
// # AssignmentID (Required) ID
//
// Form Parameters:
// # Comment (Optional) Include a textual comment with the submission.
// # Submission (Required) . Must be one of online_text_entry, online_url, online_upload, media_recording, basic_lti_launch, student_annotationThe type of submission being made. The assignment submission_types must
//    include this submission type as an allowed option, or the submission will be rejected with a 400 error.
//
//    The submission_type given determines which of the following parameters is
//    used. For instance, to submit a URL, submission [submission_type] must be
//    set to "online_url", otherwise the submission [url] parameter will be
//    ignored.
// # Submission (Optional) Submit the assignment as an HTML document snippet. Note this HTML snippet
//    will be sanitized using the same ruleset as a submission made from the
//    Canvas web UI. The sanitized HTML will be returned in the response as the
//    submission body. Requires a submission_type of "online_text_entry".
// # Submission (Optional) Submit the assignment as a URL. The URL scheme must be "http" or "https",
//    no "ftp" or other URL schemes are allowed. If no scheme is given (e.g.
//    "www.example.com") then "http" will be assumed. Requires a submission_type
//    of "online_url" or "basic_lti_launch".
// # Submission (Optional) Submit the assignment as a set of one or more previously uploaded files
//    residing in the submitting user's files section (or the group's files
//    section, for group assignments).
//
//    To upload a new file to submit, see the submissions {api:SubmissionsApiController#create_file Upload a file API}.
//
//    Requires a submission_type of "online_upload".
// # Submission (Optional) The media comment id to submit. Media comment ids can be submitted via
//    this API, however, note that there is not yet an API to generate or list
//    existing media comments, so this functionality is currently of limited use.
//
//    Requires a submission_type of "media_recording".
// # Submission (Optional) . Must be one of audio, videoThe type of media comment being submitted.
// # Submission (Optional) Submit on behalf of the given user. Requires grading permission.
// # Submission (Optional) The Attachment ID of the document being annotated. This should match
//    the annotatable_attachment_id on the assignment.
//
//    Requires a submission_type of "student_annotation".
// # Submission (Optional) Choose the time the submission is listed as submitted at.  Requires grading permission.
//
type SubmitAssignmentSections struct {
	Path struct {
		SectionID    string `json:"section_id"`    //  (Required)
		AssignmentID string `json:"assignment_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Comment struct {
			TextComment string `json:"text_comment"` //  (Optional)
		} `json:"comment"`

		Submission struct {
			SubmissionType          string    `json:"submission_type"`           //  (Required) . Must be one of online_text_entry, online_url, online_upload, media_recording, basic_lti_launch, student_annotation
			Body                    string    `json:"body"`                      //  (Optional)
			Url                     string    `json:"url"`                       //  (Optional)
			FileIDs                 []int64   `json:"file_ids"`                  //  (Optional)
			MediaCommentID          string    `json:"media_comment_id"`          //  (Optional)
			MediaCommentType        string    `json:"media_comment_type"`        //  (Optional) . Must be one of audio, video
			UserID                  int64     `json:"user_id"`                   //  (Optional)
			AnnotatableAttachmentID int64     `json:"annotatable_attachment_id"` //  (Optional)
			SubmittedAt             time.Time `json:"submitted_at"`              //  (Optional)
		} `json:"submission"`
	} `json:"form"`
}

func (t *SubmitAssignmentSections) GetMethod() string {
	return "POST"
}

func (t *SubmitAssignmentSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	return path
}

func (t *SubmitAssignmentSections) GetQuery() (string, error) {
	return "", nil
}

func (t *SubmitAssignmentSections) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *SubmitAssignmentSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'SectionID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'AssignmentID' is required")
	}
	if t.Form.Submission.SubmissionType == "" {
		errs = append(errs, "'Submission' is required")
	}
	if !string_utils.Include([]string{"online_text_entry", "online_url", "online_upload", "media_recording", "basic_lti_launch", "student_annotation"}, t.Form.Submission.SubmissionType) {
		errs = append(errs, "Submission must be one of online_text_entry, online_url, online_upload, media_recording, basic_lti_launch, student_annotation")
	}
	if !string_utils.Include([]string{"audio", "video"}, t.Form.Submission.MediaCommentType) {
		errs = append(errs, "Submission must be one of audio, video")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SubmitAssignmentSections) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
