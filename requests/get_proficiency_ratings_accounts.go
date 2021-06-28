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

// GetProficiencyRatingsAccounts Get account-level proficiency ratings. If not defined for this account,
// it will return proficiency ratings for the nearest super-account with ratings defined.
// Will return 404 if none found.
//
//   Examples:
//     curl https://<canvas>/api/v1/accounts/<account_id>/outcome_proficiency \
//         -H 'Authorization: Bearer <token>'
// https://canvas.instructure.com/doc/api/proficiency_ratings.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetProficiencyRatingsAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetProficiencyRatingsAccounts) GetMethod() string {
	return "GET"
}

func (t *GetProficiencyRatingsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_proficiency"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetProficiencyRatingsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *GetProficiencyRatingsAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetProficiencyRatingsAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetProficiencyRatingsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetProficiencyRatingsAccounts) Do(c *canvasapi.Canvas) (*models.Proficiency, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Proficiency{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
