package requests

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// QueryByUser List authentication events for a given user.
// https://canvas.instructure.com/doc/api/authentications_log.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want events.
//    Events are stored for one year.
// # EndTime (Optional) The end of the time range from which you want events.
//
type QueryByUser struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *QueryByUser) GetMethod() string {
	return "GET"
}

func (t *QueryByUser) GetURLPath() string {
	path := "audit/authentication/users/{user_id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *QueryByUser) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *QueryByUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *QueryByUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *QueryByUser) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QueryByUser) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
