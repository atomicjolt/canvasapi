package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListUserPageViews Return a paginated list of the user's page view history in json format,
// similar to the available CSV download. Page views are returned in
// descending order, newest to oldest.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Query Parameters:
// # StartTime (Optional) The beginning of the time range from which you want page views.
// # EndTime (Optional) The end of the time range from which you want page views.
//
type ListUserPageViews struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *ListUserPageViews) GetMethod() string {
	return "GET"
}

func (t *ListUserPageViews) GetURLPath() string {
	path := "users/{user_id}/page_views"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListUserPageViews) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListUserPageViews) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListUserPageViews) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListUserPageViews) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUserPageViews) Do(c *canvasapi.Canvas) ([]*models.PageView, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.PageView{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
