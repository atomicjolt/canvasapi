package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// ShowSingleWebhookSubscription
// https://canvas.instructure.com/doc/api/webhooks_subscriptions.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type ShowSingleWebhookSubscription struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ShowSingleWebhookSubscription) GetMethod() string {
	return "GET"
}

func (t *ShowSingleWebhookSubscription) GetURLPath() string {
	path := "/lti/subscriptions/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowSingleWebhookSubscription) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowSingleWebhookSubscription) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowSingleWebhookSubscription) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowSingleWebhookSubscription) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowSingleWebhookSubscription) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
