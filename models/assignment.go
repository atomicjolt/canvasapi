package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Assignment struct {
	ID                              int64                      `json:"id"`                                   // the ID of the assignment.Example: 4
	Name                            string                     `json:"name"`                                 // the name of the assignment.Example: some assignment
	Description                     string                     `json:"description"`                          // the assignment description, in an HTML fragment.Example: <p>Do the following:</p>.
	CreatedAt                       time.Time                  `json:"created_at"`                           // The time at which this assignment was originally created.Example: 2012-07-01T23:59:00-06:00
	UpdatedAt                       time.Time                  `json:"updated_at"`                           // The time at which this assignment was last modified in any way.Example: 2012-07-01T23:59:00-06:00
	DueAt                           time.Time                  `json:"due_at"`                               // the due date for the assignment. returns null if not present. NOTE: If this assignment has assignment overrides, this field will be the due date as it applies to the user requesting information from the API..Example: 2012-07-01T23:59:00-06:00
	LockAt                          time.Time                  `json:"lock_at"`                              // the lock date (assignment is locked after this date). returns null if not present. NOTE: If this assignment has assignment overrides, this field will be the lock date as it applies to the user requesting information from the API..Example: 2012-07-01T23:59:00-06:00
	UnlockAt                        time.Time                  `json:"unlock_at"`                            // the unlock date (assignment is unlocked after this date) returns null if not present NOTE: If this assignment has assignment overrides, this field will be the unlock date as it applies to the user requesting information from the API..Example: 2012-07-01T23:59:00-06:00
	HasOverrides                    bool                       `json:"has_overrides"`                        // whether this assignment has overrides.Example: true
	AllDates                        []*AssignmentDate          `json:"all_dates"`                            // (Optional) all dates associated with the assignment, if applicable.
	CourseID                        int64                      `json:"course_id"`                            // the ID of the course the assignment belongs to.Example: 123
	HtmlUrl                         string                     `json:"html_url"`                             // the URL to the assignment's web page.Example: https://.
	SubmissionsDownloadUrl          string                     `json:"submissions_download_url"`             // the URL to download all submissions as a zip.Example: https://example.com/courses/:course_id/assignments/:id/submissions?zip=1
	AssignmentGroupID               int64                      `json:"assignment_group_id"`                  // the ID of the assignment's group.Example: 2
	DueDateRequired                 bool                       `json:"due_date_required"`                    // Boolean flag indicating whether the assignment requires a due date based on the account level setting.Example: true
	AllowedExtensions               []string                   `json:"allowed_extensions"`                   // Allowed file extensions, which take effect if submission_types includes 'online_upload'..Example: docx, ppt
	MaxNameLength                   int64                      `json:"max_name_length"`                      // An integer indicating the maximum length an assignment's name may be.Example: 15
	TurnitinEnabled                 bool                       `json:"turnitin_enabled"`                     // Boolean flag indicating whether or not Turnitin has been enabled for the assignment. NOTE: This flag will not appear unless your account has the Turnitin plugin available.Example: true
	VericiteEnabled                 bool                       `json:"vericite_enabled"`                     // Boolean flag indicating whether or not VeriCite has been enabled for the assignment. NOTE: This flag will not appear unless your account has the VeriCite plugin available.Example: true
	TurnitinSettings                *TurnitinSettings          `json:"turnitin_settings"`                    // Settings to pass along to turnitin to control what kinds of matches should be considered. originality_report_visibility can be 'immediate', 'after_grading', 'after_due_date', or 'never' exclude_small_matches_type can be null, 'percent', 'words' exclude_small_matches_value: - if type is null, this will be null also - if type is 'percent', this will be a number between 0 and 100 representing match size to exclude as a percentage of the document size. - if type is 'words', this will be number > 0 representing how many words a match must contain for it to be considered NOTE: This flag will not appear unless your account has the Turnitin plugin available.
	GradeGroupStudentsIndividually  bool                       `json:"grade_group_students_individually"`    // If this is a group assignment, boolean flag indicating whether or not students will be graded individually..
	ExternalToolTagAttributes       *ExternalToolTagAttributes `json:"external_tool_tag_attributes"`         // (Optional) assignment's settings for external tools if submission_types include 'external_tool'. Only url and new_tab are included (new_tab defaults to false).  Use the 'External Tools' API if you need more information about an external tool..
	PeerReviews                     bool                       `json:"peer_reviews"`                         // Boolean indicating if peer reviews are required for this assignment.
	AutomaticPeerReviews            bool                       `json:"automatic_peer_reviews"`               // Boolean indicating peer reviews are assigned automatically. If false, the teacher is expected to manually assign peer reviews..
	PeerReviewCount                 int64                      `json:"peer_review_count"`                    // Integer representing the amount of reviews each user is assigned. NOTE: This key is NOT present unless you have automatic_peer_reviews set to true..Example: 0
	PeerReviewsAssignAt             time.Time                  `json:"peer_reviews_assign_at"`               // String representing a date the reviews are due by. Must be a date that occurs after the default due date. If blank, or date is not after the assignment's due date, the assignment's due date will be used. NOTE: This key is NOT present unless you have automatic_peer_reviews set to true..Example: 2012-07-01T23:59:00-06:00
	IntraGroupPeerReviews           bool                       `json:"intra_group_peer_reviews"`             // Boolean representing whether or not members from within the same group on a group assignment can be assigned to peer review their own group's work.Example: false
	GroupCategoryID                 int64                      `json:"group_category_id"`                    // The ID of the assignment’s group set, if this is a group assignment. For group discussions, set group_category_id on the discussion topic, not the linked assignment..Example: 1
	NeedsGradingCount               int64                      `json:"needs_grading_count"`                  // if the requesting user has grading rights, the number of submissions that need grading..Example: 17
	NeedsGradingCountBySection      []*NeedsGradingCount       `json:"needs_grading_count_by_section"`       // if the requesting user has grading rights and the 'needs_grading_count_by_section' flag is specified, the number of submissions that need grading split out by section. NOTE: This key is NOT present unless you pass the 'needs_grading_count_by_section' argument as true.  ANOTHER NOTE: it's possible to be enrolled in multiple sections, and if a student is setup that way they will show an assignment that needs grading in multiple sections (effectively the count will be duplicated between sections).Example: {'section_id'=>'123456', 'needs_grading_count'=>5}, {'section_id'=>'654321', 'needs_grading_count'=>0}
	Position                        int64                      `json:"position"`                             // the sorting order of the assignment in the group.Example: 1
	PostToSIS                       bool                       `json:"post_to_sis"`                          // (optional, present if Sync Grades to SIS feature is enabled).Example: true
	IntegrationID                   string                     `json:"integration_id"`                       // (optional, Third Party unique identifier for Assignment).Example: 12341234
	IntegrationData                 string                     `json:"integration_data"`                     // (optional, Third Party integration data for assignment).Example: 0954
	PointsPossible                  float64                    `json:"points_possible"`                      // the maximum points possible for the assignment.Example: 12.0
	SubmissionTypes                 string                     `json:"submission_types"`                     // the types of submissions allowed for this assignment list containing one or more of the following: 'discussion_topic', 'online_quiz', 'on_paper', 'none', 'external_tool', 'online_text_entry', 'online_url', 'online_upload', 'media_recording', 'student_annotation'.Example: online_text_entry
	HasSubmittedSubmissions         bool                       `json:"has_submitted_submissions"`            // If true, the assignment has been submitted to by at least one student.Example: true
	GradingType                     string                     `json:"grading_type"`                         // The type of grading the assignment receives; one of 'pass_fail', 'percent', 'letter_grade', 'gpa_scale', 'points'.Example: points
	GradingStandardID               int64                      `json:"grading_standard_id"`                  // The id of the grading standard being applied to this assignment. Valid if grading_type is 'letter_grade' or 'gpa_scale'..
	Published                       bool                       `json:"published"`                            // Whether the assignment is published.Example: true
	Unpublishable                   bool                       `json:"unpublishable"`                        // Whether the assignment's 'published' state can be changed to false. Will be false if there are student submissions for the assignment..
	OnlyVisibleToOverrides          bool                       `json:"only_visible_to_overrides"`            // Whether the assignment is only visible to overrides..
	LockedForUser                   bool                       `json:"locked_for_user"`                      // Whether or not this is locked for the user..
	LockInfo                        *LockInfo                  `json:"lock_info"`                            // (Optional) Information for the user about the lock. Present when locked_for_user is true..
	LockExplanation                 string                     `json:"lock_explanation"`                     // (Optional) An explanation of why this is locked for the user. Present when locked_for_user is true..Example: This assignment is locked until September 1 at 12:00am
	QuizID                          int64                      `json:"quiz_id"`                              // (Optional) id of the associated quiz (applies only when submission_types is ['online_quiz']).Example: 620
	AnonymousSubmissions            bool                       `json:"anonymous_submissions"`                // (Optional) whether anonymous submissions are accepted (applies only to quiz assignments).
	DiscussionTopic                 *DiscussionTopic           `json:"discussion_topic"`                     // (Optional) the DiscussionTopic associated with the assignment, if applicable.
	FreezeOnCopy                    bool                       `json:"freeze_on_copy"`                       // (Optional) Boolean indicating if assignment will be frozen when it is copied. NOTE: This field will only be present if the AssignmentFreezer plugin is available for your account..
	Frozen                          bool                       `json:"frozen"`                               // (Optional) Boolean indicating if assignment is frozen for the calling user. NOTE: This field will only be present if the AssignmentFreezer plugin is available for your account..
	FrozenAttributes                []string                   `json:"frozen_attributes"`                    // (Optional) Array of frozen attributes for the assignment. Only account administrators currently have permission to change an attribute in this list. Will be empty if no attributes are frozen for this assignment. Possible frozen attributes are: title, description, lock_at, points_possible, grading_type, submission_types, assignment_group_id, allowed_extensions, group_category_id, notify_of_update, peer_reviews NOTE: This field will only be present if the AssignmentFreezer plugin is available for your account..Example: title
	Submission                      *Submission                `json:"submission"`                           // (Optional) If 'submission' is included in the 'include' parameter, includes a Submission object that represents the current user's (user who is requesting information from the api) current submission for the assignment. See the Submissions API for an example response. If the user does not have a submission, this key will be absent..
	UseRubricForGrading             bool                       `json:"use_rubric_for_grading"`               // (Optional) If true, the rubric is directly tied to grading the assignment. Otherwise, it is only advisory. Included if there is an associated rubric..Example: true
	RubricSettings                  string                     `json:"rubric_settings"`                      // (Optional) An object describing the basic attributes of the rubric, including the point total. Included if there is an associated rubric..Example: {'points_possible'=>12}
	Rubric                          []*RubricCriteria          `json:"rubric"`                               // (Optional) A list of scoring criteria and ratings for each rubric criterion. Included if there is an associated rubric..
	AssignmentVisibility            []int64                    `json:"assignment_visibility"`                // (Optional) If 'assignment_visibility' is included in the 'include' parameter, includes an array of student IDs who can see this assignment..Example: 137, 381, 572
	Overrides                       []*AssignmentOverride      `json:"overrides"`                            // (Optional) If 'overrides' is included in the 'include' parameter, includes an array of assignment override objects..
	OmitFromFinalGrade              bool                       `json:"omit_from_final_grade"`                // (Optional) If true, the assignment will be omitted from the student's final grade.Example: true
	ModeratedGrading                bool                       `json:"moderated_grading"`                    // Boolean indicating if the assignment is moderated..Example: true
	GraderCount                     int64                      `json:"grader_count"`                         // The maximum number of provisional graders who may issue grades for this assignment. Only relevant for moderated assignments. Must be a positive value, and must be set to 1 if the course has fewer than two active instructors. Otherwise, the maximum value is the number of active instructors in the course minus one, or 10 if the course has more than 11 active instructors..Example: 3
	FinalGraderID                   int64                      `json:"final_grader_id"`                      // The user ID of the grader responsible for choosing final grades for this assignment. Only relevant for moderated assignments..Example: 3
	GraderCommentsVisibleToGraders  bool                       `json:"grader_comments_visible_to_graders"`   // Boolean indicating if provisional graders' comments are visible to other provisional graders. Only relevant for moderated assignments..Example: true
	GradersAnonymousToGraders       bool                       `json:"graders_anonymous_to_graders"`         // Boolean indicating if provisional graders' identities are hidden from other provisional graders. Only relevant for moderated assignments with grader_comments_visible_to_graders set to true..Example: true
	GraderNamesVisibleToFinalGrader bool                       `json:"grader_names_visible_to_final_grader"` // Boolean indicating if provisional grader identities are visible to the final grader. Only relevant for moderated assignments..Example: true
	AnonymousGrading                bool                       `json:"anonymous_grading"`                    // Boolean indicating if the assignment is graded anonymously. If true, graders cannot see student identities..Example: true
	AllowedAttempts                 int64                      `json:"allowed_attempts"`                     // The number of submission attempts a student can make for this assignment. -1 is considered unlimited..Example: 2
	PostManually                    bool                       `json:"post_manually"`                        // Whether the assignment has manual posting enabled. Only relevant for courses using New Gradebook..Example: true
	ScoreStatistics                 *ScoreStatistic            `json:"score_statistics"`                     // (Optional) If 'score_statistics' and 'submission' are included in the 'include' parameter and statistics are available, includes the min, max, and mode for this assignment.
	CanSubmit                       bool                       `json:"can_submit"`                           // (Optional) If retrieving a single assignment and 'can_submit' is included in the 'include' parameter, flags whether user has the right to submit the assignment (i.e. checks enrollment dates, submission types, locked status, attempts remaining, etc...). Including 'can submit' automatically includes 'submission' in the include parameter. Not available when observed_users are included..Example: true
}

func (t *Assignment) HasError() error {
	var s []string
	s = []string{"discussion_topic", "online_quiz", "on_paper", "not_graded", "none", "external_tool", "online_text_entry", "online_url", "online_upload", "media_recording", "student_annotation"}
	if !string_utils.Include(s, t.SubmissionTypes) {
		return fmt.Errorf("expected 'submission_types' to be one of %v", s)
	}
	s = []string{"pass_fail", "percent", "letter_grade", "gpa_scale", "points"}
	if !string_utils.Include(s, t.GradingType) {
		return fmt.Errorf("expected 'grading_type' to be one of %v", s)
	}
	return nil
}
