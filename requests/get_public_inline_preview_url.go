package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetPublicInlinePreviewUrl Determine the URL that should be used for inline preview of the file.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # SubmissionID (Optional) The id of the submission the file is associated with.  Provide this argument to gain access to a file
//    that has been submitted to an assignment (Canvas will verify that the file belongs to the submission
//    and the calling user has rights to view the submission).
//
type GetPublicInlinePreviewUrl struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SubmissionID int64 `json:"submission_id" url:"submission_id,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetPublicInlinePreviewUrl) GetMethod() string {
	return "GET"
}

func (t *GetPublicInlinePreviewUrl) GetURLPath() string {
	path := "files/{id}/public_url"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetPublicInlinePreviewUrl) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetPublicInlinePreviewUrl) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetPublicInlinePreviewUrl) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetPublicInlinePreviewUrl) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetPublicInlinePreviewUrl) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
