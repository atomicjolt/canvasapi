package requests

import (
	"fmt"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// SearchAccountDomains Returns a list of up to 5 matching account domains
//
// Partial match on name / domain are supported
// https://canvas.instructure.com/doc/api/account_domain_lookups.html
//
// Query Parameters:
// # Name (Optional) campus name
// # Domain (Optional) no description
// # Latitude (Optional) no description
// # Longitude (Optional) no description
//
type SearchAccountDomains struct {
	Query struct {
		Name      string  `json:"name"`      //  (Optional)
		Domain    string  `json:"domain"`    //  (Optional)
		Latitude  float64 `json:"latitude"`  //  (Optional)
		Longitude float64 `json:"longitude"` //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *SearchAccountDomains) GetBody() (string, error) {
	return "", nil
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
