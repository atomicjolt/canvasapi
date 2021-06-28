package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetCustomColor Returns the custom colors that have been saved for a user for a given context.
//
// The asset_string parameter should be in the format 'context_id', for example
// 'course_42'.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
// # AssetString (Required) ID
//
type GetCustomColor struct {
	Path struct {
		ID          string `json:"id" url:"id,omitempty"`                     //  (Required)
		AssetString string `json:"asset_string" url:"asset_string,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetCustomColor) GetMethod() string {
	return "GET"
}

func (t *GetCustomColor) GetURLPath() string {
	path := "users/{id}/colors/{asset_string}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{asset_string}", fmt.Sprintf("%v", t.Path.AssetString))
	return path
}

func (t *GetCustomColor) GetQuery() (string, error) {
	return "", nil
}

func (t *GetCustomColor) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetCustomColor) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetCustomColor) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.AssetString == "" {
		errs = append(errs, "'AssetString' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetCustomColor) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
