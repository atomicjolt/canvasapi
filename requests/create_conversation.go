package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateConversation Create a new conversation with one or more recipients. If there is already
// an existing private conversation with the given recipients, it will be
// reused.
// https://canvas.instructure.com/doc/api/conversations.html
//
// Form Parameters:
// # Recipients (Required) An array of recipient ids. These may be user ids or course/group ids
//    prefixed with "course_" or "group_" respectively, e.g.
//    recipients[]=1&recipients[]=2&recipients[]=course_3. If the course/group
//    has over 100 enrollments, 'bulk_message' and 'group_conversation' must be
//    set to true.
// # Subject (Optional) The subject of the conversation. This is ignored when reusing a
//    conversation. Maximum length is 255 characters.
// # Body (Required) The message to be sent
// # ForceNew (Optional) Forces a new message to be created, even if there is an existing private conversation.
// # GroupConversation (Optional) Defaults to false.  When false, individual private conversations will be
//    created with each recipient. If true, this will be a group conversation
//    (i.e. all recipients may see all messages and replies). Must be set true if
//    the number of recipients is over the set maximum (default is 100).
// # AttachmentIDs (Optional) An array of attachments ids. These must be files that have been previously
//    uploaded to the sender's "conversation attachments" folder.
// # MediaCommentID (Optional) Media comment id of an audio of video file to be associated with this
//    message.
// # MediaCommentType (Optional) . Must be one of audio, videoType of the associated media file
// # UserNote (Optional) Will add a faculty journal entry for each recipient as long as the user
//    making the api call has permission, the recipient is a student and
//    faculty journals are enabled in the account.
// # Mode (Optional) . Must be one of sync, asyncDetermines whether the messages will be created/sent synchronously or
//    asynchronously. Defaults to sync, and this option is ignored if this is a
//    group conversation or there is just one recipient (i.e. it must be a bulk
//    private message). When sent async, the response will be an empty array
//    (batch status can be queried via the {api:ConversationsController#batches batches API})
// # Scope (Optional) . Must be one of unread, starred, archivedUsed when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # Filter (Optional) Used when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # FilterMode (Optional) . Must be one of and, or, default orUsed when generating "visible" in the API response. See the explanation
//    under the {api:ConversationsController#index index API action}
// # ContextCode (Optional) The course or group that is the context for this conversation. Same format
//    as courses or groups in the recipients argument.
//
type CreateConversation struct {
	Form struct {
		Recipients        []string `json:"recipients"`         //  (Required)
		Subject           string   `json:"subject"`            //  (Optional)
		Body              string   `json:"body"`               //  (Required)
		ForceNew          bool     `json:"force_new"`          //  (Optional)
		GroupConversation bool     `json:"group_conversation"` //  (Optional)
		AttachmentIDs     []string `json:"attachment_ids"`     //  (Optional)
		MediaCommentID    string   `json:"media_comment_id"`   //  (Optional)
		MediaCommentType  string   `json:"media_comment_type"` //  (Optional) . Must be one of audio, video
		UserNote          bool     `json:"user_note"`          //  (Optional)
		Mode              string   `json:"mode"`               //  (Optional) . Must be one of sync, async
		Scope             string   `json:"scope"`              //  (Optional) . Must be one of unread, starred, archived
		Filter            []string `json:"filter"`             //  (Optional)
		FilterMode        string   `json:"filter_mode"`        //  (Optional) . Must be one of and, or, default or
		ContextCode       string   `json:"context_code"`       //  (Optional)
	} `json:"form"`
}

func (t *CreateConversation) GetMethod() string {
	return "POST"
}

func (t *CreateConversation) GetURLPath() string {
	return ""
}

func (t *CreateConversation) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateConversation) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateConversation) HasErrors() error {
	errs := []string{}
	if t.Form.Recipients == nil {
		errs = append(errs, "'Recipients' is required")
	}
	if t.Form.Body == "" {
		errs = append(errs, "'Body' is required")
	}
	if !string_utils.Include([]string{"audio", "video"}, t.Form.MediaCommentType) {
		errs = append(errs, "MediaCommentType must be one of audio, video")
	}
	if !string_utils.Include([]string{"sync", "async"}, t.Form.Mode) {
		errs = append(errs, "Mode must be one of sync, async")
	}
	if !string_utils.Include([]string{"unread", "starred", "archived"}, t.Form.Scope) {
		errs = append(errs, "Scope must be one of unread, starred, archived")
	}
	if !string_utils.Include([]string{"and", "or", "default or"}, t.Form.FilterMode) {
		errs = append(errs, "FilterMode must be one of and, or, default or")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateConversation) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
