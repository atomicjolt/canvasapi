package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListAccounts A paginated list of accounts that the current user can view or manage.
// Typically, students and even teachers will get an empty list in response,
// only account admins can view the accounts that they are in.
// https://canvas.instructure.com/doc/api/accounts.html
//
// Query Parameters:
// # Include (Optional) . Must be one of lti_guid, registration_settings, servicesArray of additional information to include.
//
//    "lti_guid":: the 'tool_consumer_instance_guid' that will be sent for this account on LTI launches
//    "registration_settings":: returns info about the privacy policy and terms of use
//    "services":: returns services and whether they are enabled (requires account management permissions)
//
type ListAccounts struct {
	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of lti_guid, registration_settings, services
	} `json:"query"`
}

func (t *ListAccounts) GetMethod() string {
	return "GET"
}

func (t *ListAccounts) GetURLPath() string {
	return ""
}

func (t *ListAccounts) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *ListAccounts) HasErrors() error {
	errs := []string{}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"lti_guid", "registration_settings", "services"}, v) {
			errs = append(errs, "Include must be one of lti_guid, registration_settings, services")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListAccounts) Do(c *canvasapi.Canvas) ([]*models.Account, error) {
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
