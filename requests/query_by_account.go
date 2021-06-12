package requests

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// QueryByAccount List authentication events for a given account.
// https://canvas.instructure.com/doc/api/authentications_log.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want events.
//    Events are stored for one year.
// # EndTime (Optional) The end of the time range from which you want events.
//
type QueryByAccount struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time"` //  (Optional)
		EndTime   time.Time `json:"end_time"`   //  (Optional)
	} `json:"query"`
}

func (t *QueryByAccount) GetMethod() string {
	return "GET"
}

func (t *QueryByAccount) GetURLPath() string {
	path := "audit/authentication/accounts/{account_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *QueryByAccount) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *QueryByAccount) GetBody() (string, error) {
	return "", nil
}

func (t *QueryByAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *QueryByAccount) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
