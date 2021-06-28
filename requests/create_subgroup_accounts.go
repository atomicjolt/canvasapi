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

// CreateSubgroupAccounts Creates a new empty subgroup under the outcome group with the given title
// and description.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # AccountID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # Title (Required) The title of the new outcome group.
// # Description (Optional) The description of the new outcome group.
// # VendorGuid (Optional) A custom GUID for the learning standard
//
type CreateSubgroupAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Form struct {
		Title       string `json:"title" url:"title,omitempty"`             //  (Required)
		Description string `json:"description" url:"description,omitempty"` //  (Optional)
		VendorGuid  string `json:"vendor_guid" url:"vendor_guid,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *CreateSubgroupAccounts) GetMethod() string {
	return "POST"
}

func (t *CreateSubgroupAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_groups/{id}/subgroups"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *CreateSubgroupAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateSubgroupAccounts) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateSubgroupAccounts) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateSubgroupAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.Title == "" {
		errs = append(errs, "'Title' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateSubgroupAccounts) Do(c *canvasapi.Canvas) (*models.OutcomeGroup, error) {
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
