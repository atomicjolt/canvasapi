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

// CreateUpdateProficiencyRatingsAccounts Create or update account-level proficiency ratings. These ratings will apply to all
// sub-accounts, unless they have their own account-level proficiency ratings defined.
// https://canvas.instructure.com/doc/api/proficiency_ratings.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # Ratings (Optional) The description of the rating level.
// # Ratings (Optional) The non-negative number of points of the rating level. Points across ratings should be strictly decreasing in value.
// # Ratings (Optional) Indicates the rating level where mastery is first achieved. Only one rating in a proficiency should be marked for mastery.
// # Ratings (Optional) The color associated with the rating level. Should be a hex color code like '00FFFF'.
//
type CreateUpdateProficiencyRatingsAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Ratings struct {
			Description []string `json:"description"` //  (Optional)
			Points      []int64  `json:"points"`      //  (Optional)
			Mastery     []int64  `json:"mastery"`     //  (Optional)
			Color       []int64  `json:"color"`       //  (Optional)
		} `json:"ratings"`
	} `json:"form"`
}

func (t *CreateUpdateProficiencyRatingsAccounts) GetMethod() string {
	return "POST"
}

func (t *CreateUpdateProficiencyRatingsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/outcome_proficiency"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateUpdateProficiencyRatingsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateUpdateProficiencyRatingsAccounts) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateUpdateProficiencyRatingsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateUpdateProficiencyRatingsAccounts) Do(c *canvasapi.Canvas) (*models.Proficiency, error) {
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
