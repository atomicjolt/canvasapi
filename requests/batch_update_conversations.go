package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// BatchUpdateConversations Perform a change on a set of conversations. Operates asynchronously; use the {api:ProgressController#show progress endpoint}
// to query the status of an operation.
// https://canvas.instructure.com/doc/api/conversations.html
//
// Form Parameters:
// # Form.ConversationIDs (Required) List of conversations to update. Limited to 500 conversations.
// # Form.Event (Required) . Must be one of mark_as_read, mark_as_unread, star, unstar, archive, destroyThe action to take on each conversation.
//
type BatchUpdateConversations struct {
	Form struct {
		ConversationIDs []string `json:"conversation_ids" url:"conversation_ids,omitempty"` //  (Required)
		Event           string   `json:"event" url:"event,omitempty"`                       //  (Required) . Must be one of mark_as_read, mark_as_unread, star, unstar, archive, destroy
	} `json:"form"`
}

func (t *BatchUpdateConversations) GetMethod() string {
	return "PUT"
}

func (t *BatchUpdateConversations) GetURLPath() string {
	return ""
}

func (t *BatchUpdateConversations) GetQuery() (string, error) {
	return "", nil
}

func (t *BatchUpdateConversations) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *BatchUpdateConversations) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *BatchUpdateConversations) HasErrors() error {
	errs := []string{}
	if t.Form.ConversationIDs == nil {
		errs = append(errs, "'Form.ConversationIDs' is required")
	}
	if t.Form.Event == "" {
		errs = append(errs, "'Form.Event' is required")
	}
	if t.Form.Event != "" && !string_utils.Include([]string{"mark_as_read", "mark_as_unread", "star", "unstar", "archive", "destroy"}, t.Form.Event) {
		errs = append(errs, "Event must be one of mark_as_read, mark_as_unread, star, unstar, archive, destroy")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *BatchUpdateConversations) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
