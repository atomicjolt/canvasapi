package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListFeaturesAccounts A paginated list of all features that apply to a given Account, Course, or User.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListFeaturesAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *ListFeaturesAccounts) GetMethod() string {
	return "GET"
}

func (t *ListFeaturesAccounts) GetURLPath() string {
	path := "accounts/{account_id}/features"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListFeaturesAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListFeaturesAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *ListFeaturesAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListFeaturesAccounts) Do(c *canvasapi.Canvas) ([]*models.Feature, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Feature{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
