package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListRecentHistoryForUser Return a paginated list of the user's recent history. History entries are returned in descending order,
// newest to oldest. You may list history entries for yourself (use +self+ as the user_id), for a student you observe,
// or for a user you manage as an administrator. Note that the +per_page+ pagination argument is not supported
// and the number of history entries returned per page will vary.
// https://canvas.instructure.com/doc/api/history.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ListRecentHistoryForUser struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListRecentHistoryForUser) GetMethod() string {
	return "GET"
}

func (t *ListRecentHistoryForUser) GetURLPath() string {
	path := "users/{user_id}/history"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListRecentHistoryForUser) GetQuery() (string, error) {
	return "", nil
}

func (t *ListRecentHistoryForUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListRecentHistoryForUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListRecentHistoryForUser) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListRecentHistoryForUser) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.HistoryEntry, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.HistoryEntry{}
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
