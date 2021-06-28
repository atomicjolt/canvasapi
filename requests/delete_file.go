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
)

// DeleteFile Remove the specified file. Unlike most other DELETE endpoints, using this
// endpoint will result in comprehensive, irretrievable destruction of the file.
// It should be used with the `replace` parameter set to true in cases where the
// file preview also needs to be destroyed (such as to remove files that violate
// privacy laws).
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # Replace (Optional) This action is irreversible.
//    If replace is set to true the file contents will be replaced with a
//    generic "file has been removed" file. This also destroys any previews
//    that have been generated for the file.
//    Must have manage files and become other users permissions
//
type DeleteFile struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Replace bool `json:"replace" url:"replace,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *DeleteFile) GetMethod() string {
	return "DELETE"
}

func (t *DeleteFile) GetURLPath() string {
	path := "files/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteFile) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *DeleteFile) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteFile) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteFile) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteFile) Do(c *canvasapi.Canvas) (*models.File, error) {
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
