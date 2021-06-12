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

// GetSubAccountsOfAccount List accounts that are sub-accounts of the given account.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Query Parameters:
// # Recursive (Optional) If true, the entire account tree underneath
//    this account will be returned (though still paginated). If false, only
//    direct sub-accounts of this account will be returned. Defaults to false.
//
type GetSubAccountsOfAccount struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Recursive bool `json:"recursive"` //  (Optional)
	} `json:"query"`
}

func (t *GetSubAccountsOfAccount) GetMethod() string {
	return "GET"
}

func (t *GetSubAccountsOfAccount) GetURLPath() string {
	path := "accounts/{account_id}/sub_accounts"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetSubAccountsOfAccount) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSubAccountsOfAccount) GetBody() (string, error) {
	return "", nil
}

func (t *GetSubAccountsOfAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSubAccountsOfAccount) Do(c *canvasapi.Canvas) ([]*models.Account, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Account{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
