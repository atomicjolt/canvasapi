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

// GetSISImportErrorListSISImports Returns the list of SIS import errors for an account or a SIS import. Import
// errors are only stored for 30 days.
//
// Example:
//   curl 'https://<canvas>/api/v1/accounts/<account_id>/sis_imports/<id>/sis_import_errors' \
//     -H "Authorization: Bearer <token>"
//
// Example:
//   curl 'https://<canvas>/api/v1/accounts/<account_id>/sis_import_errors' \
//     -H "Authorization: Bearer <token>"
// https://canvas.instructure.com/doc/api/sis_import_errors.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Failure (Optional) If set, only shows errors on a sis import that would cause a failure.
//
type GetSISImportErrorListSISImports struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Query struct {
		Failure bool `json:"failure" url:"failure,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetSISImportErrorListSISImports) GetMethod() string {
	return "GET"
}

func (t *GetSISImportErrorListSISImports) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports/{id}/errors"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSISImportErrorListSISImports) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSISImportErrorListSISImports) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSISImportErrorListSISImports) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSISImportErrorListSISImports) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSISImportErrorListSISImports) Do(c *canvasapi.Canvas) ([]*models.SISImportError, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.SISImportError{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
