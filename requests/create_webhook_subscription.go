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
// # Form.Subscription.ContextID (Required) The id of the context for the subscription.
// # Form.Subscription.ContextType (Required) The type of context for the subscription. Must be 'assignment',
//    'account', or 'course'.
// # Form.Subscription.EventTypes (Required) Array of strings representing the event types for
//    the subscription.
// # Form.Subscription.Format (Required) Format to deliver the live events. Must be 'live-event' or 'caliper'.
// # Form.Subscription.TransportMetadata (Required) An object with a single key: 'Url'. Example: { "Url": "sqs.example" }
// # Form.Subscription.TransportType (Required) Must be either 'sqs' or 'https'.
//
type CreateWebhookSubscription struct {
	Form struct {
		Subscription struct {
			ContextID         string                   `json:"context_id" url:"context_id,omitempty"`                 //  (Required)
			ContextType       string                   `json:"context_type" url:"context_type,omitempty"`             //  (Required)
			EventTypes        string                   `json:"event_types" url:"event_types,omitempty"`               //  (Required)
			Format            string                   `json:"format" url:"format,omitempty"`                         //  (Required)
			TransportMetadata map[string](interface{}) `json:"transport_metadata" url:"transport_metadata,omitempty"` //  (Required)
			TransportType     string                   `json:"transport_type" url:"transport_type,omitempty"`         //  (Required)
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
		errs = append(errs, "'Form.Subscription.ContextID' is required")
	}
	if t.Form.Subscription.ContextType == "" {
		errs = append(errs, "'Form.Subscription.ContextType' is required")
	}
	if t.Form.Subscription.EventTypes == "" {
		errs = append(errs, "'Form.Subscription.EventTypes' is required")
	}
	if t.Form.Subscription.Format == "" {
		errs = append(errs, "'Form.Subscription.Format' is required")
	}
	if t.Form.Subscription.TransportMetadata == nil {
		errs = append(errs, "'Form.Subscription.TransportMetadata' is required")
	}
	if t.Form.Subscription.TransportType == "" {
		errs = append(errs, "'Form.Subscription.TransportType' is required")
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
