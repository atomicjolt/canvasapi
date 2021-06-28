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
		Recipients        []string `json:"recipients" url:"recipients,omitempty"`                 //  (Required)
		Subject           string   `json:"subject" url:"subject,omitempty"`                       //  (Optional)
		Body              string   `json:"body" url:"body,omitempty"`                             //  (Required)
		ForceNew          bool     `json:"force_new" url:"force_new,omitempty"`                   //  (Optional)
		GroupConversation bool     `json:"group_conversation" url:"group_conversation,omitempty"` //  (Optional)
		AttachmentIDs     []string `json:"attachment_ids" url:"attachment_ids,omitempty"`         //  (Optional)
		MediaCommentID    string   `json:"media_comment_id" url:"media_comment_id,omitempty"`     //  (Optional)
		MediaCommentType  string   `json:"media_comment_type" url:"media_comment_type,omitempty"` //  (Optional) . Must be one of audio, video
		UserNote          bool     `json:"user_note" url:"user_note,omitempty"`                   //  (Optional)
		Mode              string   `json:"mode" url:"mode,omitempty"`                             //  (Optional) . Must be one of sync, async
		Scope             string   `json:"scope" url:"scope,omitempty"`                           //  (Optional) . Must be one of unread, starred, archived
		Filter            []string `json:"filter" url:"filter,omitempty"`                         //  (Optional)
		FilterMode        string   `json:"filter_mode" url:"filter_mode,omitempty"`               //  (Optional) . Must be one of and, or, default or
		ContextCode       string   `json:"context_code" url:"context_code,omitempty"`             //  (Optional)
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

func (t *CreateConversation) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateConversation) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateConversation) HasErrors() error {
	errs := []string{}
	if t.Form.Recipients == nil {
		errs = append(errs, "'Recipients' is required")
	}
	if t.Form.Body == "" {
		errs = append(errs, "'Body' is required")
	}
	if t.Form.MediaCommentType != "" && !string_utils.Include([]string{"audio", "video"}, t.Form.MediaCommentType) {
		errs = append(errs, "MediaCommentType must be one of audio, video")
	}
	if t.Form.Mode != "" && !string_utils.Include([]string{"sync", "async"}, t.Form.Mode) {
		errs = append(errs, "Mode must be one of sync, async")
	}
	if t.Form.Scope != "" && !string_utils.Include([]string{"unread", "starred", "archived"}, t.Form.Scope) {
		errs = append(errs, "Scope must be one of unread, starred, archived")
	}
	if t.Form.FilterMode != "" && !string_utils.Include([]string{"and", "or", "default or"}, t.Form.FilterMode) {
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
