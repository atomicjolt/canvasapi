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

// CreateNewSubAccount Add a new sub-account to a given account.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # Account (Required) The name of the new sub-account.
// # Account (Optional) The account's identifier in the Student Information System.
// # Account (Optional) The default course storage quota to be used, if not otherwise specified.
// # Account (Optional) The default user storage quota to be used, if not otherwise specified.
// # Account (Optional) The default group storage quota to be used, if not otherwise specified.
//
type CreateNewSubAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Account struct {
			Name                       string `json:"name" url:"name,omitempty"`                                                     //  (Required)
			SISAccountID               string `json:"sis_account_id" url:"sis_account_id,omitempty"`                                 //  (Optional)
			DefaultStorageQuotaMb      int64  `json:"default_storage_quota_mb" url:"default_storage_quota_mb,omitempty"`             //  (Optional)
			DefaultUserStorageQuotaMb  int64  `json:"default_user_storage_quota_mb" url:"default_user_storage_quota_mb,omitempty"`   //  (Optional)
			DefaultGroupStorageQuotaMb int64  `json:"default_group_storage_quota_mb" url:"default_group_storage_quota_mb,omitempty"` //  (Optional)
		} `json:"account" url:"account,omitempty"`
	} `json:"form"`
}

func (t *CreateNewSubAccount) GetMethod() string {
	return "POST"
}

func (t *CreateNewSubAccount) GetURLPath() string {
	path := "accounts/{account_id}/sub_accounts"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateNewSubAccount) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateNewSubAccount) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateNewSubAccount) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateNewSubAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.Account.Name == "" {
		errs = append(errs, "'Account' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateNewSubAccount) Do(c *canvasapi.Canvas) (*models.Account, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Account{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
