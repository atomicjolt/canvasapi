package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// GetSISImportList Returns the list of SIS imports for an account
//
// Example:
//   curl https://<canvas>/api/v1/accounts/<account_id>/sis_imports \
//     -H 'Authorization: Bearer <token>'
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # CreatedSince (Optional) If set, only shows imports created after the specified date (use ISO8601 format)
// # WorkflowState (Optional) . Must be one of initializing, created, importing, cleanup_batch, imported, imported_with_messages, aborted, failed, failed_with_messages, restoring, partially_restored, restoredIf set, only returns imports that are in the given state.
//
type GetSISImportList struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		CreatedSince  time.Time `json:"created_since"`  //  (Optional)
		WorkflowState []string  `json:"workflow_state"` //  (Optional) . Must be one of initializing, created, importing, cleanup_batch, imported, imported_with_messages, aborted, failed, failed_with_messages, restoring, partially_restored, restored
	} `json:"query"`
}

func (t *GetSISImportList) GetMethod() string {
	return "GET"
}

func (t *GetSISImportList) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetSISImportList) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSISImportList) GetBody() (string, error) {
	return "", nil
}

func (t *GetSISImportList) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	for _, v := range t.Query.WorkflowState {
		if !string_utils.Include([]string{"initializing", "created", "importing", "cleanup_batch", "imported", "imported_with_messages", "aborted", "failed", "failed_with_messages", "restoring", "partially_restored", "restored"}, v) {
			errs = append(errs, "WorkflowState must be one of initializing, created, importing, cleanup_batch, imported, imported_with_messages, aborted, failed, failed_with_messages, restoring, partially_restored, restored")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSISImportList) Do(c *canvasapi.Canvas) ([]*models.SISImport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.SISImport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
