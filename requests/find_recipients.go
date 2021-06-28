package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// FindRecipients Deprecated, see the {api:SearchController#recipients Find recipients endpoint} in the Search API
// https://canvas.instructure.com/doc/api/conversations.html
//
type FindRecipients struct {
}

func (t *FindRecipients) GetMethod() string {
	return "GET"
}

func (t *FindRecipients) GetURLPath() string {
	return ""
}

func (t *FindRecipients) GetQuery() (string, error) {
	return "", nil
}

func (t *FindRecipients) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *FindRecipients) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *FindRecipients) HasErrors() error {
	return nil
}

func (t *FindRecipients) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
