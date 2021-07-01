package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListMediaTracksForMediaObject List the media tracks associated with a media object
// https://canvas.instructure.com/doc/api/media_objects.html
//
// Path Parameters:
// # Path.MediaObjectID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of content, webvtt_content, updated_at, created_atBy default, index returns id, locale, kind, media_object_id, and user_id for each of the
//    result MediaTracks. Use include[] to
//    add additional fields. For example include[]=content
//
type ListMediaTracksForMediaObject struct {
	Path struct {
		MediaObjectID string `json:"media_object_id" url:"media_object_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of content, webvtt_content, updated_at, created_at
	} `json:"query"`
}

func (t *ListMediaTracksForMediaObject) GetMethod() string {
	return "GET"
}

func (t *ListMediaTracksForMediaObject) GetURLPath() string {
	path := "media_objects/{media_object_id}/media_tracks"
	path = strings.ReplaceAll(path, "{media_object_id}", fmt.Sprintf("%v", t.Path.MediaObjectID))
	return path
}

func (t *ListMediaTracksForMediaObject) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListMediaTracksForMediaObject) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMediaTracksForMediaObject) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMediaTracksForMediaObject) HasErrors() error {
	errs := []string{}
	if t.Path.MediaObjectID == "" {
		errs = append(errs, "'Path.MediaObjectID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"content", "webvtt_content", "updated_at", "created_at"}, v) {
			errs = append(errs, "Include must be one of content, webvtt_content, updated_at, created_at")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMediaTracksForMediaObject) Do(c *canvasapi.Canvas) ([]*models.MediaTrack, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.MediaTrack{}
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
