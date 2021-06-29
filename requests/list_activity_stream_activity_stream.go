package requests

import (
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ListActivityStreamActivityStream Returns the current user's global activity stream, paginated.
//
// There are many types of objects that can be returned in the activity
// stream. All object types have the same basic set of shared attributes:
//   !!!javascript
//   {
//     'created_at': '2011-07-13T09:12:00Z',
//     'updated_at': '2011-07-25T08:52:41Z',
//     'id': 1234,
//     'title': 'Stream Item Subject',
//     'message': 'This is the body text of the activity stream item. It is plain-text, and can be multiple paragraphs.',
//     'type': 'DiscussionTopic|Conversation|Message|Submission|Conference|Collaboration|AssessmentRequest...',
//     'read_state': false,
//     'context_type': 'course', // course|group
//     'course_id': 1,
//     'group_id': null,
//     'html_url': "http://..." // URL to the Canvas web UI for this stream item
//   }
//
// In addition, each item type has its own set of attributes available.
//
// DiscussionTopic:
//
//   !!!javascript
//   {
//     'type': 'DiscussionTopic',
//     'discussion_topic_id': 1234,
//     'total_root_discussion_entries': 5,
//     'require_initial_post': true,
//     'user_has_posted': true,
//     'root_discussion_entries': {
//       ...
//     }
//   }
//
// For DiscussionTopic, the message is truncated at 4kb.
//
// Announcement:
//
//   !!!javascript
//   {
//     'type': 'Announcement',
//     'announcement_id': 1234,
//     'total_root_discussion_entries': 5,
//     'require_initial_post': true,
//     'user_has_posted': null,
//     'root_discussion_entries': {
//       ...
//     }
//   }
//
// For Announcement, the message is truncated at 4kb.
//
// Conversation:
//
//   !!!javascript
//   {
//     'type': 'Conversation',
//     'conversation_id': 1234,
//     'private': false,
//     'participant_count': 3,
//   }
//
// Message:
//
//   !!!javascript
//   {
//     'type': 'Message',
//     'message_id': 1234,
//     'notification_category': 'Assignment Graded'
//   }
//
// Submission:
//
// Returns an {api:Submissions:Submission Submission} with its Course and Assignment data.
//
// Conference:
//
//   !!!javascript
//   {
//     'type': 'Conference',
//     'web_conference_id': 1234
//   }
//
// Collaboration:
//
//   !!!javascript
//   {
//     'type': 'Collaboration',
//     'collaboration_id': 1234
//   }
//
// AssessmentRequest:
//
//   !!!javascript
//   {
//     'type': 'AssessmentRequest',
//     'assessment_request_id': 1234
//   }
// https://canvas.instructure.com/doc/api/users.html
//
// Query Parameters:
// # Query.OnlyActiveCourses (Optional) If true, will only return objects for courses the user is actively participating in
//
type ListActivityStreamActivityStream struct {
	Query struct {
		OnlyActiveCourses bool `json:"only_active_courses" url:"only_active_courses,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListActivityStreamActivityStream) GetMethod() string {
	return "GET"
}

func (t *ListActivityStreamActivityStream) GetURLPath() string {
	return ""
}

func (t *ListActivityStreamActivityStream) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListActivityStreamActivityStream) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListActivityStreamActivityStream) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListActivityStreamActivityStream) HasErrors() error {
	return nil
}

func (t *ListActivityStreamActivityStream) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
