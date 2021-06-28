package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateAssignment Create a new assignment for this course. The assignment is created in the
// active state.
// https://canvas.instructure.com/doc/api/assignments.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Assignment (Required) The assignment name.
// # Assignment (Optional) The position of this assignment in the group when displaying
//    assignment lists.
// # Assignment (Optional) . Must be one of online_quiz, none, on_paper, discussion_topic, external_tool, online_upload, online_text_entry, online_url, media_recording, student_annotationList of supported submission types for the assignment.
//    Unless the assignment is allowing online submissions, the array should
//    only have one element.
//
//    If not allowing online submissions, your options are:
//      "online_quiz"
//      "none"
//      "on_paper"
//      "discussion_topic"
//      "external_tool"
//
//    If you are allowing online submissions, you can have one or many
//    allowed submission types:
//
//      "online_upload"
//      "online_text_entry"
//      "online_url"
//      "media_recording" (Only valid when the Kaltura plugin is enabled)
//      "student_annotation"
// # Assignment (Optional) Allowed extensions if submission_types includes "online_upload"
//
//    Example:
//      allowed_extensions: ["docx","ppt"]
// # Assignment (Optional) Only applies when the Turnitin plugin is enabled for a course and
//    the submission_types array includes "online_upload".
//    Toggles Turnitin submissions for the assignment.
//    Will be ignored if Turnitin is not available for the course.
// # Assignment (Optional) Only applies when the VeriCite plugin is enabled for a course and
//    the submission_types array includes "online_upload".
//    Toggles VeriCite submissions for the assignment.
//    Will be ignored if VeriCite is not available for the course.
// # Assignment (Optional) Settings to send along to turnitin. See Assignment object definition for
//    format.
// # Assignment (Optional) Data used for SIS integrations. Requires admin-level token with the "Manage SIS" permission. JSON string required.
// # Assignment (Optional) Unique ID from third party integrations
// # Assignment (Optional) If submission_types does not include external_tool,discussion_topic,
//    online_quiz, or on_paper, determines whether or not peer reviews
//    will be turned on for the assignment.
// # Assignment (Optional) Whether peer reviews will be assigned automatically by Canvas or if
//    teachers must manually assign peer reviews. Does not apply if peer reviews
//    are not enabled.
// # Assignment (Optional) If true, Canvas will send a notification to students in the class
//    notifying them that the content has changed.
// # Assignment (Optional) If present, the assignment will become a group assignment assigned
//    to the group.
// # Assignment (Optional) If this is a group assignment, teachers have the options to grade
//    students individually. If false, Canvas will apply the assignment's
//    score to each member of the group. If true, the teacher can manually
//    assign scores to each member of the group.
// # Assignment (Optional) Hash of external tool parameters if submission_types is ["external_tool"].
//    See Assignment object definition for format.
// # Assignment (Optional) The maximum points possible on the assignment.
// # Assignment (Optional) . Must be one of pass_fail, percent, letter_grade, gpa_scale, points, not_gradedThe strategy used for grading the assignment.
//    The assignment defaults to "points" if this field is omitted.
// # Assignment (Optional) The day/time the assignment is due. Must be between the lock dates if there are lock dates.
//    Accepts times in ISO 8601 format, e.g. 2014-10-21T18:48:00Z.
// # Assignment (Optional) The day/time the assignment is locked after. Must be after the due date if there is a due date.
//    Accepts times in ISO 8601 format, e.g. 2014-10-21T18:48:00Z.
// # Assignment (Optional) The day/time the assignment is unlocked. Must be before the due date if there is a due date.
//    Accepts times in ISO 8601 format, e.g. 2014-10-21T18:48:00Z.
// # Assignment (Optional) The assignment's description, supports HTML.
// # Assignment (Optional) The assignment group id to put the assignment in.
//    Defaults to the top assignment group in the course.
// # Assignment (Optional) List of overrides for the assignment.
// # Assignment (Optional) Whether this assignment is only visible to overrides
//    (Only useful if 'differentiated assignments' account setting is on)
// # Assignment (Optional) Whether this assignment is published.
//    (Only useful if 'draft state' account setting is on)
//    Unpublished assignments are not visible to students.
// # Assignment (Optional) The grading standard id to set for the course.  If no value is provided for this argument the current grading_standard will be un-set from this course.
//    This will update the grading_type for the course to 'letter_grade' unless it is already 'gpa_scale'.
// # Assignment (Optional) Whether this assignment is counted towards a student's final grade.
// # Assignment (Optional) Whether this assignment should use the Quizzes 2 LTI tool. Sets the
//    submission type to 'external_tool' and configures the external tool
//    attributes to use the Quizzes 2 LTI tool configured for this course.
//    Has no effect if no Quizzes 2 LTI tool is configured.
// # Assignment (Optional) Whether this assignment is moderated.
// # Assignment (Optional) The maximum number of provisional graders who may issue grades for this
//    assignment. Only relevant for moderated assignments. Must be a positive
//    value, and must be set to 1 if the course has fewer than two active
//    instructors. Otherwise, the maximum value is the number of active
//    instructors in the course minus one, or 10 if the course has more than 11
//    active instructors.
// # Assignment (Optional) The user ID of the grader responsible for choosing final grades for this
//    assignment. Only relevant for moderated assignments.
// # Assignment (Optional) Boolean indicating if provisional graders' comments are visible to other
//    provisional graders. Only relevant for moderated assignments.
// # Assignment (Optional) Boolean indicating if provisional graders' identities are hidden from
//    other provisional graders. Only relevant for moderated assignments.
// # Assignment (Optional) Boolean indicating if provisional grader identities are visible to the
//    the final grader. Only relevant for moderated assignments.
// # Assignment (Optional) Boolean indicating if the assignment is graded anonymously. If true,
//    graders cannot see student identities.
// # Assignment (Optional) The number of submission attempts allowed for this assignment. Set to -1 for unlimited attempts.
// # Assignment (Optional) The Attachment ID of the document being annotated.
//
//    Only applies when submission_types includes "student_annotation".
//
type CreateAssignment struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Assignment struct {
			Name                             string                       `json:"name" url:"name,omitempty"`                                                                   //  (Required)
			Position                         int64                        `json:"position" url:"position,omitempty"`                                                           //  (Optional)
			SubmissionTypes                  []string                     `json:"submission_types" url:"submission_types,omitempty"`                                           //  (Optional) . Must be one of online_quiz, none, on_paper, discussion_topic, external_tool, online_upload, online_text_entry, online_url, media_recording, student_annotation
			AllowedExtensions                []string                     `json:"allowed_extensions" url:"allowed_extensions,omitempty"`                                       //  (Optional)
			TurnitinEnabled                  bool                         `json:"turnitin_enabled" url:"turnitin_enabled,omitempty"`                                           //  (Optional)
			VericiteEnabled                  bool                         `json:"vericite_enabled" url:"vericite_enabled,omitempty"`                                           //  (Optional)
			TurnitinSettings                 string                       `json:"turnitin_settings" url:"turnitin_settings,omitempty"`                                         //  (Optional)
			IntegrationData                  string                       `json:"integration_data" url:"integration_data,omitempty"`                                           //  (Optional)
			IntegrationID                    string                       `json:"integration_id" url:"integration_id,omitempty"`                                               //  (Optional)
			PeerReviews                      bool                         `json:"peer_reviews" url:"peer_reviews,omitempty"`                                                   //  (Optional)
			AutomaticPeerReviews             bool                         `json:"automatic_peer_reviews" url:"automatic_peer_reviews,omitempty"`                               //  (Optional)
			NotifyOfUpdate                   bool                         `json:"notify_of_update" url:"notify_of_update,omitempty"`                                           //  (Optional)
			GroupCategoryID                  int64                        `json:"group_category_id" url:"group_category_id,omitempty"`                                         //  (Optional)
			GradeGroupStudentsIndividually   int64                        `json:"grade_group_students_individually" url:"grade_group_students_individually,omitempty"`         //  (Optional)
			ExternalToolTagAttributes        string                       `json:"external_tool_tag_attributes" url:"external_tool_tag_attributes,omitempty"`                   //  (Optional)
			PointsPossible                   float64                      `json:"points_possible" url:"points_possible,omitempty"`                                             //  (Optional)
			GradingType                      string                       `json:"grading_type" url:"grading_type,omitempty"`                                                   //  (Optional) . Must be one of pass_fail, percent, letter_grade, gpa_scale, points, not_graded
			DueAt                            time.Time                    `json:"due_at" url:"due_at,omitempty"`                                                               //  (Optional)
			LockAt                           time.Time                    `json:"lock_at" url:"lock_at,omitempty"`                                                             //  (Optional)
			UnlockAt                         time.Time                    `json:"unlock_at" url:"unlock_at,omitempty"`                                                         //  (Optional)
			Description                      string                       `json:"description" url:"description,omitempty"`                                                     //  (Optional)
			AssignmentGroupID                int64                        `json:"assignment_group_id" url:"assignment_group_id,omitempty"`                                     //  (Optional)
			AssignmentOverrides              []*models.AssignmentOverride `json:"assignment_overrides" url:"assignment_overrides,omitempty"`                                   //  (Optional)
			OnlyVisibleToOverrides           bool                         `json:"only_visible_to_overrides" url:"only_visible_to_overrides,omitempty"`                         //  (Optional)
			Published                        bool                         `json:"published" url:"published,omitempty"`                                                         //  (Optional)
			GradingStandardID                int64                        `json:"grading_standard_id" url:"grading_standard_id,omitempty"`                                     //  (Optional)
			OmitFromFinalGrade               bool                         `json:"omit_from_final_grade" url:"omit_from_final_grade,omitempty"`                                 //  (Optional)
			QuizLti                          bool                         `json:"quiz_lti" url:"quiz_lti,omitempty"`                                                           //  (Optional)
			ModeratedGrading                 bool                         `json:"moderated_grading" url:"moderated_grading,omitempty"`                                         //  (Optional)
			GraderCount                      int64                        `json:"grader_count" url:"grader_count,omitempty"`                                                   //  (Optional)
			FinalGraderID                    int64                        `json:"final_grader_id" url:"final_grader_id,omitempty"`                                             //  (Optional)
			GraderCommentsVisibleToGraders   bool                         `json:"grader_comments_visible_to_graders" url:"grader_comments_visible_to_graders,omitempty"`       //  (Optional)
			GradersAnonymousToGraders        bool                         `json:"graders_anonymous_to_graders" url:"graders_anonymous_to_graders,omitempty"`                   //  (Optional)
			GradersNamesVisibleToFinalGrader bool                         `json:"graders_names_visible_to_final_grader" url:"graders_names_visible_to_final_grader,omitempty"` //  (Optional)
			AnonymousGrading                 bool                         `json:"anonymous_grading" url:"anonymous_grading,omitempty"`                                         //  (Optional)
			AllowedAttempts                  int64                        `json:"allowed_attempts" url:"allowed_attempts,omitempty"`                                           //  (Optional)
			AnnotatableAttachmentID          int64                        `json:"annotatable_attachment_id" url:"annotatable_attachment_id,omitempty"`                         //  (Optional)
		} `json:"assignment" url:"assignment,omitempty"`
	} `json:"form"`
}

func (t *CreateAssignment) GetMethod() string {
	return "POST"
}

func (t *CreateAssignment) GetURLPath() string {
	path := "courses/{course_id}/assignments"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateAssignment) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateAssignment) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateAssignment) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateAssignment) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Assignment.Name == "" {
		errs = append(errs, "'Assignment' is required")
	}
	for _, v := range t.Form.Assignment.SubmissionTypes {
		if v != "" && !string_utils.Include([]string{"online_quiz", "none", "on_paper", "discussion_topic", "external_tool", "online_upload", "online_text_entry", "online_url", "media_recording", "student_annotation"}, v) {
			errs = append(errs, "Assignment must be one of online_quiz, none, on_paper, discussion_topic, external_tool, online_upload, online_text_entry, online_url, media_recording, student_annotation")
		}
	}
	if t.Form.Assignment.GradingType != "" && !string_utils.Include([]string{"pass_fail", "percent", "letter_grade", "gpa_scale", "points", "not_graded"}, t.Form.Assignment.GradingType) {
		errs = append(errs, "Assignment must be one of pass_fail, percent, letter_grade, gpa_scale, points, not_graded")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateAssignment) Do(c *canvasapi.Canvas) (*models.Assignment, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Assignment{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
