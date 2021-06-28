package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetFullTopicCourses Return a cached structure of the discussion topic, containing all entries,
// their authors, and their message bodies.
//
// May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403
// Forbidden status and the body 'require_initial_post'.
//
// In some rare situations, this cached structure may not be available yet. In
// that case, the server will respond with a 503 error, and the caller should
// try again soon.
//
// The response is an object containing the following keys:
// * "participants": A list of summary information on users who have posted to
//   the discussion. Each value is an object containing their id, display_name,
//   and avatar_url.
// * "unread_entries": A list of entry ids that are unread by the current
//   user. this implies that any entry not in this list is read.
// * "entry_ratings": A map of entry ids to ratings by the current user. Entries
//   not in this list have no rating. Only populated if rating is enabled.
// * "forced_entries": A list of entry ids that have forced_read_state set to
//   true. This flag is meant to indicate the entry's read_state has been
//   manually set to 'unread' by the user, so the entry should not be
//   automatically marked as read.
// * "view": A threaded view of all the entries in the discussion, containing
//   the id, user_id, and message.
// * "new_entries": Because this view is eventually consistent, it's possible
//   that newly created or updated entries won't yet be reflected in the view.
//   If the application wants to also get a flat list of all entries not yet
//   reflected in the view, pass include_new_entries=1 to the request and this
//   array of entries will be returned. These entries are returned in a flat
//   array, in ascending created_at order.
// https://canvas.instructure.com/doc/api/discussion_topics.html
//
// Path Parameters:
// # CourseID (Required) ID
// # TopicID (Required) ID
//
type GetFullTopicCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		TopicID  string `json:"topic_id" url:"topic_id,omitempty"`   //  (Required)
	} `json:"path"`
}

func (t *GetFullTopicCourses) GetMethod() string {
	return "GET"
}

func (t *GetFullTopicCourses) GetURLPath() string {
	path := "courses/{course_id}/discussion_topics/{topic_id}/view"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{topic_id}", fmt.Sprintf("%v", t.Path.TopicID))
	return path
}

func (t *GetFullTopicCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *GetFullTopicCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetFullTopicCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetFullTopicCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.TopicID == "" {
		errs = append(errs, "'TopicID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFullTopicCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
