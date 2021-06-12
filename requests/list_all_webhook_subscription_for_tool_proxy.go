package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ListAllWebhookSubscriptionForToolProxy This endpoint returns a paginated list with a default limit of 100 items per result set.
// You can retrieve the next result set by setting a 'StartKey' header in your next request
// with the value of the 'EndKey' header in the response.
//
// Example use of a 'StartKey' header object:
//   { "Id":"71d6dfba-0547-477d-b41d-db8cb528c6d1","DeveloperKey":"10000000000001" }
// https://canvas.instructure.com/doc/api/webhooks_subscriptions.html
//
type ListAllWebhookSubscriptionForToolProxy struct {
}

func (t *ListAllWebhookSubscriptionForToolProxy) GetMethod() string {
	return "GET"
}

func (t *ListAllWebhookSubscriptionForToolProxy) GetURLPath() string {
	return ""
}

func (t *ListAllWebhookSubscriptionForToolProxy) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAllWebhookSubscriptionForToolProxy) GetBody() (string, error) {
	return "", nil
}

func (t *ListAllWebhookSubscriptionForToolProxy) HasErrors() error {
	return nil
}

func (t *ListAllWebhookSubscriptionForToolProxy) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
