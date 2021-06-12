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

// DeprecatedSelfRegisterUser Self register and return a new user and pseudonym for an account.
//
// If self-registration is enabled on the account, you can use this
// endpoint to self register new users.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # User (Required) The full name of the user. This name will be used by teacher for grading.
// # User (Optional) User's name as it will be displayed in discussions, messages, and comments.
// # User (Optional) User's name as used to sort alphabetically in lists.
// # User (Optional) The time zone for the user. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # User (Optional) The user's preferred language, from the list of languages Canvas supports.
//    This is in RFC-5646 format.
// # User (Required) Whether the user accepts the terms of use.
// # Pseudonym (Required) User's login ID. Must be a valid email address.
// # CommunicationChannel (Optional) The communication channel type, e.g. 'email' or 'sms'.
// # CommunicationChannel (Optional) The communication channel address, e.g. the user's email address.
//
type DeprecatedSelfRegisterUser struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			Name         string `json:"name"`          //  (Required)
			ShortName    string `json:"short_name"`    //  (Optional)
			SortableName string `json:"sortable_name"` //  (Optional)
			TimeZone     string `json:"time_zone"`     //  (Optional)
			Locale       string `json:"locale"`        //  (Optional)
			TermsOfUse   bool   `json:"terms_of_use"`  //  (Required)
		} `json:"user"`

		Pseudonym struct {
			UniqueID string `json:"unique_id"` //  (Required)
		} `json:"pseudonym"`

		CommunicationChannel struct {
			Type    string `json:"type"`    //  (Optional)
			Address string `json:"address"` //  (Optional)
		} `json:"communication_channel"`
	} `json:"form"`
}

func (t *DeprecatedSelfRegisterUser) GetMethod() string {
	return "POST"
}

func (t *DeprecatedSelfRegisterUser) GetURLPath() string {
	path := "accounts/{account_id}/self_registration"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *DeprecatedSelfRegisterUser) GetQuery() (string, error) {
	return "", nil
}

func (t *DeprecatedSelfRegisterUser) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *DeprecatedSelfRegisterUser) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.User.Name == "" {
		errs = append(errs, "'User' is required")
	}
	if t.Form.Pseudonym.UniqueID == "" {
		errs = append(errs, "'Pseudonym' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeprecatedSelfRegisterUser) Do(c *canvasapi.Canvas) (*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
