package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// GetRunningBatches Returns any currently running conversation batches for the current user.
// Conversation batches are created when a bulk private message is sent
// asynchronously (see the mode argument to the {api:ConversationsController#create create API action}).
// https://canvas.instructure.com/doc/api/conversations.html
//
type GetRunningBatches struct {
}

func (t *GetRunningBatches) GetMethod() string {
	return "GET"
}

func (t *GetRunningBatches) GetURLPath() string {
	return ""
}

func (t *GetRunningBatches) GetQuery() (string, error) {
	return "", nil
}

func (t *GetRunningBatches) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetRunningBatches) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetRunningBatches) HasErrors() error {
	return nil
}

func (t *GetRunningBatches) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
