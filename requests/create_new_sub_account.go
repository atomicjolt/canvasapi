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
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Account struct {
			Name                       string `json:"name"`                           //  (Required)
			SISAccountID               string `json:"sis_account_id"`                 //  (Optional)
			DefaultStorageQuotaMb      int64  `json:"default_storage_quota_mb"`       //  (Optional)
			DefaultUserStorageQuotaMb  int64  `json:"default_user_storage_quota_mb"`  //  (Optional)
			DefaultGroupStorageQuotaMb int64  `json:"default_group_storage_quota_mb"` //  (Optional)
		} `json:"account"`
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

func (t *CreateNewSubAccount) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
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
