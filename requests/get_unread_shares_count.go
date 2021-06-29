package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetUnreadSharesCount Return the number of content shares a user has received that have not yet been read. Use +self+ as the user_id
// to retrieve your own content shares. Only linked observers and administrators may view other users'
// content shares.
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type GetUnreadSharesCount struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetUnreadSharesCount) GetMethod() string {
	return "GET"
}

func (t *GetUnreadSharesCount) GetURLPath() string {
	path := "users/{user_id}/content_shares/unread_count"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GetUnreadSharesCount) GetQuery() (string, error) {
	return "", nil
}

func (t *GetUnreadSharesCount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetUnreadSharesCount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetUnreadSharesCount) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUnreadSharesCount) Do(c *canvasapi.Canvas) (*canvasapi.UnreadCount, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := canvasapi.UnreadCount{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
