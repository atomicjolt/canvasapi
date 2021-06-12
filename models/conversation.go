package models

import (
	"time"
)

type Conversation struct {
	ID               int64                      `json:"id"`                // the unique identifier for the conversation..Example: 2
	Subject          string                     `json:"subject"`           // the subject of the conversation..Example: 2
	WorkflowState    string                     `json:"workflow_state"`    // The current state of the conversation (read, unread or archived)..Example: unread
	LastMessage      string                     `json:"last_message"`      // A <=100 character preview from the most recent message..Example: sure thing, here's the file
	StartAt          time.Time                  `json:"start_at"`          // the date and time at which the last message was sent..Example: 2011-09-02T12:00:00Z
	MessageCount     int64                      `json:"message_count"`     // the number of messages in the conversation..Example: 2
	Subscribed       bool                       `json:"subscribed"`        // whether the current user is subscribed to the conversation..Example: true
	Private          bool                       `json:"private"`           // whether the conversation is private..Example: true
	Starred          bool                       `json:"starred"`           // whether the conversation is starred..Example: true
	Properties       []string                   `json:"properties"`        // Additional conversation flags (last_author, attachments, media_objects). Each listed property means the flag is set to true (i.e. the current user is the most recent author, there are attachments, or there are media objects).
	Audience         []int64                    `json:"audience"`          // Array of user ids who are involved in the conversation, ordered by participation level, then alphabetical. Excludes current user, unless this is a monologue..
	AudienceContexts []string                   `json:"audience_contexts"` // Most relevant shared contexts (courses and groups) between current user and other participants. If there is only one participant, it will also include that user's enrollment(s)/ membership type(s) in each course/group..
	AvatarUrl        string                     `json:"avatar_url"`        // URL to appropriate icon for this conversation (custom, individual or group avatar, depending on audience)..Example: https://canvas.instructure.com/images/messages/avatar-group-50.png
	Participants     []*ConversationParticipant `json:"participants"`      // Array of users participating in the conversation. Includes current user..
	Visible          bool                       `json:"visible"`           // indicates whether the conversation is visible under the current scope and filter. This attribute is always true in the index API response, and is primarily useful in create/update responses so that you can know if the record should be displayed in the UI. The default scope is assumed, unless a scope or filter is passed to the create/update API call..Example: true
	ContextName      string                     `json:"context_name"`      // Name of the course or group in which the conversation is occurring..Example: Canvas 101
}

func (t *Conversation) HasError() error {
	return nil
}
