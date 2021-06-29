package models

import (
	"time"
)

type Conversation struct {
	ID               int64                      `json:"id" url:"id,omitempty"`                               // the unique identifier for the conversation..Example: 2
	Subject          string                     `json:"subject" url:"subject,omitempty"`                     // the subject of the conversation..Example: 2
	WorkflowState    string                     `json:"workflow_state" url:"workflow_state,omitempty"`       // The current state of the conversation (read, unread or archived)..Example: unread
	LastMessage      string                     `json:"last_message" url:"last_message,omitempty"`           // A <=100 character preview from the most recent message..Example: sure thing, here's the file
	StartAt          time.Time                  `json:"start_at" url:"start_at,omitempty"`                   // the date and time at which the last message was sent..Example: 2011-09-02T12:00:00Z
	MessageCount     int64                      `json:"message_count" url:"message_count,omitempty"`         // the number of messages in the conversation..Example: 2
	Subscribed       bool                       `json:"subscribed" url:"subscribed,omitempty"`               // whether the current user is subscribed to the conversation..Example: true
	Private          bool                       `json:"private" url:"private,omitempty"`                     // whether the conversation is private..Example: true
	Starred          bool                       `json:"starred" url:"starred,omitempty"`                     // whether the conversation is starred..Example: true
	Properties       []string                   `json:"properties" url:"properties,omitempty"`               // Additional conversation flags (last_author, attachments, media_objects). Each listed property means the flag is set to true (i.e. the current user is the most recent author, there are attachments, or there are media objects).
	Audience         []string                   `json:"audience" url:"audience,omitempty"`                   // Array of user ids who are involved in the conversation, ordered by participation level, then alphabetical. Excludes current user, unless this is a monologue..
	AudienceContexts []string                   `json:"audience_contexts" url:"audience_contexts,omitempty"` // Most relevant shared contexts (courses and groups) between current user and other participants. If there is only one participant, it will also include that user's enrollment(s)/ membership type(s) in each course/group..
	AvatarUrl        string                     `json:"avatar_url" url:"avatar_url,omitempty"`               // URL to appropriate icon for this conversation (custom, individual or group avatar, depending on audience)..Example: https://canvas.instructure.com/images/messages/avatar-group-50.png
	Participants     []*ConversationParticipant `json:"participants" url:"participants,omitempty"`           // Array of users participating in the conversation. Includes current user..
	Visible          bool                       `json:"visible" url:"visible,omitempty"`                     // indicates whether the conversation is visible under the current scope and filter. This attribute is always true in the index API response, and is primarily useful in create/update responses so that you can know if the record should be displayed in the UI. The default scope is assumed, unless a scope or filter is passed to the create/update API call..Example: true
	ContextName      string                     `json:"context_name" url:"context_name,omitempty"`           // Name of the course or group in which the conversation is occurring..Example: Canvas 101
}

func (t *Conversation) HasErrors() error {
	return nil
}
