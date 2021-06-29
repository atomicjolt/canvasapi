package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// UpdateWebhookSubscription This endpoint uses the same parameters as the create endpoint
// https://canvas.instructure.com/doc/api/webhooks_subscriptions.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type UpdateWebhookSubscription struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *UpdateWebhookSubscription) GetMethod() string {
	return "PUT"
}

func (t *UpdateWebhookSubscription) GetURLPath() string {
	path := "/lti/subscriptions/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateWebhookSubscription) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateWebhookSubscription) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UpdateWebhookSubscription) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UpdateWebhookSubscription) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateWebhookSubscription) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
