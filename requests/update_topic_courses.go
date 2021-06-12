package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// UpdateTopicCourses Update an existing discussion topic for the course or group.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
//
// Form Parameters:
// # Title (Optional) no description
// # Message (Optional) no description
// # DiscussionType (Optional) . Must be one of side_comment, threadedThe type of discussion. Defaults to side_comment if not value is given. Accepted values are 'side_comment', for discussions that only allow one level of nested comments, and 'threaded' for fully threaded discussions.
// # Published (Optional) Whether this topic is published (true) or draft state (false). Only
//    teachers and TAs have the ability to create draft state topics.
// # DelayedPostAt (Optional) If a timestamp is given, the topic will not be published until that time.
// # LockAt (Optional) If a timestamp is given, the topic will be scheduled to lock at the
//    provided timestamp. If the timestamp is in the past, the topic will be
//    locked.
// # PodcastEnabled (Optional) If true, the topic will have an associated podcast feed.
// # PodcastHasStudentPosts (Optional) If true, the podcast will include posts from students as well. Implies
//    podcast_enabled.
// # RequireInitialPost (Optional) If true then a user may not respond to other replies until that user has
//    made an initial reply. Defaults to false.
// # Assignment (Optional) To create an assignment discussion, pass the assignment parameters as a
//    sub-object. See the {api:AssignmentsApiController#create Create an Assignment API}
//    for the available parameters. The name parameter will be ignored, as it's
//    taken from the discussion title. If you want to make a discussion that was
//    an assignment NOT an assignment, pass set_assignment = false as part of
//    the assignment object
// # IsAnnouncement (Optional) If true, this topic is an announcement. It will appear in the
//    announcement's section rather than the discussions section. This requires
//    announcment-posting permissions.
// # Pinned (Optional) If true, this topic will be listed in the "Pinned Discussion" section
// # PositionAfter (Optional) By default, discussions are sorted chronologically by creation date, you
//    can pass the id of another topic to have this one show up after the other
//    when they are listed.
// # GroupCategoryID (Optional) If present, the topic will become a group discussion assigned
//    to the group.
// # AllowRating (Optional) If true, users will be allowed to rate entries.
// # OnlyGradersCanRate (Optional) If true, only graders will be allowed to rate entries.
// # SortByRating (Optional) If true, entries will be sorted by rating.
// # SpecificSections (Optional) A comma-separated list of sections ids to which the discussion topic
//    should be made specific too.  If it is not desired to make the discussion
//    topic specific to sections, then this parameter may be omitted or set to
//    "all".  Can only be present only on announcements and only those that are
//    for a course (as opposed to a group).
//
type UpdateTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		TopicID  string `json:"topic_id"`  //  (Required)
	} `json:"path"`

	Form struct {
		Title                  string    `json:"title"`                     //  (Optional)
		Message                string    `json:"message"`                   //  (Optional)
		DiscussionType         string    `json:"discussion_type"`           //  (Optional) . Must be one of side_comment, threaded
		Published              bool      `json:"published"`                 //  (Optional)
		DelayedPostAt          time.Time `json:"delayed_post_at"`           //  (Optional)
		LockAt                 time.Time `json:"lock_at"`                   //  (Optional)
		PodcastEnabled         bool      `json:"podcast_enabled"`           //  (Optional)
		PodcastHasStudentPosts bool      `json:"podcast_has_student_posts"` //  (Optional)
		RequireInitialPost     bool      `json:"require_initial_post"`      //  (Optional)
		Assignment             string    `json:"assignment"`                //  (Optional)
		IsAnnouncement         bool      `json:"is_announcement"`           //  (Optional)
		Pinned                 bool      `json:"pinned"`                    //  (Optional)
		PositionAfter          string    `json:"position_after"`            //  (Optional)
		GroupCategoryID        int64     `json:"group_category_id"`         //  (Optional)
		AllowRating            bool      `json:"allow_rating"`              //  (Optional)
		OnlyGradersCanRate     bool      `json:"only_graders_can_rate"`     //  (Optional)
		SortByRating           bool      `json:"sort_by_rating"`            //  (Optional)
		SpecificSections       string    `json:"specific_sections"`         //  (Optional)
	} `json:"form"`
}

func (t *UpdateTopicCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *UpdateTopicCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateTopicCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateTopicCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	if !string_utils.Include([]string{"side_comment", "threaded"}, t.Form.DiscussionType) {
		errs = append(errs, "DiscussionType must be one of side_comment, threaded")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateTopicCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
