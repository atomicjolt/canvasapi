package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// CreateWebhookSubscription Creates a webook subscription for the specified event type and
// context.
// https://canvas.instructure.com/doc/api/webhooks_subscriptions.html
//
// Form Parameters:
// # Subscription (Required) The id of the context for the subscription.
// # Subscription (Required) The type of context for the subscription. Must be 'assignment',
//    'account', or 'course'.
// # Subscription (Required) Array of strings representing the event types for
//    the subscription.
// # Subscription (Required) Format to deliver the live events. Must be 'live-event' or 'caliper'.
// # Subscription (Required) An object with a single key: 'Url'. Example: { "Url": "sqs.example" }
// # Subscription (Required) Must be either 'sqs' or 'https'.
//
type CreateWebhookSubscription struct {
	Form struct {
		Subscription struct {
			ContextID         string `json:"context_id" url:"context_id,omitempty"`                 //  (Required)
			ContextType       string `json:"context_type" url:"context_type,omitempty"`             //  (Required)
			EventTypes        string `json:"event_types" url:"event_types,omitempty"`               //  (Required)
			Format            string `json:"format" url:"format,omitempty"`                         //  (Required)
			TransportMetadata string `json:"transport_metadata" url:"transport_metadata,omitempty"` //  (Required)
			TransportType     string `json:"transport_type" url:"transport_type,omitempty"`         //  (Required)
		} `json:"subscription" url:"subscription,omitempty"`
	} `json:"form"`
}

func (t *CreateWebhookSubscription) GetMethod() string {
	return "POST"
}

func (t *CreateWebhookSubscription) GetURLPath() string {
	return ""
}

func (t *CreateWebhookSubscription) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateWebhookSubscription) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateWebhookSubscription) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateWebhookSubscription) HasErrors() error {
	errs := []string{}
	if t.Form.Subscription.ContextID == "" {
		errs = append(errs, "'Subscription' is required")
	}
	if t.Form.Subscription.ContextType == "" {
		errs = append(errs, "'Subscription' is required")
	}
	if t.Form.Subscription.EventTypes == "" {
		errs = append(errs, "'Subscription' is required")
	}
	if t.Form.Subscription.Format == "" {
		errs = append(errs, "'Subscription' is required")
	}
	if t.Form.Subscription.TransportMetadata == "" {
		errs = append(errs, "'Subscription' is required")
	}
	if t.Form.Subscription.TransportType == "" {
		errs = append(errs, "'Subscription' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateWebhookSubscription) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
