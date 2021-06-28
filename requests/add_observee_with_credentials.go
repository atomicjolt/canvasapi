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

// AddObserveeWithCredentials Register the given user to observe another user, given the observee's credentials.
//
// *Note:* all users are allowed to add their own observees, given the observee's
// credentials or access token are provided. Administrators can add observees given credentials, access token or
// the {api:UserObserveesController#update observee's id}.
// https://canvas.instructure.com/doc/api/user_observees.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Form Parameters:
// # Observee (Optional) The login id for the user to observe.  Required if access_token is omitted.
// # Observee (Optional) The password for the user to observe. Required if access_token is omitted.
// # AccessToken (Optional) The access token for the user to observe.  Required if <tt>observee[unique_id]</tt> or <tt>observee[password]</tt> are omitted.
// # PairingCode (Optional) A generated pairing code for the user to observe. Required if the Observer pairing code feature flag is enabled
// # RootAccountID (Optional) The ID for the root account to associate with the observation link.
//    Defaults to the current domain account.
//    If 'all' is specified, a link will be created for each root account associated
//    to both the observer and observee.
//
type AddObserveeWithCredentials struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Observee struct {
			UniqueID string `json:"unique_id" url:"unique_id,omitempty"` //  (Optional)
			Password string `json:"password" url:"password,omitempty"`   //  (Optional)
		} `json:"observee" url:"observee,omitempty"`

		AccessToken   string `json:"access_token" url:"access_token,omitempty"`       //  (Optional)
		PairingCode   string `json:"pairing_code" url:"pairing_code,omitempty"`       //  (Optional)
		RootAccountID int64  `json:"root_account_id" url:"root_account_id,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *AddObserveeWithCredentials) GetMethod() string {
	return "POST"
}

func (t *AddObserveeWithCredentials) GetURLPath() string {
	path := "users/{user_id}/observees"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *AddObserveeWithCredentials) GetQuery() (string, error) {
	return "", nil
}

func (t *AddObserveeWithCredentials) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AddObserveeWithCredentials) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AddObserveeWithCredentials) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AddObserveeWithCredentials) Do(c *canvasapi.Canvas) (*models.User, error) {
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
