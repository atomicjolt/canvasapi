package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// DeleteExternalFeedGroups Deletes the external feed.
// https://canvas.instructure.com/doc/api/announcement_external_feeds.html
//
// Path Parameters:
// # GroupID (Required) ID
// # ExternalFeedID (Required) ID
//
type DeleteExternalFeedGroups struct {
	Path struct {
		GroupID        string `json:"group_id" url:"group_id,omitempty"`                 //  (Required)
		ExternalFeedID string `json:"external_feed_id" url:"external_feed_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *DeleteExternalFeedGroups) GetMethod() string {
	return "DELETE"
}

func (t *DeleteExternalFeedGroups) GetURLPath() string {
	path := "groups/{group_id}/external_feeds/{external_feed_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{external_feed_id}", fmt.Sprintf("%v", t.Path.ExternalFeedID))
	return path
}

func (t *DeleteExternalFeedGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *DeleteExternalFeedGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteExternalFeedGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteExternalFeedGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.ExternalFeedID == "" {
		errs = append(errs, "'ExternalFeedID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteExternalFeedGroups) Do(c *canvasapi.Canvas) (*models.ExternalFeed, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ExternalFeed{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
