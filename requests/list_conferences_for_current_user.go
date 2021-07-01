package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListConferencesForCurrentUser Retrieve the paginated list of conferences for all courses and groups
// the current user belongs to
//
// This API returns a JSON object containing the list of conferences.
// The key for the list of conferences is "conferences".
// https://canvas.instructure.com/doc/api/conferences.html
//
// Query Parameters:
// # Query.State (Optional) If set to "live", returns only conferences that are live (i.e., have
//    started and not finished yet). If omitted, returns all conferences for
//    this user's groups and courses.
//
type ListConferencesForCurrentUser struct {
	Query struct {
		State string `json:"state" url:"state,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListConferencesForCurrentUser) GetMethod() string {
	return "GET"
}

func (t *ListConferencesForCurrentUser) GetURLPath() string {
	return ""
}

func (t *ListConferencesForCurrentUser) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListConferencesForCurrentUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListConferencesForCurrentUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListConferencesForCurrentUser) HasErrors() error {
	return nil
}

func (t *ListConferencesForCurrentUser) Do(c *canvasapi.Canvas) ([]*models.Conference, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Conference{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
