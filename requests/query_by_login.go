package requests

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// QueryByLogin List authentication events for a given login.
// https://canvas.instructure.com/doc/api/authentications_log.html
//
// Path Parameters:
// # LoginID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want events.
//    Events are stored for one year.
// # EndTime (Optional) The end of the time range from which you want events.
//
type QueryByLogin struct {
	Path struct {
		LoginID string `json:"login_id" url:"login_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *QueryByLogin) GetMethod() string {
	return "GET"
}

func (t *QueryByLogin) GetURLPath() string {
	path := "audit/authentication/logins/{login_id}"
	path = strings.ReplaceAll(path, "{login_id}", fmt.Sprintf("%v", t.Path.LoginID))
	return path
}

func (t *QueryByLogin) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *QueryByLogin) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *QueryByLogin) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *QueryByLogin) HasErrors() error {
	errs := []string{}
	if t.Path.LoginID == "" {
		errs = append(errs, "'LoginID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QueryByLogin) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
