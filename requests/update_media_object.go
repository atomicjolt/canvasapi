package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateMediaObject
// https://canvas.instructure.com/doc/api/media_objects.html
//
// Path Parameters:
// # MediaObjectID (Required) ID
//
// Form Parameters:
// # UserEnteredTitle (Optional) The new title.
//
type UpdateMediaObject struct {
	Path struct {
		MediaObjectID string `json:"media_object_id"` //  (Required)
	} `json:"path"`

	Form struct {
		UserEnteredTitle string `json:"user_entered_title"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateMediaObject) GetMethod() string {
	return "PUT"
}

func (t *UpdateMediaObject) GetURLPath() string {
	path := "media_objects/{media_object_id}"
	path = strings.ReplaceAll(path, "{media_object_id}", fmt.Sprintf("%v", t.Path.MediaObjectID))
	return path
}

func (t *UpdateMediaObject) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateMediaObject) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateMediaObject) HasErrors() error {
	errs := []string{}
	if t.Path.MediaObjectID == "" {
		errs = append(errs, "'MediaObjectID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateMediaObject) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
