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

// RestoreWorkflowStatesOfSISImportedItems This will restore the the workflow_state for all the items that changed
// their workflow_state during the import being restored.
// This will restore states for items imported with the following importers:
// accounts.csv terms.csv courses.csv sections.csv group_categories.csv
// groups.csv users.csv admins.csv
// This also restores states for other items that changed during the import.
// An example would be if an enrollment was deleted from a sis import and the
// group_membership was also deleted as a result of the enrollment deletion,
// both items would be restored when the sis batch is restored.
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # BatchMode (Optional) If set, will only restore items that were deleted from batch_mode.
// # UndeleteOnly (Optional) If set, will only restore items that were deleted. This will ignore any
//    items that were created or modified.
// # UnconcludeOnly (Optional) If set, will only restore enrollments that were concluded. This will
//    ignore any items that were created or deleted.
//
type RestoreWorkflowStatesOfSISImportedItems struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`

	Form struct {
		BatchMode      bool `json:"batch_mode"`      //  (Optional)
		UndeleteOnly   bool `json:"undelete_only"`   //  (Optional)
		UnconcludeOnly bool `json:"unconclude_only"` //  (Optional)
	} `json:"form"`
}

func (t *RestoreWorkflowStatesOfSISImportedItems) GetMethod() string {
	return "PUT"
}

func (t *RestoreWorkflowStatesOfSISImportedItems) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports/{id}/restore_states"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *RestoreWorkflowStatesOfSISImportedItems) GetQuery() (string, error) {
	return "", nil
}

func (t *RestoreWorkflowStatesOfSISImportedItems) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *RestoreWorkflowStatesOfSISImportedItems) HasErrors() error {
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

func (t *RestoreWorkflowStatesOfSISImportedItems) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
