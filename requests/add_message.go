package requests

import (
	"fmt"
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
// # ID (Required) ID
//
// Form Parameters:
// # Body (Required) The message to be sent.
// # AttachmentIDs (Optional) An array of attachments ids. These must be files that have been previously
//    uploaded to the sender's "conversation attachments" folder.
// # MediaCommentID (Optional) Media comment id of an audio of video file to be associated with this
//    message.
// # MediaCommentType (Optional) . Must be one of audio, videoType of the associated media file.
// # Recipients (Optional) no description
// # IncludedMessages (Optional) no description
// # UserNote (Optional) Will add a faculty journal entry for each recipient as long as the user
//    making the api call has permission, the recipient is a student and
//    faculty journals are enabled in the account.
//
type AddMessage struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Body             string   `json:"body"`               //  (Required)
		AttachmentIDs    []string `json:"attachment_ids"`     //  (Optional)
		MediaCommentID   string   `json:"media_comment_id"`   //  (Optional)
		MediaCommentType string   `json:"media_comment_type"` //  (Optional) . Must be one of audio, video
		Recipients       []string `json:"recipients"`         //  (Optional)
		IncludedMessages []string `json:"included_messages"`  //  (Optional)
		UserNote         bool     `json:"user_note"`          //  (Optional)
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

func (t *AddMessage) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *AddMessage) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Body == "" {
		errs = append(errs, "'Body' is required")
	}
	if !string_utils.Include([]string{"audio", "video"}, t.Form.MediaCommentType) {
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
