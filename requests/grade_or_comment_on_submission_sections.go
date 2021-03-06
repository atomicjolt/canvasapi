package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// GradeOrCommentOnSubmissionSections Comment on and/or update the grading for a student's assignment submission.
// If any submission or rubric_assessment arguments are provided, the user
// must have permission to manage grades in the appropriate context (course or
// section).
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # Path.SectionID (Required) ID
// # Path.AssignmentID (Required) ID
// # Path.UserID (Required) ID
//
// Form Parameters:
// # Form.Comment.TextComment (Optional) Add a textual comment to the submission.
// # Form.Comment.GroupComment (Optional) Whether or not this comment should be sent to the entire group (defaults
//    to false). Ignored if this is not a group assignment or if no text_comment
//    is provided.
// # Form.Comment.MediaCommentID (Optional) Add an audio/video comment to the submission. Media comments can be added
//    via this API, however, note that there is not yet an API to generate or
//    list existing media comments, so this functionality is currently of
//    limited use.
// # Form.Comment.MediaCommentType (Optional) . Must be one of audio, videoThe type of media comment being added.
// # Form.Comment.FileIDs (Optional) Attach files to this comment that were previously uploaded using the
//    Submission Comment API's files action
// # Form.Include.Visibility (Optional) Whether this assignment is visible to the owner of the submission
// # Form.Submission.PostedGrade (Optional) Assign a score to the submission, updating both the "score" and "grade"
//    fields on the submission record. This parameter can be passed in a few
//    different formats:
//
//    points:: A floating point or integral value, such as "13.5". The grade
//      will be interpreted directly as the score of the assignment.
//      Values above assignment.points_possible are allowed, for awarding
//      extra credit.
//    percentage:: A floating point value appended with a percent sign, such as
//       "40%". The grade will be interpreted as a percentage score on the
//       assignment, where 100% == assignment.points_possible. Values above 100%
//       are allowed, for awarding extra credit.
//    letter grade:: A letter grade, following the assignment's defined letter
//       grading scheme. For example, "A-". The resulting score will be the high
//       end of the defined range for the letter grade. For instance, if "B" is
//       defined as 86% to 84%, a letter grade of "B" will be worth 86%. The
//       letter grade will be rejected if the assignment does not have a defined
//       letter grading scheme. For more fine-grained control of scores, pass in
//       points or percentage rather than the letter grade.
//    "pass/complete/fail/incomplete":: A string value of "pass" or "complete"
//       will give a score of 100%. "fail" or "incomplete" will give a score of
//       0.
//
//    Note that assignments with grading_type of "pass_fail" can only be
//    assigned a score of 0 or assignment.points_possible, nothing inbetween. If
//    a posted_grade in the "points" or "percentage" format is sent, the grade
//    will only be accepted if the grade equals one of those two values.
// # Form.Submission.Excuse (Optional) Sets the "excused" status of an assignment.
// # Form.Submission.LatePolicyStatus (Optional) Sets the late policy status to either "late", "missing", "none", or null.
// # Form.Submission.SecondsLateOverride (Optional) Sets the seconds late if late policy status is "late"
// # Form.RubricAssessment (Optional) Assign a rubric assessment to this assignment submission. The
//    sub-parameters here depend on the rubric for the assignment. The general
//    format is, for each row in the rubric:
//
//    The points awarded for this row.
//      rubric_assessment[criterion_id][points]
//
//    The rating id for the row.
//      rubric_assessment[criterion_id][rating_id]
//
//    Comments to add for this row.
//      rubric_assessment[criterion_id][comments]
//
//    For example, if the assignment rubric is (in JSON format):
//      !!!javascript
//      [
//        {
//          'id': 'crit1',
//          'points': 10,
//          'description': 'Criterion 1',
//          'ratings':
//          [
//            { 'id': 'rat1', 'description': 'Good', 'points': 10 },
//            { 'id': 'rat2', 'description': 'Poor', 'points': 3 }
//          ]
//        },
//        {
//          'id': 'crit2',
//          'points': 5,
//          'description': 'Criterion 2',
//          'ratings':
//          [
//            { 'id': 'rat1', 'description': 'Exemplary', 'points': 5 },
//            { 'id': 'rat2', 'description': 'Complete', 'points': 5 },
//            { 'id': 'rat3', 'description': 'Incomplete', 'points': 0 }
//          ]
//        }
//      ]
//
//    Then a possible set of values for rubric_assessment would be:
//        rubric_assessment[crit1][points]=3&rubric_assessment[crit1][rating_id]=rat1&rubric_assessment[crit2][points]=5&rubric_assessment[crit2][rating_id]=rat2&rubric_assessment[crit2][comments]=Well%20Done.
//
type GradeOrCommentOnSubmissionSections struct {
	Path struct {
		SectionID    string `json:"section_id" url:"section_id,omitempty"`       //  (Required)
		AssignmentID string `json:"assignment_id" url:"assignment_id,omitempty"` //  (Required)
		UserID       string `json:"user_id" url:"user_id,omitempty"`             //  (Required)
	} `json:"path"`

	Form struct {
		Comment struct {
			TextComment      string   `json:"text_comment" url:"text_comment,omitempty"`             //  (Optional)
			GroupComment     bool     `json:"group_comment" url:"group_comment,omitempty"`           //  (Optional)
			MediaCommentID   string   `json:"media_comment_id" url:"media_comment_id,omitempty"`     //  (Optional)
			MediaCommentType string   `json:"media_comment_type" url:"media_comment_type,omitempty"` //  (Optional) . Must be one of audio, video
			FileIDs          []string `json:"file_ids" url:"file_ids,omitempty"`                     //  (Optional)
		} `json:"comment" url:"comment,omitempty"`

		Include struct {
			Visibility string `json:"visibility" url:"visibility,omitempty"` //  (Optional)
		} `json:"include" url:"include,omitempty"`

		Submission struct {
			PostedGrade         string `json:"posted_grade" url:"posted_grade,omitempty"`                   //  (Optional)
			Excuse              bool   `json:"excuse" url:"excuse,omitempty"`                               //  (Optional)
			LatePolicyStatus    string `json:"late_policy_status" url:"late_policy_status,omitempty"`       //  (Optional)
			SecondsLateOverride int64  `json:"seconds_late_override" url:"seconds_late_override,omitempty"` //  (Optional)
		} `json:"submission" url:"submission,omitempty"`

		RubricAssessment *models.RubricAssessment `json:"rubric_assessment" url:"rubric_assessment,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *GradeOrCommentOnSubmissionSections) GetMethod() string {
	return "PUT"
}

func (t *GradeOrCommentOnSubmissionSections) GetURLPath() string {
	path := "sections/{section_id}/assignments/{assignment_id}/submissions/{user_id}"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	path = strings.ReplaceAll(path, "{assignment_id}", fmt.Sprintf("%v", t.Path.AssignmentID))
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GradeOrCommentOnSubmissionSections) GetQuery() (string, error) {
	return "", nil
}

func (t *GradeOrCommentOnSubmissionSections) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *GradeOrCommentOnSubmissionSections) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *GradeOrCommentOnSubmissionSections) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'Path.SectionID' is required")
	}
	if t.Path.AssignmentID == "" {
		errs = append(errs, "'Path.AssignmentID' is required")
	}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Form.Comment.MediaCommentType != "" && !string_utils.Include([]string{"audio", "video"}, t.Form.Comment.MediaCommentType) {
		errs = append(errs, "Comment must be one of audio, video")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GradeOrCommentOnSubmissionSections) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
