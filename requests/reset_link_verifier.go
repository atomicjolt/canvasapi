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

// ResetLinkVerifier Resets the link verifier. Any existing links to the file using
// the previous hard-coded "verifier" parameter will no longer
// automatically grant access.
//
// Must have manage files and become other users permissions
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
type ResetLinkVerifier struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ResetLinkVerifier) GetMethod() string {
	return "POST"
}

func (t *ResetLinkVerifier) GetURLPath() string {
	path := "files/{id}/reset_verifier"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ResetLinkVerifier) GetQuery() (string, error) {
	return "", nil
}

func (t *ResetLinkVerifier) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ResetLinkVerifier) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ResetLinkVerifier) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ResetLinkVerifier) Do(c *canvasapi.Canvas) (*models.File, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.File{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
