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

// UnlinkOutcomeAccounts Unlinking an outcome only deletes the outcome itself if this was the last
// link to the outcome in any group in any context. Aligned outcomes cannot be
// deleted; as such, if this is the last link to an aligned outcome, the
// unlinking will fail.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
// # OutcomeID (Required) ID
//
type UnlinkOutcomeAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
		OutcomeID string `json:"outcome_id" url:"outcome_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *UnlinkOutcomeAccounts) GetMethod() string {
	return "DELETE"
}

func (t *UnlinkOutcomeAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_groups/{id}/outcomes/{outcome_id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{outcome_id}", fmt.Sprintf("%v", t.Path.OutcomeID))
	return path
}

func (t *UnlinkOutcomeAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *UnlinkOutcomeAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UnlinkOutcomeAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UnlinkOutcomeAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.OutcomeID == "" {
		errs = append(errs, "'OutcomeID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UnlinkOutcomeAccounts) Do(c *canvasapi.Canvas) (*models.OutcomeLink, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeLink{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
