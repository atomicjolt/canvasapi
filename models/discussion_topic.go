package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type DiscussionTopic struct {
	ID                      int64             `json:"id"`                        // The ID of this topic..Example: 1
	Title                   string            `json:"title"`                     // The topic title..Example: Topic 1
	Message                 string            `json:"message"`                   // The HTML content of the message body..Example: <p>content here</p>
	HtmlUrl                 string            `json:"html_url"`                  // The URL to the discussion topic in canvas..Example: https://<canvas>/courses/1/discussion_topics/2
	PostedAt                time.Time         `json:"posted_at"`                 // The datetime the topic was posted. If it is null it hasn't been posted yet. (see delayed_post_at).Example: 2037-07-21T13:29:31Z
	LastReplyAt             time.Time         `json:"last_reply_at"`             // The datetime for when the last reply was in the topic..Example: 2037-07-28T19:38:31Z
	RequireInitialPost      bool              `json:"require_initial_post"`      // If true then a user may not respond to other replies until that user has made an initial reply. Defaults to false..
	UserCanSeePosts         bool              `json:"user_can_see_posts"`        // Whether or not posts in this topic are visible to the user..Example: true
	DiscussionSubentryCount int64             `json:"discussion_subentry_count"` // The count of entries in the topic..Example: 0
	ReadState               string            `json:"read_state"`                // The read_state of the topic for the current user, 'read' or 'unread'..Example: read
	UnreadCount             int64             `json:"unread_count"`              // The count of unread entries of this topic for the current user..Example: 0
	Subscribed              bool              `json:"subscribed"`                // Whether or not the current user is subscribed to this topic..Example: true
	SubscriptionHold        string            `json:"subscription_hold"`         // (Optional) Why the user cannot subscribe to this topic. Only one reason will be returned even if multiple apply. Can be one of: 'initial_post_required': The user must post a reply first; 'not_in_group_set': The user is not in the group set for this graded group discussion; 'not_in_group': The user is not in this topic's group; 'topic_is_announcement': This topic is an announcement.Example: not_in_group_set
	AssignmentID            int64             `json:"assignment_id"`             // The unique identifier of the assignment if the topic is for grading, otherwise null..
	DelayedPostAt           time.Time         `json:"delayed_post_at"`           // The datetime to publish the topic (if not right away)..
	Published               bool              `json:"published"`                 // Whether this discussion topic is published (true) or draft state (false).Example: true
	LockAt                  time.Time         `json:"lock_at"`                   // The datetime to lock the topic (if ever)..
	Locked                  bool              `json:"locked"`                    // Whether or not the discussion is 'closed for comments'..
	Pinned                  bool              `json:"pinned"`                    // Whether or not the discussion has been 'pinned' by an instructor.
	LockedForUser           bool              `json:"locked_for_user"`           // Whether or not this is locked for the user..Example: true
	LockInfo                *LockInfo         `json:"lock_info"`                 // (Optional) Information for the user about the lock. Present when locked_for_user is true..
	LockExplanation         string            `json:"lock_explanation"`          // (Optional) An explanation of why this is locked for the user. Present when locked_for_user is true..Example: This discussion is locked until September 1 at 12:00am
	UserName                string            `json:"user_name"`                 // The username of the topic creator..Example: User Name
	TopicChildren           []int64           `json:"topic_children"`            // DEPRECATED An array of topic_ids for the group discussions the user is a part of..Example: 5, 7, 10
	GroupTopicChildren      string            `json:"group_topic_children"`      // An array of group discussions the user is a part of. Fields include: id, group_id.Example: {'id'=>5, 'group_id'=>1}, {'id'=>7, 'group_id'=>5}, {'id'=>10, 'group_id'=>4}
	RootTopicID             int64             `json:"root_topic_id"`             // If the topic is for grading and a group assignment this will point to the original topic in the course..
	PodcastUrl              string            `json:"podcast_url"`               // If the topic is a podcast topic this is the feed url for the current user..Example: /feeds/topics/1/enrollment_1XAcepje4u228rt4mi7Z1oFbRpn3RAkTzuXIGOPe.rss
	DiscussionType          string            `json:"discussion_type"`           // The type of discussion. Values are 'side_comment', for discussions that only allow one level of nested comments, and 'threaded' for fully threaded discussions..Example: side_comment
	GroupCategoryID         int64             `json:"group_category_id"`         // The unique identifier of the group category if the topic is a group discussion, otherwise null..
	Attachments             []*FileAttachment `json:"attachments"`               // Array of file attachments..
	Permissions             string            `json:"permissions"`               // The current user's permissions on this topic..Example: true
	AllowRating             bool              `json:"allow_rating"`              // Whether or not users can rate entries in this topic..Example: true
	OnlyGradersCanRate      bool              `json:"only_graders_can_rate"`     // Whether or not grade permissions are required to rate entries..Example: true
	SortByRating            bool              `json:"sort_by_rating"`            // Whether or not entries should be sorted by rating..Example: true
}

func (t *DiscussionTopic) HasError() error {
	var s []string
	s = []string{"read", "unread"}
	if !string_utils.Include(s, t.ReadState) {
		return fmt.Errorf("expected 'read_state' to be one of %v", s)
	}
	s = []string{"initial_post_required", "not_in_group_set", "not_in_group", "topic_is_announcement"}
	if !string_utils.Include(s, t.SubscriptionHold) {
		return fmt.Errorf("expected 'subscription_hold' to be one of %v", s)
	}
	s = []string{"side_comment", "threaded"}
	if !string_utils.Include(s, t.DiscussionType) {
		return fmt.Errorf("expected 'discussion_type' to be one of %v", s)
	}
	return nil
}
