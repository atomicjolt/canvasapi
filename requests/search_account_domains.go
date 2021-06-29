package requests

import (
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SearchAccountDomains Returns a list of up to 5 matching account domains
//
// Partial match on name / domain are supported
// https://canvas.instructure.com/doc/api/account_domain_lookups.html
//
// Query Parameters:
// # Query.Name (Optional) campus name
// # Query.Domain (Optional) no description
// # Query.Latitude (Optional) no description
// # Query.Longitude (Optional) no description
//
type SearchAccountDomains struct {
	Query struct {
		Name      string  `json:"name" url:"name,omitempty"`           //  (Optional)
		Domain    string  `json:"domain" url:"domain,omitempty"`       //  (Optional)
		Latitude  float64 `json:"latitude" url:"latitude,omitempty"`   //  (Optional)
		Longitude float64 `json:"longitude" url:"longitude,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *SearchAccountDomains) GetMethod() string {
	return "GET"
}

func (t *SearchAccountDomains) GetURLPath() string {
	return ""
}

func (t *SearchAccountDomains) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *SearchAccountDomains) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *SearchAccountDomains) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *SearchAccountDomains) HasErrors() error {
	return nil
}

func (t *SearchAccountDomains) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
