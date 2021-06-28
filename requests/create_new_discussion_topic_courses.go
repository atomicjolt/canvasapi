package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateNewDiscussionTopicCourses Create an new discussion topic for the course or group.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # Title (Optional) no description
// # Message (Optional) no description
// # DiscussionType (Optional) . Must be one of side_comment, threadedThe type of discussion. Defaults to side_comment if not value is given. Accepted values are 'side_comment', for discussions that only allow one level of nested comments, and 'threaded' for fully threaded discussions.
// # Published (Optional) Whether this topic is published (true) or draft state (false). Only
//    teachers and TAs have the ability to create draft state topics.
// # DelayedPostAt (Optional) If a timestamp is given, the topic will not be published until that time.
// # AllowRating (Optional) Whether or not users can rate entries in this topic.
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
// # OnlyGradersCanRate (Optional) If true, only graders will be allowed to rate entries.
// # SortByRating (Optional) If true, entries will be sorted by rating.
// # Attachment (Optional) A multipart/form-data form-field-style attachment.
//    Attachments larger than 1 kilobyte are subject to quota restrictions.
// # SpecificSections (Optional) A comma-separated list of sections ids to which the discussion topic
//    should be made specific to.  If it is not desired to make the discussion
//    topic specific to sections, then this parameter may be omitted or set to
//    "all".  Can only be present only on announcements and only those that are
//    for a course (as opposed to a group).
//
type CreateNewDiscussionTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Title                  string    `json:"title" url:"title,omitempty"`                                         //  (Optional)
		Message                string    `json:"message" url:"message,omitempty"`                                     //  (Optional)
		DiscussionType         string    `json:"discussion_type" url:"discussion_type,omitempty"`                     //  (Optional) . Must be one of side_comment, threaded
		Published              bool      `json:"published" url:"published,omitempty"`                                 //  (Optional)
		DelayedPostAt          time.Time `json:"delayed_post_at" url:"delayed_post_at,omitempty"`                     //  (Optional)
		AllowRating            bool      `json:"allow_rating" url:"allow_rating,omitempty"`                           //  (Optional)
		LockAt                 time.Time `json:"lock_at" url:"lock_at,omitempty"`                                     //  (Optional)
		PodcastEnabled         bool      `json:"podcast_enabled" url:"podcast_enabled,omitempty"`                     //  (Optional)
		PodcastHasStudentPosts bool      `json:"podcast_has_student_posts" url:"podcast_has_student_posts,omitempty"` //  (Optional)
		RequireInitialPost     bool      `json:"require_initial_post" url:"require_initial_post,omitempty"`           //  (Optional)
		Assignment             string    `json:"assignment" url:"assignment,omitempty"`                               //  (Optional)
		IsAnnouncement         bool      `json:"is_announcement" url:"is_announcement,omitempty"`                     //  (Optional)
		Pinned                 bool      `json:"pinned" url:"pinned,omitempty"`                                       //  (Optional)
		PositionAfter          string    `json:"position_after" url:"position_after,omitempty"`                       //  (Optional)
		GroupCategoryID        int64     `json:"group_category_id" url:"group_category_id,omitempty"`                 //  (Optional)
		OnlyGradersCanRate     bool      `json:"only_graders_can_rate" url:"only_graders_can_rate,omitempty"`         //  (Optional)
		SortByRating           bool      `json:"sort_by_rating" url:"sort_by_rating,omitempty"`                       //  (Optional)
		Attachment             string    `json:"attachment" url:"attachment,omitempty"`                               //  (Optional)
		SpecificSections       string    `json:"specific_sections" url:"specific_sections,omitempty"`                 //  (Optional)
	} `json:"form"`
}

func (t *CreateNewDiscussionTopicCourses) GetMethod() string {
	return "POST"
}

func (t *CreateNewDiscussionTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateNewDiscussionTopicCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateNewDiscussionTopicCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateNewDiscussionTopicCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateNewDiscussionTopicCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.DiscussionType != "" && !string_utils.Include([]string{"side_comment", "threaded"}, t.Form.DiscussionType) {
		errs = append(errs, "DiscussionType must be one of side_comment, threaded")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateNewDiscussionTopicCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
