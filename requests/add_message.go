package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// AddMessage Add a message to an existing conversation. Response is similar to the
// GET/show action, except that only includes the
// latest message (i.e. what we just sent)
//
// An array of user ids. Defaults to all of the current conversation
// recipients. To explicitly send a message to no other recipients,
// this array should consist of the logged-in user id.
//
// An array of message ids from this conversation to send to recipients
// of the new message. Recipients who already had a copy of included
// messages will not be affected.
// https://canvas.instructure.com/doc/api/conversations.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Body (Required) The message to be sent.
// # Form.AttachmentIDs (Optional) An array of attachments ids. These must be files that have been previously
//    uploaded to the sender's "conversation attachments" folder.
// # Form.MediaCommentID (Optional) Media comment id of an audio of video file to be associated with this
//    message.
// # Form.MediaCommentType (Optional) . Must be one of audio, videoType of the associated media file.
// # Form.Recipients (Optional) no description
// # Form.IncludedMessages (Optional) no description
// # Form.UserNote (Optional) Will add a faculty journal entry for each recipient as long as the user
//    making the api call has permission, the recipient is a student and
//    faculty journals are enabled in the account.
//
type AddMessage struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Body             string   `json:"body" url:"body,omitempty"`                             //  (Required)
		AttachmentIDs    []string `json:"attachment_ids" url:"attachment_ids,omitempty"`         //  (Optional)
		MediaCommentID   string   `json:"media_comment_id" url:"media_comment_id,omitempty"`     //  (Optional)
		MediaCommentType string   `json:"media_comment_type" url:"media_comment_type,omitempty"` //  (Optional) . Must be one of audio, video
		Recipients       []string `json:"recipients" url:"recipients,omitempty"`                 //  (Optional)
		IncludedMessages []string `json:"included_messages" url:"included_messages,omitempty"`   //  (Optional)
		UserNote         bool     `json:"user_note" url:"user_note,omitempty"`                   //  (Optional)
	} `json:"form"`
}

func (t *AddMessage) GetMethod() string {
	return "POST"
}

func (t *AddMessage) GetURLPath() string {
	path := "conversations/{id}/add_message"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *AddMessage) GetQuery() (string, error) {
	return "", nil
}

func (t *AddMessage) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AddMessage) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AddMessage) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.Body == "" {
		errs = append(errs, "'Form.Body' is required")
	}
	if t.Form.MediaCommentType != "" && !string_utils.Include([]string{"audio", "video"}, t.Form.MediaCommentType) {
		errs = append(errs, "MediaCommentType must be one of audio, video")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddMessage) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
