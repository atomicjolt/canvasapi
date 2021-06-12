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

// ImportOutcomeGroupAccounts Creates a new subgroup of the outcome group with the same title and
// description as the source group, then creates links in that new subgroup to
// the same outcomes that are linked in the source group. Recurses on the
// subgroups of the source group, importing them each in turn into the new
// subgroup.
//
// Allows you to copy organizational structure, but does not create copies of
// the outcomes themselves, only new links.
//
// The source group must be either global, from the same context as this
// outcome group, or from an associated account. The source group cannot be
// the root outcome group of its context.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # SourceOutcomeGroupID (Required) The ID of the source outcome group.
// # Async (Optional) If true, perform action asynchronously.  In that case, this endpoint
//    will return a Progress object instead of an OutcomeGroup.
//    Use the {api:ProgressController#show progress endpoint}
//    to query the status of the operation.  The imported outcome group id
//    and url will be returned in the results of the Progress object
//    as "outcome_group_id" and "outcome_group_url"
//
type ImportOutcomeGroupAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
		ID        string `json:"id"`         //  (Required)
	} `json:"path"`

	Form struct {
		SourceOutcomeGroupID int64 `json:"source_outcome_group_id"` //  (Required)
		Async                bool  `json:"async"`                   //  (Optional)
	} `json:"form"`
}

func (t *ImportOutcomeGroupAccounts) GetMethod() string {
	return "POST"
}

func (t *ImportOutcomeGroupAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_groups/{id}/import"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ImportOutcomeGroupAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ImportOutcomeGroupAccounts) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ImportOutcomeGroupAccounts) HasErrors() error {
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

func (t *ImportOutcomeGroupAccounts) Do(c *canvasapi.Canvas) (*models.OutcomeGroup, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeGroup{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
