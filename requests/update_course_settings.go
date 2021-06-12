package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateCourseSettings Can update the following course settings:
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # AllowStudentDiscussionTopics (Optional) Let students create discussion topics
// # AllowStudentForumAttachments (Optional) Let students attach files to discussions
// # AllowStudentDiscussionEditing (Optional) Let students edit or delete their own discussion posts
// # AllowStudentOrganizedGroups (Optional) Let students organize their own groups
// # FilterSpeedGraderByStudentGroup (Optional) Filter SpeedGrader to only the selected student group
// # HideFinalGrades (Optional) Hide totals in student grades summary
// # HideDistributionGraphs (Optional) Hide grade distribution graphs from students
// # HideSectionsOnCourseUsersPage (Optional) Disallow students from viewing students in sections they do not belong to
// # LockAllAnnouncements (Optional) Disable comments on announcements
// # UsageRightsRequired (Optional) Copyright and license information must be provided for files before they are published.
// # RestrictStudentPastView (Optional) Restrict students from viewing courses after end date
// # RestrictStudentFutureView (Optional) Restrict students from viewing courses before start date
// # ShowAnnouncementsOnHomePage (Optional) Show the most recent announcements on the Course home page (if a Wiki, defaults to five announcements, configurable via home_page_announcement_limit)
// # HomePageAnnouncementLimit (Optional) Limit the number of announcements on the home page if enabled via show_announcements_on_home_page
// # SyllabusCourseSummary (Optional) Show the course summary (list of assignments and calendar events) on the syllabus page. Default is true.
//
type UpdateCourseSettings struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		AllowStudentDiscussionTopics    bool  `json:"allow_student_discussion_topics"`      //  (Optional)
		AllowStudentForumAttachments    bool  `json:"allow_student_forum_attachments"`      //  (Optional)
		AllowStudentDiscussionEditing   bool  `json:"allow_student_discussion_editing"`     //  (Optional)
		AllowStudentOrganizedGroups     bool  `json:"allow_student_organized_groups"`       //  (Optional)
		FilterSpeedGraderByStudentGroup bool  `json:"filter_speed_grader_by_student_group"` //  (Optional)
		HideFinalGrades                 bool  `json:"hide_final_grades"`                    //  (Optional)
		HideDistributionGraphs          bool  `json:"hide_distribution_graphs"`             //  (Optional)
		HideSectionsOnCourseUsersPage   bool  `json:"hide_sections_on_course_users_page"`   //  (Optional)
		LockAllAnnouncements            bool  `json:"lock_all_announcements"`               //  (Optional)
		UsageRightsRequired             bool  `json:"usage_rights_required"`                //  (Optional)
		RestrictStudentPastView         bool  `json:"restrict_student_past_view"`           //  (Optional)
		RestrictStudentFutureView       bool  `json:"restrict_student_future_view"`         //  (Optional)
		ShowAnnouncementsOnHomePage     bool  `json:"show_announcements_on_home_page"`      //  (Optional)
		HomePageAnnouncementLimit       int64 `json:"home_page_announcement_limit"`         //  (Optional)
		SyllabusCourseSummary           bool  `json:"syllabus_course_summary"`              //  (Optional)
	} `json:"form"`
}

func (t *UpdateCourseSettings) GetMethod() string {
	return "PUT"
}

func (t *UpdateCourseSettings) GetURLPath() string {
	path := "courses/{course_id}/settings"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *UpdateCourseSettings) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCourseSettings) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateCourseSettings) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCourseSettings) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
