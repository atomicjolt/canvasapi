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

// ListExternalFeedsGroups Returns the paginated list of External Feeds this course or group.
// https://canvas.instructure.com/doc/api/announcement_external_feeds.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ListExternalFeedsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListExternalFeedsGroups) GetMethod() string {
	return "GET"
}

func (t *ListExternalFeedsGroups) GetURLPath() string {
	path := "groups/{group_id}/external_feeds"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListExternalFeedsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListExternalFeedsGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListExternalFeedsGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListExternalFeedsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListExternalFeedsGroups) Do(c *canvasapi.Canvas) ([]*models.ExternalFeed, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.ExternalFeed{}
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
