package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// UpdateMediaTracks Replace the media tracks associated with a media object with
// the array of tracks provided in the body.
// Update will
// delete any existing tracks not listed,
// leave untouched any tracks with no content field,
// and update or create tracks with a content field.
// https://canvas.instructure.com/doc/api/media_objects.html
//
// Path Parameters:
// # MediaObjectID (Required) ID
//
// Form Parameters:
// # Include (Optional) Retuns a listing of the resulting set of MediaTracks.
//    Like List Media Objects, use the include[] parameter to
//    add additional fields.
//
type UpdateMediaTracks struct {
	Path struct {
		MediaObjectID string `json:"media_object_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Include []string `json:"include"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateMediaTracks) GetMethod() string {
	return "PUT"
}

func (t *UpdateMediaTracks) GetURLPath() string {
	path := "media_objects/{media_object_id}/media_tracks"
	path = strings.ReplaceAll(path, "{media_object_id}", fmt.Sprintf("%v", t.Path.MediaObjectID))
	return path
}

func (t *UpdateMediaTracks) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMediaTracks) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateMediaTracks) HasErrors() error {
	errs := []string{}
	if t.Path.MediaObjectID == "" {
		errs = append(errs, "'MediaObjectID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMediaTracks) Do(c *canvasapi.Canvas) ([]*models.MediaTrack, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.MediaTrack{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
