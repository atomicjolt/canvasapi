package requests

import (
	"fmt"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// DeleteWebhookSubscription
// https://canvas.instructure.com/doc/api/webhooks_subscriptions.html
//
// Path Parameters:
// # ID (Required) ID
//
type DeleteWebhookSubscription struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`
}

func (t *DeleteWebhookSubscription) GetMethod() string {
	return "DELETE"
}

func (t *DeleteWebhookSubscription) GetURLPath() string {
	path := "/lti/subscriptions/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteWebhookSubscription) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteWebhookSubscription) GetBody() (string, error) {
	return "", nil
}

func (t *DeleteWebhookSubscription) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteWebhookSubscription) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
